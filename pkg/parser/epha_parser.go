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
		"", "'import'", "'as'", "'{'", "'}'", "'='", "'.'", "'['", "','", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "IDENTIFIER", "STRING", "NUMBER",
		"LINE_COMMENT", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "statement", "importStatement", "resourceDefinition", "resourceBody",
		"resourceProperty", "resourcePropertyBody", "propertyKey", "value",
		"valueList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 90, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 4,
		0, 22, 8, 0, 11, 0, 12, 0, 23, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 34, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 5,
		4, 43, 8, 4, 10, 4, 12, 4, 46, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 61, 8, 5, 1, 6, 5, 6, 64,
		8, 6, 10, 6, 12, 6, 67, 9, 6, 1, 7, 1, 7, 1, 7, 5, 7, 72, 8, 7, 10, 7,
		12, 7, 75, 9, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 83, 8, 9, 10,
		9, 12, 9, 86, 9, 9, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 0, 1, 1, 0, 11, 12, 88, 0, 21, 1, 0, 0, 0, 2, 27, 1, 0, 0,
		0, 4, 29, 1, 0, 0, 0, 6, 35, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10, 60, 1,
		0, 0, 0, 12, 65, 1, 0, 0, 0, 14, 68, 1, 0, 0, 0, 16, 76, 1, 0, 0, 0, 18,
		78, 1, 0, 0, 0, 20, 22, 3, 2, 1, 0, 21, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0,
		0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 1, 1, 0, 0, 0, 25, 28, 3,
		4, 2, 0, 26, 28, 3, 6, 3, 0, 27, 25, 1, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28,
		3, 1, 0, 0, 0, 29, 30, 5, 1, 0, 0, 30, 33, 5, 10, 0, 0, 31, 32, 5, 2, 0,
		0, 32, 34, 5, 10, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 5,
		1, 0, 0, 0, 35, 36, 5, 10, 0, 0, 36, 37, 5, 10, 0, 0, 37, 38, 5, 3, 0,
		0, 38, 39, 3, 8, 4, 0, 39, 40, 5, 4, 0, 0, 40, 7, 1, 0, 0, 0, 41, 43, 3,
		10, 5, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44,
		45, 1, 0, 0, 0, 45, 9, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 3, 14, 7,
		0, 48, 49, 5, 5, 0, 0, 49, 50, 3, 16, 8, 0, 50, 61, 1, 0, 0, 0, 51, 52,
		3, 14, 7, 0, 52, 53, 5, 5, 0, 0, 53, 54, 3, 18, 9, 0, 54, 61, 1, 0, 0,
		0, 55, 56, 3, 14, 7, 0, 56, 57, 5, 3, 0, 0, 57, 58, 3, 12, 6, 0, 58, 59,
		5, 4, 0, 0, 59, 61, 1, 0, 0, 0, 60, 47, 1, 0, 0, 0, 60, 51, 1, 0, 0, 0,
		60, 55, 1, 0, 0, 0, 61, 11, 1, 0, 0, 0, 62, 64, 3, 10, 5, 0, 63, 62, 1,
		0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66,
		13, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 73, 5, 10, 0, 0, 69, 70, 5, 6,
		0, 0, 70, 72, 5, 10, 0, 0, 71, 69, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73,
		71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 15, 1, 0, 0, 0, 75, 73, 1, 0, 0,
		0, 76, 77, 7, 0, 0, 0, 77, 17, 1, 0, 0, 0, 78, 79, 5, 7, 0, 0, 79, 84,
		3, 16, 8, 0, 80, 81, 5, 8, 0, 0, 81, 83, 3, 16, 8, 0, 82, 80, 1, 0, 0,
		0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87,
		1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 88, 5, 9, 0, 0, 88, 19, 1, 0, 0, 0,
		8, 23, 27, 33, 44, 60, 65, 73, 84,
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
	EphaParserEOF          = antlr.TokenEOF
	EphaParserT__0         = 1
	EphaParserT__1         = 2
	EphaParserT__2         = 3
	EphaParserT__3         = 4
	EphaParserT__4         = 5
	EphaParserT__5         = 6
	EphaParserT__6         = 7
	EphaParserT__7         = 8
	EphaParserT__8         = 9
	EphaParserIDENTIFIER   = 10
	EphaParserSTRING       = 11
	EphaParserNUMBER       = 12
	EphaParserLINE_COMMENT = 13
	EphaParserWHITESPACE   = 14
)

