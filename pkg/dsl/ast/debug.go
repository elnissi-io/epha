// In epha/pkg/dsl/ast/ast.go
package ast

import (
	"bytes"
	"fmt"
)

// String returns the string representation of the AST node.
func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

// ImportStatement represents import statements.
func (is *ImportStatement) String() string {
	var out bytes.Buffer

	out.WriteString("ImportStatement(")
	out.WriteString("Module=")
	out.WriteString(is.Module.String())

	if is.Alias != nil {
		out.WriteString(", Alias=")
		out.WriteString(is.Alias.String())
	}

	out.WriteString(", Imports=")
	for i, imp := range is.Imports {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(imp.String())
	}

	out.WriteString(")")

	return out.String()
}

// ArrayLiteral represents array literals.
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	for i, elem := range al.Elements {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(fmt.Sprintf("%s", elem))
	}
	out.WriteString("]")
	return out.String()
}

// Identifier represents identifiers.
func (id *Identifier) String() string {
	return id.Value
}

// VariableDeclaration represents variable declarations.
func (vd *VariableDeclaration) String() string {
	var out bytes.Buffer

	out.WriteString("VariableDeclaration(")
	out.WriteString("Name=")
	out.WriteString(vd.Name.String())
	out.WriteString(", Value=")
	out.WriteString(vd.Value.String())
	out.WriteString(")\n")

	return out.String()
}

// StringLiteral represents string literals.
func (sl *StringLiteral) String() string {
	return fmt.Sprintf("\"%s\"", sl.Value)
}

// HashLiteral represents hash literals.
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for key, value := range hl.Pairs {
		out.WriteString(fmt.Sprintf("%s: %s, ", key, value))
	}
	out.WriteString("}")
	return out.String()
}

// AssignmentStatement represents assignment statements.
func (as *AssignmentStatement) String() string {
	var out bytes.Buffer

	out.WriteString("AssignmentStatement(")
	out.WriteString("Name=")
	out.WriteString(as.Name.String())
	out.WriteString(", Value=")
	out.WriteString(as.Value.String())
	out.WriteString(")\n")

	return out.String()
}

// IntegerLiteral represents integer literals.
func (il *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

// BooleanLiteral represents boolean literals.
func (bl *BooleanLiteral) String() string {
	return fmt.Sprintf("%t", bl.Value)
}
