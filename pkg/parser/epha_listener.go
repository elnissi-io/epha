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

	// EnterImportStmt is called when entering the importStmt production.
	EnterImportStmt(c *ImportStmtContext)

	// EnterVariableDecl is called when entering the variableDecl production.
	EnterVariableDecl(c *VariableDeclContext)

	// EnterHelmChart is called when entering the helmChart production.
	EnterHelmChart(c *HelmChartContext)

	// EnterK8sResource is called when entering the k8sResource production.
	EnterK8sResource(c *K8sResourceContext)

	// EnterHelmStmt is called when entering the helmStmt production.
	EnterHelmStmt(c *HelmStmtContext)

	// EnterK8sStmt is called when entering the k8sStmt production.
	EnterK8sStmt(c *K8sStmtContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterHashLiteral is called when entering the hashLiteral production.
	EnterHashLiteral(c *HashLiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitImportStmt is called when exiting the importStmt production.
	ExitImportStmt(c *ImportStmtContext)

	// ExitVariableDecl is called when exiting the variableDecl production.
	ExitVariableDecl(c *VariableDeclContext)

	// ExitHelmChart is called when exiting the helmChart production.
	ExitHelmChart(c *HelmChartContext)

	// ExitK8sResource is called when exiting the k8sResource production.
	ExitK8sResource(c *K8sResourceContext)

	// ExitHelmStmt is called when exiting the helmStmt production.
	ExitHelmStmt(c *HelmStmtContext)

	// ExitK8sStmt is called when exiting the k8sStmt production.
	ExitK8sStmt(c *K8sStmtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitHashLiteral is called when exiting the hashLiteral production.
	ExitHashLiteral(c *HashLiteralContext)
}
