package ast

import "epha/pkg/dsl/token"

type Node interface {
	TokenLiteral() string
}

type Identifier struct {
	Token token.Token
	Value string
}

// The root node of every AST
type Program struct {
	Statements []Statement
}

// Represents variable declarations
type VariableDeclaration struct {
	Token token.Token // The token.IDENTIFIER token
	Name  *Identifier // The variable name
	Value Expression  // The value/expression assigned to the variable
}

// Function calls like k8s.service(...)
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  *Identifier // The callee, e.g., k8s.service
	Arguments []Expression
}
