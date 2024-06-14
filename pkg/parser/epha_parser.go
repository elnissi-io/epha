// Code generated from parser/Epha.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Epha
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type EphaParser struct {
	*antlr.BaseParser
}

var EphaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ephaParserInit() {
	staticData := &EphaParserStaticData
	staticData.LiteralNames = []string{
		"", "'import'", "'from'", "','", "'='", "'helm'", "'{'", "'}'", "'k8s'",
		"':'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER", "STRING",
		"NUMBER", "BOOLEAN", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "importStmt", "variableDecl", "helmChart", "k8sResource",
		"helmStmt", "k8sStmt", "expression", "arrayLiteral", "hashLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 124, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 35, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 45, 8, 2, 10, 2, 12, 2, 48, 9, 2, 3, 2, 50, 8, 2, 3, 2, 52,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 62, 8, 4, 10,
		4, 12, 4, 65, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 73, 8, 5,
		10, 5, 12, 5, 76, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 94, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 5, 9, 100, 8, 9, 10, 9, 12, 9, 103, 9, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		117, 8, 10, 10, 10, 12, 10, 120, 9, 10, 1, 10, 1, 10, 1, 10, 0, 0, 11,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 0, 128, 0, 25, 1, 0, 0, 0, 2,
		34, 1, 0, 0, 0, 4, 36, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0,
		10, 68, 1, 0, 0, 0, 12, 79, 1, 0, 0, 0, 14, 83, 1, 0, 0, 0, 16, 93, 1,
		0, 0, 0, 18, 95, 1, 0, 0, 0, 20, 106, 1, 0, 0, 0, 22, 24, 3, 2, 1, 0, 23,
		22, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0,
		0, 26, 28, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 5, 0, 0, 1, 29, 1, 1,
		0, 0, 0, 30, 35, 3, 4, 2, 0, 31, 35, 3, 6, 3, 0, 32, 35, 3, 8, 4, 0, 33,
		35, 3, 10, 5, 0, 34, 30, 1, 0, 0, 0, 34, 31, 1, 0, 0, 0, 34, 32, 1, 0,
		0, 0, 34, 33, 1, 0, 0, 0, 35, 3, 1, 0, 0, 0, 36, 37, 5, 1, 0, 0, 37, 51,
		5, 12, 0, 0, 38, 39, 5, 2, 0, 0, 39, 49, 5, 12, 0, 0, 40, 41, 5, 1, 0,
		0, 41, 46, 5, 12, 0, 0, 42, 43, 5, 3, 0, 0, 43, 45, 5, 12, 0, 0, 44, 42,
		1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0,
		47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49, 40, 1, 0, 0, 0, 49, 50, 1,
		0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 38, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		5, 1, 0, 0, 0, 53, 54, 5, 12, 0, 0, 54, 55, 5, 4, 0, 0, 55, 56, 3, 16,
		8, 0, 56, 7, 1, 0, 0, 0, 57, 58, 5, 5, 0, 0, 58, 59, 5, 12, 0, 0, 59, 63,
		5, 6, 0, 0, 60, 62, 3, 12, 6, 0, 61, 60, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0,
		63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 63, 1,
		0, 0, 0, 66, 67, 5, 7, 0, 0, 67, 9, 1, 0, 0, 0, 68, 69, 5, 8, 0, 0, 69,
		70, 5, 12, 0, 0, 70, 74, 5, 6, 0, 0, 71, 73, 3, 14, 7, 0, 72, 71, 1, 0,
		0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77,
		1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 78, 5, 7, 0, 0, 78, 11, 1, 0, 0, 0,
		79, 80, 5, 12, 0, 0, 80, 81, 5, 9, 0, 0, 81, 82, 3, 16, 8, 0, 82, 13, 1,
		0, 0, 0, 83, 84, 5, 12, 0, 0, 84, 85, 5, 9, 0, 0, 85, 86, 3, 16, 8, 0,
		86, 15, 1, 0, 0, 0, 87, 94, 5, 13, 0, 0, 88, 94, 5, 14, 0, 0, 89, 94, 5,
		15, 0, 0, 90, 94, 5, 12, 0, 0, 91, 94, 3, 18, 9, 0, 92, 94, 3, 20, 10,
		0, 93, 87, 1, 0, 0, 0, 93, 88, 1, 0, 0, 0, 93, 89, 1, 0, 0, 0, 93, 90,
		1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 17, 1, 0, 0, 0,
		95, 96, 5, 10, 0, 0, 96, 101, 3, 16, 8, 0, 97, 98, 5, 3, 0, 0, 98, 100,
		3, 16, 8, 0, 99, 97, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0,
		0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104,
		105, 5, 11, 0, 0, 105, 19, 1, 0, 0, 0, 106, 107, 5, 6, 0, 0, 107, 108,
		3, 16, 8, 0, 108, 109, 5, 9, 0, 0, 109, 110, 3, 16, 8, 0, 110, 118, 1,
		0, 0, 0, 111, 112, 5, 3, 0, 0, 112, 113, 3, 16, 8, 0, 113, 114, 5, 9, 0,
		0, 114, 115, 3, 16, 8, 0, 115, 117, 1, 0, 0, 0, 116, 111, 1, 0, 0, 0, 117,
		120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 121,
		1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 122, 5, 7, 0, 0, 122, 21, 1, 0,
		0, 0, 10, 25, 34, 46, 49, 51, 63, 74, 93, 101, 118,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// EphaParserInit initializes any static state used to implement EphaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEphaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EphaParserInit() {
	staticData := &EphaParserStaticData
	staticData.once.Do(ephaParserInit)
}

// NewEphaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEphaParser(input antlr.TokenStream) *EphaParser {
	EphaParserInit()
	this := new(EphaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EphaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
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

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(EphaParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4386) != 0 {
		{
			p.SetState(22)
			p.Statement()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(28)
		p.Match(EphaParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ImportStmt() IImportStmtContext
	VariableDecl() IVariableDeclContext
	HelmChart() IHelmChartContext
	K8sResource() IK8sResourceContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ImportStmt() IImportStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStmtContext)
}

func (s *StatementContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *StatementContext) HelmChart() IHelmChartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHelmChartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHelmChartContext)
}

func (s *StatementContext) K8sResource() IK8sResourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IK8sResourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

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
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_importStmt
	return p
}

