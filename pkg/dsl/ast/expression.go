package ast

import (
	"bytes"
	"epha/pkg/dsl/token"
	"strings"
)

type Expression interface {
    Node
    expressionNode()
    String() string
}

type CallExpression struct {
    Token     token.Token // The token. For instance, the function name
    Function  Expression  // Identifier or function literal
    Arguments []Expression
}

func (i *Identifier) expressionNode() {}
func (sl *StringLiteral) expressionNode() {}
func (hl *HashLiteral) expressionNode() {}
func (il *IntegerLiteral) expressionNode() {}
func (bl *BooleanLiteral) expressionNode() {}
func (al *ArrayLiteral) expressionNode() {}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) statementNode()  {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal } // Added

func (ce *CallExpression) String() string {
    var out bytes.Buffer
    args := []string{}
    for _, a := range ce.Arguments {
        args = append(args, a.String())
    }
    out.WriteString(ce.Function.String())
    out.WriteString("(")
    out.WriteString(strings.Join(args, ", "))
    out.WriteString(")")
    return out.String()
}