// EphaParser rules.
const (
	EphaParserRULE_program              = 0
	EphaParserRULE_statement            = 1
	EphaParserRULE_importStatement      = 2
	EphaParserRULE_resourceDefinition   = 3
	EphaParserRULE_resourceBody         = 4
	EphaParserRULE_resourceProperty     = 5
	EphaParserRULE_resourcePropertyBody = 6
	EphaParserRULE_propertyKey          = 7
	EphaParserRULE_value                = 8
	EphaParserRULE_valueList            = 9
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EphaParserT__0 || _la == EphaParserIDENTIFIER {
		{
			p.SetState(20)
			p.Statement()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
	ImportStatement() IImportStatementContext
	ResourceDefinition() IResourceDefinitionContext

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

func (s *StatementContext) ImportStatement() IImportStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *StatementContext) ResourceDefinition() IResourceDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourceDefinitionContext)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EphaParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)
			p.ImportStatement()
		}

	case EphaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.ResourceDefinition()
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

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(EphaParserIDENTIFIER)
}

func (s *ImportStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, i)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *EphaParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EphaParserRULE_importStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(EphaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EphaParserT__1 {
		{
			p.SetState(31)
			p.Match(EphaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.Match(EphaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IResourceDefinitionContext is an interface to support dynamic dispatch.
type IResourceDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	ResourceBody() IResourceBodyContext

	// IsResourceDefinitionContext differentiates from other interfaces.
	IsResourceDefinitionContext()
}

type ResourceDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceDefinitionContext() *ResourceDefinitionContext {
	var p = new(ResourceDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_resourceDefinition
	return p
}

func InitEmptyResourceDefinitionContext(p *ResourceDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_resourceDefinition
}

func (*ResourceDefinitionContext) IsResourceDefinitionContext() {}

func NewResourceDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceDefinitionContext {
	var p = new(ResourceDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_resourceDefinition

	return p
}

func (s *ResourceDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceDefinitionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(EphaParserIDENTIFIER)
}

func (s *ResourceDefinitionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, i)
}

func (s *ResourceDefinitionContext) ResourceBody() IResourceBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourceBodyContext)
}

func (s *ResourceDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterResourceDefinition(s)
	}
}

func (s *ResourceDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitResourceDefinition(s)
	}
}

func (p *EphaParser) ResourceDefinition() (localctx IResourceDefinitionContext) {
	localctx = NewResourceDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EphaParserRULE_resourceDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.Match(EphaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.ResourceBody()
	}
	{
		p.SetState(39)
		p.Match(EphaParserT__3)
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

// IResourceBodyContext is an interface to support dynamic dispatch.
type IResourceBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllResourceProperty() []IResourcePropertyContext
	ResourceProperty(i int) IResourcePropertyContext

	// IsResourceBodyContext differentiates from other interfaces.
	IsResourceBodyContext()
}

type ResourceBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceBodyContext() *ResourceBodyContext {
	var p = new(ResourceBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_resourceBody
	return p
}

func InitEmptyResourceBodyContext(p *ResourceBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_resourceBody
}

func (*ResourceBodyContext) IsResourceBodyContext() {}

func NewResourceBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceBodyContext {
	var p = new(ResourceBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_resourceBody

	return p
}

func (s *ResourceBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceBodyContext) AllResourceProperty() []IResourcePropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResourcePropertyContext); ok {
			len++
		}
	}

	tst := make([]IResourcePropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResourcePropertyContext); ok {
			tst[i] = t.(IResourcePropertyContext)
			i++
		}
	}

	return tst
}

func (s *ResourceBodyContext) ResourceProperty(i int) IResourcePropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourcePropertyContext); ok {
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

	return t.(IResourcePropertyContext)
}

func (s *ResourceBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterResourceBody(s)
	}
}

func (s *ResourceBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitResourceBody(s)
	}
}

func (p *EphaParser) ResourceBody() (localctx IResourceBodyContext) {
	localctx = NewResourceBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EphaParserRULE_resourceBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserIDENTIFIER {
		{
			p.SetState(41)
			p.ResourceProperty()
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IResourcePropertyContext is an interface to support dynamic dispatch.
type IResourcePropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyKey() IPropertyKeyContext
	Value() IValueContext
	ValueList() IValueListContext
	ResourcePropertyBody() IResourcePropertyBodyContext

	// IsResourcePropertyContext differentiates from other interfaces.
	IsResourcePropertyContext()
}

type ResourcePropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourcePropertyContext() *ResourcePropertyContext {
	var p = new(ResourcePropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_resourceProperty
	return p
}

func InitEmptyResourcePropertyContext(p *ResourcePropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_resourceProperty
}

func (*ResourcePropertyContext) IsResourcePropertyContext() {}

func NewResourcePropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourcePropertyContext {
	var p = new(ResourcePropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_resourceProperty

	return p
}

func (s *ResourcePropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourcePropertyContext) PropertyKey() IPropertyKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyKeyContext)
}

func (s *ResourcePropertyContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ResourcePropertyContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *ResourcePropertyContext) ResourcePropertyBody() IResourcePropertyBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourcePropertyBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourcePropertyBodyContext)
}

