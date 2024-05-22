// Code generated from antlr/Epha.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Epha
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 126,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	7, 2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 37, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	7, 4, 47, 10, 4, 12, 4, 14, 4, 50, 11, 4, 5, 4, 52, 10, 4, 5, 4, 54, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 64, 10, 6, 12,
	6, 14, 6, 67, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 75, 10,
	7, 12, 7, 14, 7, 78, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 96, 10,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 102, 10, 11, 12, 11, 14, 11, 105,
	11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 119, 10, 12, 12, 12, 14, 12, 122, 11, 12, 3,
	12, 3, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2,
	2, 2, 130, 2, 27, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8,
	55, 3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 81, 3, 2, 2,
	2, 16, 85, 3, 2, 2, 2, 18, 95, 3, 2, 2, 2, 20, 97, 3, 2, 2, 2, 22, 108,
	3, 2, 2, 2, 24, 26, 5, 4, 3, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2,
	27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2, 29, 27, 3,
	2, 2, 2, 30, 31, 7, 2, 2, 3, 31, 3, 3, 2, 2, 2, 32, 37, 5, 6, 4, 2, 33,
	37, 5, 8, 5, 2, 34, 37, 5, 10, 6, 2, 35, 37, 5, 12, 7, 2, 36, 32, 3, 2,
	2, 2, 36, 33, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 35, 3, 2, 2, 2, 37, 5,
	3, 2, 2, 2, 38, 39, 7, 3, 2, 2, 39, 53, 7, 14, 2, 2, 40, 41, 7, 4, 2, 2,
	41, 51, 7, 14, 2, 2, 42, 43, 7, 3, 2, 2, 43, 48, 7, 14, 2, 2, 44, 45, 7,
	5, 2, 2, 45, 47, 7, 14, 2, 2, 46, 44, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48,
	46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2,
	2, 51, 42, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 40,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 7, 3, 2, 2, 2, 55, 56, 7, 14, 2, 2,
	56, 57, 7, 6, 2, 2, 57, 58, 5, 18, 10, 2, 58, 9, 3, 2, 2, 2, 59, 60, 7,
	7, 2, 2, 60, 61, 7, 14, 2, 2, 61, 65, 7, 8, 2, 2, 62, 64, 5, 14, 8, 2,
	63, 62, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3,
	2, 2, 2, 66, 68, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 69, 7, 9, 2, 2, 69,
	11, 3, 2, 2, 2, 70, 71, 7, 10, 2, 2, 71, 72, 7, 14, 2, 2, 72, 76, 7, 8,
	2, 2, 73, 75, 5, 16, 9, 2, 74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76,
	74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 79, 3, 2, 2, 2, 78, 76, 3, 2, 2,
	2, 79, 80, 7, 9, 2, 2, 80, 13, 3, 2, 2, 2, 81, 82, 7, 14, 2, 2, 82, 83,
	7, 11, 2, 2, 83, 84, 5, 18, 10, 2, 84, 15, 3, 2, 2, 2, 85, 86, 7, 14, 2,
	2, 86, 87, 7, 11, 2, 2, 87, 88, 5, 18, 10, 2, 88, 17, 3, 2, 2, 2, 89, 96,
	7, 15, 2, 2, 90, 96, 7, 16, 2, 2, 91, 96, 7, 17, 2, 2, 92, 96, 7, 14, 2,
	2, 93, 96, 5, 20, 11, 2, 94, 96, 5, 22, 12, 2, 95, 89, 3, 2, 2, 2, 95,
	90, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2,
	2, 95, 94, 3, 2, 2, 2, 96, 19, 3, 2, 2, 2, 97, 98, 7, 12, 2, 2, 98, 103,
	5, 18, 10, 2, 99, 100, 7, 5, 2, 2, 100, 102, 5, 18, 10, 2, 101, 99, 3,
	2, 2, 2, 102, 105, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2,
	2, 104, 106, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 106, 107, 7, 13, 2, 2, 107,
	21, 3, 2, 2, 2, 108, 109, 7, 8, 2, 2, 109, 110, 5, 18, 10, 2, 110, 111,
	7, 11, 2, 2, 111, 112, 5, 18, 10, 2, 112, 120, 3, 2, 2, 2, 113, 114, 7,
	5, 2, 2, 114, 115, 5, 18, 10, 2, 115, 116, 7, 11, 2, 2, 116, 117, 5, 18,
	10, 2, 117, 119, 3, 2, 2, 2, 118, 113, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2,
	120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122,
	120, 3, 2, 2, 2, 123, 124, 7, 9, 2, 2, 124, 23, 3, 2, 2, 2, 12, 27, 36,
	48, 51, 53, 65, 76, 95, 103, 120,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'import'", "'from'", "','", "'='", "'helm'", "'{'", "'}'", "'k8s'",
	"':'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER", "STRING",
	"NUMBER", "BOOLEAN", "WS",
}

