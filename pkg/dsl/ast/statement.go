package ast

import "epha/pkg/dsl/token"

type Statement interface {
	Node
	statementNode()
	String() string // Define the String() method for Statement
}

// Represents import statements
type ImportStatement struct {
	Token   token.Token   // The 'import' or 'from' token
	Module  *Identifier   // The module name (e.g., 'kubernetes', 'crds')
	Imports []*Identifier // Specific imports from the module
	Alias   *Identifier   // Optional alias for the import
}

// Represents variable declarations/assignments
type AssignmentStatement struct {
	Token token.Token // The token.IDENTIFIER token for the variable name
	Name  *Identifier // The variable name
	Value Expression  // The value/expression assigned to the variable
}

// Statement node method implementations
func (is *ImportStatement) statementNode()     {}
func (as *AssignmentStatement) statementNode() {}
func (vd *VariableDeclaration) statementNode() {}