func (s *ResourcePropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourcePropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourcePropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterResourceProperty(s)
	}
}

func (s *ResourcePropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitResourceProperty(s)
	}
}

func (p *EphaParser) ResourceProperty() (localctx IResourcePropertyContext) {
	localctx = NewResourcePropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EphaParserRULE_resourceProperty)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.PropertyKey()
		}
		{
			p.SetState(48)
			p.Match(EphaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.PropertyKey()
		}
		{
			p.SetState(52)
			p.Match(EphaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.ValueList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.PropertyKey()
		}
		{
			p.SetState(56)
			p.Match(EphaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.ResourcePropertyBody()
		}
		{
			p.SetState(58)
			p.Match(EphaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IResourcePropertyBodyContext is an interface to support dynamic dispatch.
type IResourcePropertyBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllResourceProperty() []IResourcePropertyContext
	ResourceProperty(i int) IResourcePropertyContext

	// IsResourcePropertyBodyContext differentiates from other interfaces.
	IsResourcePropertyBodyContext()
}

type ResourcePropertyBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourcePropertyBodyContext() *ResourcePropertyBodyContext {
	var p = new(ResourcePropertyBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_resourcePropertyBody
	return p
}

func InitEmptyResourcePropertyBodyContext(p *ResourcePropertyBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_resourcePropertyBody
}

func (*ResourcePropertyBodyContext) IsResourcePropertyBodyContext() {}

func NewResourcePropertyBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourcePropertyBodyContext {
	var p = new(ResourcePropertyBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_resourcePropertyBody

	return p
}

func (s *ResourcePropertyBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourcePropertyBodyContext) AllResourceProperty() []IResourcePropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResourcePropertyContext); ok {
			len++
		}
	}

	tst := make([]IResourcePropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResourcePropertyContext); ok {
			tst[i] = t.(IResourcePropertyContext)
			i++
		}
	}

	return tst
}

func (s *ResourcePropertyBodyContext) ResourceProperty(i int) IResourcePropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourcePropertyContext); ok {
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

	return t.(IResourcePropertyContext)
}

func (s *ResourcePropertyBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourcePropertyBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourcePropertyBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterResourcePropertyBody(s)
	}
}

func (s *ResourcePropertyBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitResourcePropertyBody(s)
	}
}

func (p *EphaParser) ResourcePropertyBody() (localctx IResourcePropertyBodyContext) {
	localctx = NewResourcePropertyBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EphaParserRULE_resourcePropertyBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserIDENTIFIER {
		{
			p.SetState(62)
			p.ResourceProperty()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IPropertyKeyContext is an interface to support dynamic dispatch.
type IPropertyKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsPropertyKeyContext differentiates from other interfaces.
	IsPropertyKeyContext()
}

type PropertyKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyKeyContext() *PropertyKeyContext {
	var p = new(PropertyKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_propertyKey
	return p
}

func InitEmptyPropertyKeyContext(p *PropertyKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_propertyKey
}

func (*PropertyKeyContext) IsPropertyKeyContext() {}

func NewPropertyKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyKeyContext {
	var p = new(PropertyKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_propertyKey

	return p
}

func (s *PropertyKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyKeyContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(EphaParserIDENTIFIER)
}

func (s *PropertyKeyContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(EphaParserIDENTIFIER, i)
}

func (s *PropertyKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterPropertyKey(s)
	}
}

func (s *PropertyKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitPropertyKey(s)
	}
}

func (p *EphaParser) PropertyKey() (localctx IPropertyKeyContext) {
	localctx = NewPropertyKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EphaParserRULE_propertyKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserT__5 {
		{
			p.SetState(69)
			p.Match(EphaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(EphaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(EphaParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(EphaParserNUMBER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *EphaParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EphaParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EphaParserSTRING || _la == EphaParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IValueListContext is an interface to support dynamic dispatch.
type IValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsValueListContext differentiates from other interfaces.
	IsValueListContext()
}

type ValueListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueListContext() *ValueListContext {
	var p = new(ValueListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_valueList
	return p
}

func InitEmptyValueListContext(p *ValueListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EphaParserRULE_valueList
}

func (*ValueListContext) IsValueListContext() {}

func NewValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueListContext {
	var p = new(ValueListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EphaParserRULE_valueList

	return p
}

func (s *ValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueListContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ValueListContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.EnterValueList(s)
	}
}

func (s *ValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EphaListener); ok {
		listenerT.ExitValueList(s)
	}
}

func (p *EphaParser) ValueList() (localctx IValueListContext) {
	localctx = NewValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, EphaParserRULE_valueList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(EphaParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Value()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserT__7 {
		{
			p.SetState(80)
			p.Match(EphaParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Value()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(EphaParserT__8)
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