var ruleNames = []string{
	"program", "statement", "importStmt", "variableDecl", "helmChart", "k8sResource",
	"helmStmt", "k8sStmt", "expression", "arrayLiteral", "hashLiteral",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type EphaParser struct {
	*antlr.BaseParser
}

func NewEphaParser(input antlr.TokenStream) *EphaParser {
	this := new(EphaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Epha.g4"

	return this
}

// EphaParser tokens.
const (
	EphaParserEOF        = antlr.TokenEOF
	EphaParserT__0       = 1
	EphaParserT__1       = 2
	EphaParserT__2       = 3
	EphaParserT__3       = 4
	EphaParserT__4       = 5
	EphaParserT__5       = 6
	EphaParserT__6       = 7
	EphaParserT__7       = 8
	EphaParserT__8       = 9
	EphaParserT__9       = 10
	EphaParserT__10      = 11
	EphaParserIDENTIFIER = 12
	EphaParserSTRING     = 13
	EphaParserNUMBER     = 14
	EphaParserBOOLEAN    = 15
	EphaParserWS         = 16
)

// EphaParser rules.
const (
	EphaParserRULE_program      = 0
	EphaParserRULE_statement    = 1
	EphaParserRULE_importStmt   = 2
	EphaParserRULE_variableDecl = 3
	EphaParserRULE_helmChart    = 4
	EphaParserRULE_k8sResource  = 5
	EphaParserRULE_helmStmt     = 6
	EphaParserRULE_k8sStmt      = 7
	EphaParserRULE_expression   = 8
	EphaParserRULE_arrayLiteral = 9
	EphaParserRULE_hashLiteral  = 10
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(EphaParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *EphaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EphaParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EphaParserT__0)|(1<<EphaParserT__4)|(1<<EphaParserT__7)|(1<<EphaParserIDENTIFIER))) != 0 {
		{
			p.SetState(22)
			p.Statement()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(28)
		p.Match(EphaParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ImportStmt() IImportStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportStmtContext)
}

func (s *StatementContext) VariableDecl() IVariableDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *StatementContext) HelmChart() IHelmChartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHelmChartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHelmChartContext)
}

func (s *StatementContext) K8sResource() IK8sResourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IK8sResourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IK8sResourceContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *EphaParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EphaParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EphaParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.ImportStmt()
		}

	case EphaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.VariableDecl()
		}

	case EphaParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.HelmChart()
		}

	case EphaParserT__7:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(33)
			p.K8sResource()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_importStmt
	return p
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_importStmt

	return p
}

func (s *ImportStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStmtContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(EphaParserIDENTIFIER)
}

func (s *ImportStmtContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, i)
}

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterImportStmt(s)
	}
}

func (s *ImportStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitImportStmt(s)
	}
}

func (p *EphaParser) ImportStmt() (localctx IImportStmtContext) {
	localctx = NewImportStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EphaParserRULE_importStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(EphaParserT__0)
	}
	{
		p.SetState(37)
		p.Match(EphaParserIDENTIFIER)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EphaParserT__1 {
		{
			p.SetState(38)
			p.Match(EphaParserT__1)
		}
		{
			p.SetState(39)
			p.Match(EphaParserIDENTIFIER)
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(40)
				p.Match(EphaParserT__0)
			}
			{
				p.SetState(41)
				p.Match(EphaParserIDENTIFIER)
			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == EphaParserT__2 {
				{
					p.SetState(42)
					p.Match(EphaParserT__2)
				}
				{
					p.SetState(43)
					p.Match(EphaParserIDENTIFIER)
				}

				p.SetState(48)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}

	}

	return localctx
}

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_variableDecl
	return p
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *VariableDeclContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterVariableDecl(s)
	}
}

