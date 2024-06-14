package listener

import (
	"epha/pkg/dsl/ast"
	"epha/pkg/dsl/token"
	"epha/pkg/parser"
	"strconv"
)

type EphaListener struct {
    *parser.BaseEphaListener
    program *ast.Program
}

func NewEphaListener() *EphaListener {
    return &EphaListener{program: &ast.Program{}}
}

func (l *EphaListener) EnterImportStmt(ctx *parser.ImportStmtContext) {
    importStmt := &ast.ImportStatement{
        Token: token.Token{Type: token.Import, Literal: "import"},
        Module: &ast.Identifier{
            Token: token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER(0).GetText()},
            Value: ctx.IDENTIFIER(0).GetText(),
        },
    }
    if ctx.IDENTIFIER(1) != nil {
        importStmt.Alias = &ast.Identifier{
            Token: token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER(1).GetText()},
            Value: ctx.IDENTIFIER(1).GetText(),
        }
    }
    for _, imp := range ctx.AllIDENTIFIER()[2:] {
        importStmt.Imports = append(importStmt.Imports, &ast.Identifier{
            Token: token.Token{Type: token.Identifier, Literal: imp.GetText()},
            Value: imp.GetText(),
        })
    }
    l.program.Statements = append(l.program.Statements, importStmt)
}

func (l *EphaListener) EnterVariableDecl(ctx *parser.VariableDeclContext) {
    varDecl := &ast.VariableDeclaration{
        Token: token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER().GetText()},
        Name: &ast.Identifier{
            Token: token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER().GetText()},
            Value: ctx.IDENTIFIER().GetText(),
        },
        Value: l.parseExpression(ctx.Expression().(*parser.ExpressionContext)),
    }
    l.program.Statements = append(l.program.Statements, varDecl)
}

func (l *EphaListener) EnterHelmChart(ctx *parser.HelmChartContext) {
    helmChart := &ast.CallExpression{
        Token:    token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER().GetText()},
        Function: &ast.Identifier{Token: token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER().GetText()}, Value: ctx.IDENTIFIER().GetText()},
    }
    for _, stmt := range ctx.AllHelmStmt() {
        helmChart.Arguments = append(helmChart.Arguments, l.parseExpression(stmt.(*parser.HelmStmtContext).Expression().(*parser.ExpressionContext)))
    }
    l.program.Statements = append(l.program.Statements, helmChart)
}

func (l *EphaListener) EnterK8sResource(ctx *parser.K8sResourceContext) {
    k8sResource := &ast.CallExpression{
        Token:    token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER().GetText()},
        Function: &ast.Identifier{Token: token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER().GetText()}, Value: ctx.IDENTIFIER().GetText()},
    }
    for _, stmt := range ctx.AllK8sStmt() {
        k8sResource.Arguments = append(k8sResource.Arguments, l.parseExpression(stmt.(*parser.K8sStmtContext).Expression().(*parser.ExpressionContext)))
    }
    l.program.Statements = append(l.program.Statements, k8sResource)
}

func (l *EphaListener) parseExpression(ctx *parser.ExpressionContext) ast.Expression {
    switch {
    case ctx.STRING() != nil:
        return &ast.StringLiteral{Token: token.Token{Type: token.String, Literal: ctx.STRING().GetText()}, Value: ctx.STRING().GetText()}
    case ctx.NUMBER() != nil:
        value, _ := strconv.ParseInt(ctx.NUMBER().GetText(), 10, 64)
        return &ast.IntegerLiteral{Token: token.Token{Type: token.Number, Literal: ctx.NUMBER().GetText()}, Value: value}
    case ctx.BOOLEAN() != nil:
        value := ctx.BOOLEAN().GetText() == "true"
        return &ast.BooleanLiteral{Token: token.Token{Type: token.Boolean, Literal: ctx.BOOLEAN().GetText()}, Value: value}
    case ctx.IDENTIFIER() != nil:
        return &ast.Identifier{Token: token.Token{Type: token.Identifier, Literal: ctx.IDENTIFIER().GetText()}, Value: ctx.IDENTIFIER().GetText()}
    case ctx.ArrayLiteral() != nil:
        return l.parseArrayLiteral(ctx.ArrayLiteral().(*parser.ArrayLiteralContext))
    case ctx.HashLiteral() != nil:
        return l.parseHashLiteral(ctx.HashLiteral().(*parser.HashLiteralContext))
    }
    return nil
}

func (l *EphaListener) parseArrayLiteral(ctx *parser.ArrayLiteralContext) ast.Expression {
    elements := []ast.Expression{}
    for _, exprCtx := range ctx.AllExpression() {
        elements = append(elements, l.parseExpression(exprCtx.(*parser.ExpressionContext)))
    }
    return &ast.ArrayLiteral{Token: token.Token{Type: token.LBracket, Literal: "["}, Elements: elements}
}

func (l *EphaListener) parseHashLiteral(ctx *parser.HashLiteralContext) ast.Expression {
    pairs := make(map[ast.Expression]ast.Expression)
    for i := 0; i < len(ctx.AllExpression()); i += 2 {
        key := l.parseExpression(ctx.AllExpression()[i].(*parser.ExpressionContext))
        value := l.parseExpression(ctx.AllExpression()[i+1].(*parser.ExpressionContext))
        pairs[key] = value
    }
    return &ast.HashLiteral{Token: token.Token{Type: token.LBrace, Literal: "{"}, Pairs: pairs}
}

func (l *EphaListener) GetProgram() *ast.Program {
    return l.program
}
