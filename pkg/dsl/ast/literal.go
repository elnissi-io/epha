package ast

import "epha/pkg/dsl/token"

type StringLiteral struct {
    Token token.Token
    Value string
}

type ObjectLiteral struct {
    Token token.Token // The '{' token
    Pairs map[string]Expression
}

type ArrayLiteral struct {
    Token    token.Token // The '[' token
    Elements []Expression
}

type HashLiteral struct {
    Token token.Token // the '{' token
    Pairs map[Expression]Expression
}

type IntegerLiteral struct {
    Token token.Token
    Value int64
}

type BooleanLiteral struct {
    Token token.Token
    Value bool
}

func (i *Identifier) TokenLiteral() string           { return i.Token.Literal }
func (sl *StringLiteral) TokenLiteral() string       { return sl.Token.Literal }
func (il *IntegerLiteral) TokenLiteral() string      { return il.Token.Literal }
func (bl *BooleanLiteral) TokenLiteral() string      { return bl.Token.Literal }  // Fixed
func (al *ArrayLiteral) TokenLiteral() string        { return al.Token.Literal }
func (hl *HashLiteral) TokenLiteral() string         { return hl.Token.Literal }
func (is *ImportStatement) TokenLiteral() string     { return is.Token.Literal }
func (vd *VariableDeclaration) TokenLiteral() string { return vd.Token.Literal }
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Literal }