func (s *VariableDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitVariableDecl(s)
	}
}

func (p *EphaParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EphaParserRULE_variableDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(EphaParserIDENTIFIER)
	}
	{
		p.SetState(54)
		p.Match(EphaParserT__3)
	}
	{
		p.SetState(55)
		p.Expression()
	}

	return localctx
}

// IHelmChartContext is an interface to support dynamic dispatch.
type IHelmChartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHelmChartContext differentiates from other interfaces.
	IsHelmChartContext()
}

type HelmChartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelmChartContext() *HelmChartContext {
	var p = new(HelmChartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_helmChart
	return p
}

func (*HelmChartContext) IsHelmChartContext() {}

func NewHelmChartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelmChartContext {
	var p = new(HelmChartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_helmChart

	return p
}

func (s *HelmChartContext) GetParser() antlr.Parser { return s.parser }

func (s *HelmChartContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *HelmChartContext) AllHelmStmt() []IHelmStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHelmStmtContext)(nil)).Elem())
	var tst = make([]IHelmStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHelmStmtContext)
		}
	}

	return tst
}

func (s *HelmChartContext) HelmStmt(i int) IHelmStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHelmStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHelmStmtContext)
}

func (s *HelmChartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelmChartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HelmChartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterHelmChart(s)
	}
}

func (s *HelmChartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitHelmChart(s)
	}
}

func (p *EphaParser) HelmChart() (localctx IHelmChartContext) {
	localctx = NewHelmChartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EphaParserRULE_helmChart)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(EphaParserT__4)
	}
	{
		p.SetState(58)
		p.Match(EphaParserIDENTIFIER)
	}
	{
		p.SetState(59)
		p.Match(EphaParserT__5)
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserIDENTIFIER {
		{
			p.SetState(60)
			p.HelmStmt()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(EphaParserT__6)
	}

	return localctx
}

// IK8sResourceContext is an interface to support dynamic dispatch.
type IK8sResourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsK8sResourceContext differentiates from other interfaces.
	IsK8sResourceContext()
}

type K8sResourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyK8sResourceContext() *K8sResourceContext {
	var p = new(K8sResourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_k8sResource
	return p
}

func (*K8sResourceContext) IsK8sResourceContext() {}

func NewK8sResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *K8sResourceContext {
	var p = new(K8sResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_k8sResource

	return p
}

func (s *K8sResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *K8sResourceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *K8sResourceContext) AllK8sStmt() []IK8sStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IK8sStmtContext)(nil)).Elem())
	var tst = make([]IK8sStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IK8sStmtContext)
		}
	}

	return tst
}

func (s *K8sResourceContext) K8sStmt(i int) IK8sStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IK8sStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IK8sStmtContext)
}

func (s *K8sResourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *K8sResourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *K8sResourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterK8sResource(s)
	}
}

func (s *K8sResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitK8sResource(s)
	}
}

func (p *EphaParser) K8sResource() (localctx IK8sResourceContext) {
	localctx = NewK8sResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EphaParserRULE_k8sResource)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(EphaParserT__7)
	}
	{
		p.SetState(69)
		p.Match(EphaParserIDENTIFIER)
	}
	{
		p.SetState(70)
		p.Match(EphaParserT__5)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserIDENTIFIER {
		{
			p.SetState(71)
			p.K8sStmt()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(EphaParserT__6)
	}

	return localctx
}

// IHelmStmtContext is an interface to support dynamic dispatch.
type IHelmStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHelmStmtContext differentiates from other interfaces.
	IsHelmStmtContext()
}

type HelmStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelmStmtContext() *HelmStmtContext {
	var p = new(HelmStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_helmStmt
	return p
}

func (*HelmStmtContext) IsHelmStmtContext() {}

func NewHelmStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelmStmtContext {
	var p = new(HelmStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_helmStmt

	return p
}

func (s *HelmStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *HelmStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *HelmStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HelmStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelmStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HelmStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterHelmStmt(s)
	}
}

func (s *HelmStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitHelmStmt(s)
	}
}

func (p *EphaParser) HelmStmt() (localctx IHelmStmtContext) {
	localctx = NewHelmStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EphaParserRULE_helmStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(EphaParserIDENTIFIER)
	}
	{
		p.SetState(80)
		p.Match(EphaParserT__8)
	}
	{
		p.SetState(81)
		p.Expression()
	}

	return localctx
}

// IK8sStmtContext is an interface to support dynamic dispatch.
type IK8sStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsK8sStmtContext differentiates from other interfaces.
	IsK8sStmtContext()
}

type K8sStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyK8sStmtContext() *K8sStmtContext {
	var p = new(K8sStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_k8sStmt
	return p
}

func (*K8sStmtContext) IsK8sStmtContext() {}

func NewK8sStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *K8sStmtContext {
	var p = new(K8sStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_k8sStmt

	return p
}

func (s *K8sStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *K8sStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *K8sStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *K8sStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *K8sStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *K8sStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterK8sStmt(s)
	}
}

func (s *K8sStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitK8sStmt(s)
	}
}

func (p *EphaParser) K8sStmt() (localctx IK8sStmtContext) {
	localctx = NewK8sStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EphaParserRULE_k8sStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(EphaParserIDENTIFIER)
	}
	{
		p.SetState(84)
		p.Match(EphaParserT__8)
	}
	{
		p.SetState(85)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(EphaParserSTRING, 0)
}

func (s *ExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(EphaParserNUMBER, 0)
}

func (s *ExpressionContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(EphaParserBOOLEAN, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *ExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ExpressionContext) HashLiteral() IHashLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHashLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHashLiteralContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EphaParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EphaParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EphaParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(EphaParserSTRING)
		}

	case EphaParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(EphaParserNUMBER)
		}

	case EphaParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(EphaParserBOOLEAN)
		}

	case EphaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)
			p.Match(EphaParserIDENTIFIER)
		}

	case EphaParserT__9:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)
			p.ArrayLiteral()
		}

	case EphaParserT__5:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(92)
			p.HashLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *EphaParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, EphaParserRULE_arrayLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(EphaParserT__9)
	}
	{
		p.SetState(96)
		p.Expression()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserT__2 {
		{
			p.SetState(97)
			p.Match(EphaParserT__2)
		}
		{
			p.SetState(98)
			p.Expression()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
		p.Match(EphaParserT__10)
	}

	return localctx
}

// IHashLiteralContext is an interface to support dynamic dispatch.
type IHashLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHashLiteralContext differentiates from other interfaces.
	IsHashLiteralContext()
}

type HashLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHashLiteralContext() *HashLiteralContext {
	var p = new(HashLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EphaParserRULE_hashLiteral
	return p
}

func (*HashLiteralContext) IsHashLiteralContext() {}

func NewHashLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashLiteralContext {
	var p = new(HashLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_hashLiteral

	return p
}

func (s *HashLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *HashLiteralContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *HashLiteralContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HashLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HashLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterHashLiteral(s)
	}
}

func (s *HashLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitHashLiteral(s)
	}
}

func (p *EphaParser) HashLiteral() (localctx IHashLiteralContext) {
	localctx = NewHashLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, EphaParserRULE_hashLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(EphaParserT__5)
	}

	{
		p.SetState(107)
		p.Expression()
	}
	{
		p.SetState(108)
		p.Match(EphaParserT__8)
	}
	{
		p.SetState(109)
		p.Expression()
	}

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserT__2 {
		{
			p.SetState(111)
			p.Match(EphaParserT__2)
		}
		{
			p.SetState(112)
			p.Expression()
		}
		{
			p.SetState(113)
			p.Match(EphaParserT__8)
		}
		{
			p.SetState(114)
			p.Expression()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(EphaParserT__6)
	}

	return localctx
}