func InitEmptyImportStmtContext(p *ImportStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_importStmt
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(EphaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EphaParserT__1 {
		{
			p.SetState(38)
			p.Match(EphaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(EphaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(40)
				p.Match(EphaParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(41)
				p.Match(EphaParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == EphaParserT__2 {
				{
					p.SetState(42)
					p.Match(EphaParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(43)
					p.Match(EphaParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(48)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *VariableDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(EphaParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHelmChartContext is an interface to support dynamic dispatch.
type IHelmChartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllHelmStmt() []IHelmStmtContext
	HelmStmt(i int) IHelmStmtContext

	// IsHelmChartContext differentiates from other interfaces.
	IsHelmChartContext()
}

type HelmChartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelmChartContext() *HelmChartContext {
	var p = new(HelmChartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_helmChart
	return p
}

func InitEmptyHelmChartContext(p *HelmChartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_helmChart
}

func (*HelmChartContext) IsHelmChartContext() {}

func NewHelmChartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelmChartContext {
	var p = new(HelmChartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_helmChart

	return p
}

func (s *HelmChartContext) GetParser() antlr.Parser { return s.parser }

func (s *HelmChartContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *HelmChartContext) AllHelmStmt() []IHelmStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHelmStmtContext); ok {
			len++
		}
	}

	tst := make([]IHelmStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHelmStmtContext); ok {
			tst[i] = t.(IHelmStmtContext)
			i++
		}
	}

	return tst
}

func (s *HelmChartContext) HelmStmt(i int) IHelmStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHelmStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(EphaParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(EphaParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserIDENTIFIER {
		{
			p.SetState(60)
			p.HelmStmt()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(EphaParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IK8sResourceContext is an interface to support dynamic dispatch.
type IK8sResourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllK8sStmt() []IK8sStmtContext
	K8sStmt(i int) IK8sStmtContext

	// IsK8sResourceContext differentiates from other interfaces.
	IsK8sResourceContext()
}

type K8sResourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyK8sResourceContext() *K8sResourceContext {
	var p = new(K8sResourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_k8sResource
	return p
}

func InitEmptyK8sResourceContext(p *K8sResourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_k8sResource
}

func (*K8sResourceContext) IsK8sResourceContext() {}

func NewK8sResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *K8sResourceContext {
	var p = new(K8sResourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_k8sResource

	return p
}

func (s *K8sResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *K8sResourceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *K8sResourceContext) AllK8sStmt() []IK8sStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IK8sStmtContext); ok {
			len++
		}
	}

	tst := make([]IK8sStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IK8sStmtContext); ok {
			tst[i] = t.(IK8sStmtContext)
			i++
		}
	}

	return tst
}

func (s *K8sResourceContext) K8sStmt(i int) IK8sStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IK8sStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(EphaParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(EphaParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserIDENTIFIER {
		{
			p.SetState(71)
			p.K8sStmt()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(EphaParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHelmStmtContext is an interface to support dynamic dispatch.
type IHelmStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsHelmStmtContext differentiates from other interfaces.
	IsHelmStmtContext()
}

type HelmStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelmStmtContext() *HelmStmtContext {
	var p = new(HelmStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_helmStmt
	return p
}

func InitEmptyHelmStmtContext(p *HelmStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_helmStmt
}

func (*HelmStmtContext) IsHelmStmtContext() {}

func NewHelmStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelmStmtContext {
	var p = new(HelmStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_helmStmt

	return p
}

func (s *HelmStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *HelmStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *HelmStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(EphaParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IK8sStmtContext is an interface to support dynamic dispatch.
type IK8sStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsK8sStmtContext differentiates from other interfaces.
	IsK8sStmtContext()
}

type K8sStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyK8sStmtContext() *K8sStmtContext {
	var p = new(K8sStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_k8sStmt
	return p
}

func InitEmptyK8sStmtContext(p *K8sStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_k8sStmt
}

func (*K8sStmtContext) IsK8sStmtContext() {}

func NewK8sStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *K8sStmtContext {
	var p = new(K8sStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_k8sStmt

	return p
}

func (s *K8sStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *K8sStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, 0)
}

func (s *K8sStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(EphaParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext
	HashLiteral() IHashLiteralContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ExpressionContext) HashLiteral() IHashLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHashLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EphaParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(EphaParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EphaParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(EphaParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EphaParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(EphaParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EphaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)
			p.Match(EphaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(EphaParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Expression()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserT__2 {
		{
			p.SetState(97)
			p.Match(EphaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Expression()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
		p.Match(EphaParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHashLiteralContext is an interface to support dynamic dispatch.
type IHashLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsHashLiteralContext differentiates from other interfaces.
	IsHashLiteralContext()
}

type HashLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHashLiteralContext() *HashLiteralContext {
	var p = new(HashLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_hashLiteral
	return p
}

func InitEmptyHashLiteralContext(p *HashLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_hashLiteral
}

func (*HashLiteralContext) IsHashLiteralContext() {}

func NewHashLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashLiteralContext {
	var p = new(HashLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_hashLiteral

	return p
}

func (s *HashLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *HashLiteralContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *HashLiteralContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(EphaParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(107)
		p.Expression()
	}
	{
		p.SetState(108)
		p.Match(EphaParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Expression()
	}

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserT__2 {
		{
			p.SetState(111)
			p.Match(EphaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Expression()
		}
		{
			p.SetState(113)
			p.Match(EphaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Expression()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(EphaParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
