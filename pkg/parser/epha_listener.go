// Code generated from parser/Epha.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Epha
import "github.com/antlr4-go/antlr/v4"

// EphaListener is a complete listener for a parse tree produced by EphaParser.
type EphaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterResourceDefinition is called when entering the resourceDefinition production.
	EnterResourceDefinition(c *ResourceDefinitionContext)

	// EnterResourceBody is called when entering the resourceBody production.
	EnterResourceBody(c *ResourceBodyContext)

	// EnterResourceProperty is called when entering the resourceProperty production.
	EnterResourceProperty(c *ResourcePropertyContext)

	// EnterResourcePropertyBody is called when entering the resourcePropertyBody production.
	EnterResourcePropertyBody(c *ResourcePropertyBodyContext)

	// EnterPropertyKey is called when entering the propertyKey production.
	EnterPropertyKey(c *PropertyKeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitResourceDefinition is called when exiting the resourceDefinition production.
	ExitResourceDefinition(c *ResourceDefinitionContext)

	// ExitResourceBody is called when exiting the resourceBody production.
	ExitResourceBody(c *ResourceBodyContext)

	// ExitResourceProperty is called when exiting the resourceProperty production.
	ExitResourceProperty(c *ResourcePropertyContext)

	// ExitResourcePropertyBody is called when exiting the resourcePropertyBody production.
	ExitResourcePropertyBody(c *ResourcePropertyBodyContext)

	// ExitPropertyKey is called when exiting the propertyKey production.
	ExitPropertyKey(c *PropertyKeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)
}
