// Code generated from parser/Epha.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Epha
import "github.com/antlr4-go/antlr/v4"

// BaseEphaListener is a complete listener for a parse tree produced by EphaParser.
type BaseEphaListener struct{}

var _ EphaListener = &BaseEphaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEphaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEphaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEphaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEphaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseEphaListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseEphaListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseEphaListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseEphaListener) ExitStatement(ctx *StatementContext) {}

// EnterResourceDefinition is called when production resourceDefinition is entered.
func (s *BaseEphaListener) EnterResourceDefinition(ctx *ResourceDefinitionContext) {}

// ExitResourceDefinition is called when production resourceDefinition is exited.
func (s *BaseEphaListener) ExitResourceDefinition(ctx *ResourceDefinitionContext) {}

// EnterResourceBody is called when production resourceBody is entered.
func (s *BaseEphaListener) EnterResourceBody(ctx *ResourceBodyContext) {}

// ExitResourceBody is called when production resourceBody is exited.
func (s *BaseEphaListener) ExitResourceBody(ctx *ResourceBodyContext) {}

// EnterResourceProperty is called when production resourceProperty is entered.
func (s *BaseEphaListener) EnterResourceProperty(ctx *ResourcePropertyContext) {}

// ExitResourceProperty is called when production resourceProperty is exited.
func (s *BaseEphaListener) ExitResourceProperty(ctx *ResourcePropertyContext) {}

// EnterResourcePropertyBody is called when production resourcePropertyBody is entered.
func (s *BaseEphaListener) EnterResourcePropertyBody(ctx *ResourcePropertyBodyContext) {}

// ExitResourcePropertyBody is called when production resourcePropertyBody is exited.
func (s *BaseEphaListener) ExitResourcePropertyBody(ctx *ResourcePropertyBodyContext) {}

// EnterPropertyKey is called when production propertyKey is entered.
func (s *BaseEphaListener) EnterPropertyKey(ctx *PropertyKeyContext) {}

// ExitPropertyKey is called when production propertyKey is exited.
func (s *BaseEphaListener) ExitPropertyKey(ctx *PropertyKeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseEphaListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseEphaListener) ExitValue(ctx *ValueContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseEphaListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseEphaListener) ExitValueList(ctx *ValueListContext) {}
