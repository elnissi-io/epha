// Code generated from parser/Epha.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type EphaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var EphaLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ephalexerLexerInit() {
	staticData := &EphaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'import'", "'as'", "'{'", "'}'", "'='", "'.'", "'['", "','", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "IDENTIFIER", "STRING", "NUMBER",
		"LINE_COMMENT", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"IDENTIFIER", "STRING", "NUMBER", "LINE_COMMENT", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 100, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9,
		56, 8, 9, 10, 9, 12, 9, 59, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 65,
		8, 10, 10, 10, 12, 10, 68, 9, 10, 1, 10, 1, 10, 1, 11, 4, 11, 73, 8, 11,
		11, 11, 12, 11, 74, 1, 11, 1, 11, 4, 11, 79, 8, 11, 11, 11, 12, 11, 80,
		3, 11, 83, 8, 11, 1, 12, 1, 12, 5, 12, 87, 8, 12, 10, 12, 12, 12, 90, 9,
		12, 1, 12, 1, 12, 1, 13, 4, 13, 95, 8, 13, 11, 13, 12, 13, 96, 1, 13, 1,
		13, 0, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 1, 0, 6, 3, 0, 65, 90, 95, 95,
		97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92,
		92, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 107,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 1, 29, 1, 0, 0, 0, 3, 36, 1,
		0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 41, 1, 0, 0, 0, 9, 43, 1, 0, 0, 0, 11, 45,
		1, 0, 0, 0, 13, 47, 1, 0, 0, 0, 15, 49, 1, 0, 0, 0, 17, 51, 1, 0, 0, 0,
		19, 53, 1, 0, 0, 0, 21, 60, 1, 0, 0, 0, 23, 72, 1, 0, 0, 0, 25, 84, 1,
		0, 0, 0, 27, 94, 1, 0, 0, 0, 29, 30, 5, 105, 0, 0, 30, 31, 5, 109, 0, 0,
		31, 32, 5, 112, 0, 0, 32, 33, 5, 111, 0, 0, 33, 34, 5, 114, 0, 0, 34, 35,
		5, 116, 0, 0, 35, 2, 1, 0, 0, 0, 36, 37, 5, 97, 0, 0, 37, 38, 5, 115, 0,
		0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 123, 0, 0, 40, 6, 1, 0, 0, 0, 41, 42,
		5, 125, 0, 0, 42, 8, 1, 0, 0, 0, 43, 44, 5, 61, 0, 0, 44, 10, 1, 0, 0,
		0, 45, 46, 5, 46, 0, 0, 46, 12, 1, 0, 0, 0, 47, 48, 5, 91, 0, 0, 48, 14,
		1, 0, 0, 0, 49, 50, 5, 44, 0, 0, 50, 16, 1, 0, 0, 0, 51, 52, 5, 93, 0,
		0, 52, 18, 1, 0, 0, 0, 53, 57, 7, 0, 0, 0, 54, 56, 7, 1, 0, 0, 55, 54,
		1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0,
		58, 20, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 66, 5, 34, 0, 0, 61, 65, 8,
		2, 0, 0, 62, 63, 5, 92, 0, 0, 63, 65, 9, 0, 0, 0, 64, 61, 1, 0, 0, 0, 64,
		62, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0,
		0, 67, 69, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 70, 5, 34, 0, 0, 70, 22,
		1, 0, 0, 0, 71, 73, 7, 3, 0, 0, 72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0,
		74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 82, 1, 0, 0, 0, 76, 78, 5,
		46, 0, 0, 77, 79, 7, 3, 0, 0, 78, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80,
		78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 76, 1, 0, 0,
		0, 82, 83, 1, 0, 0, 0, 83, 24, 1, 0, 0, 0, 84, 88, 5, 35, 0, 0, 85, 87,
		8, 4, 0, 0, 86, 85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0,
		88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 92, 6,
		12, 0, 0, 92, 26, 1, 0, 0, 0, 93, 95, 7, 5, 0, 0, 94, 93, 1, 0, 0, 0, 95,
		96, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0,
		0, 98, 99, 6, 13, 0, 0, 99, 28, 1, 0, 0, 0, 9, 0, 57, 64, 66, 74, 80, 82,
		88, 96, 1, 6, 0, 0,
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

// EphaLexerInit initializes any static state used to implement EphaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEphaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EphaLexerInit() {
	staticData := &EphaLexerLexerStaticData
	staticData.once.Do(ephalexerLexerInit)
}

// NewEphaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEphaLexer(input antlr.CharStream) *EphaLexer {
	EphaLexerInit()
	l := new(EphaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EphaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Epha.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EphaLexer tokens.
const (
	EphaLexerT__0         = 1
	EphaLexerT__1         = 2
	EphaLexerT__2         = 3
	EphaLexerT__3         = 4
	EphaLexerT__4         = 5
	EphaLexerT__5         = 6
	EphaLexerT__6         = 7
	EphaLexerT__7         = 8
	EphaLexerT__8         = 9
	EphaLexerIDENTIFIER   = 10
	EphaLexerSTRING       = 11
	EphaLexerNUMBER       = 12
	EphaLexerLINE_COMMENT = 13
	EphaLexerWHITESPACE   = 14
)
