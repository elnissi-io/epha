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

// EnterImportStmt is called when production importStmt is entered.
func (s *BaseEphaListener) EnterImportStmt(ctx *ImportStmtContext) {}

// ExitImportStmt is called when production importStmt is exited.
func (s *BaseEphaListener) ExitImportStmt(ctx *ImportStmtContext) {}

// EnterVariableDecl is called when production variableDecl is entered.
func (s *BaseEphaListener) EnterVariableDecl(ctx *VariableDeclContext) {}

// ExitVariableDecl is called when production variableDecl is exited.
func (s *BaseEphaListener) ExitVariableDecl(ctx *VariableDeclContext) {}

// EnterHelmChart is called when production helmChart is entered.
func (s *BaseEphaListener) EnterHelmChart(ctx *HelmChartContext) {}

// ExitHelmChart is called when production helmChart is exited.
func (s *BaseEphaListener) ExitHelmChart(ctx *HelmChartContext) {}

// EnterK8sResource is called when production k8sResource is entered.
func (s *BaseEphaListener) EnterK8sResource(ctx *K8sResourceContext) {}

// ExitK8sResource is called when production k8sResource is exited.
func (s *BaseEphaListener) ExitK8sResource(ctx *K8sResourceContext) {}

// EnterHelmStmt is called when production helmStmt is entered.
func (s *BaseEphaListener) EnterHelmStmt(ctx *HelmStmtContext) {}

// ExitHelmStmt is called when production helmStmt is exited.
func (s *BaseEphaListener) ExitHelmStmt(ctx *HelmStmtContext) {}

// EnterK8sStmt is called when production k8sStmt is entered.
func (s *BaseEphaListener) EnterK8sStmt(ctx *K8sStmtContext) {}

// ExitK8sStmt is called when production k8sStmt is exited.
func (s *BaseEphaListener) ExitK8sStmt(ctx *K8sStmtContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEphaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEphaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseEphaListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseEphaListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterHashLiteral is called when production hashLiteral is entered.
func (s *BaseEphaListener) EnterHashLiteral(ctx *HashLiteralContext) {}

// ExitHashLiteral is called when production hashLiteral is exited.
func (s *BaseEphaListener) ExitHashLiteral(ctx *HashLiteralContext) {}
