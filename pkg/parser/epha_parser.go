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
		"", "'{'", "'}'", "'='", "'['", "']'", "'.'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "IDENTIFIER", "STRING", "NUMBER", "LINE_COMMENT",
		"WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "statement", "resourceDefinition", "resourceBody", "resourceProperty",
		"resourcePropertyBody", "propertyKey", "value", "valueList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 78, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 4, 0, 20, 8, 0,
		11, 0, 12, 0, 21, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		5, 3, 33, 8, 3, 10, 3, 12, 3, 36, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52, 8, 4, 1, 5,
		5, 5, 55, 8, 5, 10, 5, 12, 5, 58, 9, 5, 1, 6, 1, 6, 1, 6, 5, 6, 63, 8,
		6, 10, 6, 12, 6, 66, 9, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 73, 8, 8,
		10, 8, 12, 8, 76, 9, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0,
		1, 1, 0, 9, 10, 75, 0, 19, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 25, 1, 0,
		0, 0, 6, 34, 1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 59,
		1, 0, 0, 0, 14, 67, 1, 0, 0, 0, 16, 69, 1, 0, 0, 0, 18, 20, 3, 2, 1, 0,
		19, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1,
		0, 0, 0, 22, 1, 1, 0, 0, 0, 23, 24, 3, 4, 2, 0, 24, 3, 1, 0, 0, 0, 25,
		26, 5, 8, 0, 0, 26, 27, 5, 8, 0, 0, 27, 28, 5, 1, 0, 0, 28, 29, 3, 6, 3,
		0, 29, 30, 5, 2, 0, 0, 30, 5, 1, 0, 0, 0, 31, 33, 3, 8, 4, 0, 32, 31, 1,
		0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35,
		7, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 38, 3, 12, 6, 0, 38, 39, 5, 3, 0,
		0, 39, 40, 3, 14, 7, 0, 40, 52, 1, 0, 0, 0, 41, 42, 3, 12, 6, 0, 42, 43,
		5, 4, 0, 0, 43, 44, 3, 16, 8, 0, 44, 45, 5, 5, 0, 0, 45, 52, 1, 0, 0, 0,
		46, 47, 3, 12, 6, 0, 47, 48, 5, 1, 0, 0, 48, 49, 3, 10, 5, 0, 49, 50, 5,
		2, 0, 0, 50, 52, 1, 0, 0, 0, 51, 37, 1, 0, 0, 0, 51, 41, 1, 0, 0, 0, 51,
		46, 1, 0, 0, 0, 52, 9, 1, 0, 0, 0, 53, 55, 3, 8, 4, 0, 54, 53, 1, 0, 0,
		0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 11,
		1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 64, 5, 8, 0, 0, 60, 61, 5, 6, 0, 0,
		61, 63, 5, 8, 0, 0, 62, 60, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1,
		0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 13, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67,
		68, 7, 0, 0, 0, 68, 15, 1, 0, 0, 0, 69, 74, 3, 14, 7, 0, 70, 71, 5, 7,
		0, 0, 71, 73, 3, 14, 7, 0, 72, 70, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74,
		72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 17, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 6, 21, 34, 51, 56, 64, 74,
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
	EphaParserIDENTIFIER   = 8
	EphaParserSTRING       = 9
	EphaParserNUMBER       = 10
	EphaParserLINE_COMMENT = 11
	EphaParserWHITESPACE   = 12
)

// EphaParser rules.
const (
	EphaParserRULE_program              = 0
	EphaParserRULE_statement            = 1
	EphaParserRULE_resourceDefinition   = 2
	EphaParserRULE_resourceBody         = 3
	EphaParserRULE_resourceProperty     = 4
	EphaParserRULE_resourcePropertyBody = 5
	EphaParserRULE_propertyKey          = 6
	EphaParserRULE_value                = 7
	EphaParserRULE_valueList            = 8
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EphaParserIDENTIFIER {
		{
			p.SetState(18)
			p.Statement()
		}

		p.SetState(21)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.ResourceDefinition()
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
	p.EnterRule(localctx, 4, EphaParserRULE_resourceDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(EphaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.ResourceBody()
	}
	{
		p.SetState(29)
		p.Match(EphaParserT__1)
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
	p.EnterRule(localctx, 6, EphaParserRULE_resourceBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserIDENTIFIER {
		{
			p.SetState(31)
			p.ResourceProperty()
		}

		p.SetState(36)
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
	p.EnterRule(localctx, 8, EphaParserRULE_resourceProperty)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.PropertyKey()
		}
		{
			p.SetState(38)
			p.Match(EphaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.PropertyKey()
		}
		{
			p.SetState(42)
			p.Match(EphaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.ValueList()
		}
		{
			p.SetState(44)
			p.Match(EphaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.PropertyKey()
		}
		{
			p.SetState(47)
			p.Match(EphaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.ResourcePropertyBody()
		}
		{
			p.SetState(49)
			p.Match(EphaParserT__1)
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
	p.EnterRule(localctx, 10, EphaParserRULE_resourcePropertyBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserIDENTIFIER {
		{
			p.SetState(53)
			p.ResourceProperty()
		}

		p.SetState(58)
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
	p.EnterRule(localctx, 12, EphaParserRULE_propertyKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(EphaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserT__5 {
		{
			p.SetState(60)
			p.Match(EphaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(EphaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(66)
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
	p.EnterRule(localctx, 14, EphaParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
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
	p.EnterRule(localctx, 16, EphaParserRULE_valueList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Value()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EphaParserT__6 {
		{
			p.SetState(70)
			p.Match(EphaParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Value()
		}

		p.SetState(76)
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
