package listener

import (
	"epha/pkg/dsl/ast"
	"epha/pkg/dsl/token"
	"epha/pkg/parser"
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
	l.program.Statements = append(l.program.Statements, importStmt)
}

// Implement other listener methods for variableDecl, helmChart, k8sResource, etc.

func (l *EphaListener) GetProgram() *ast.Program {
	return l.program
}
