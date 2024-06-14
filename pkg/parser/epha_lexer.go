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
		"", "'{'", "'}'", "'='", "'['", "']'", "'.'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "IDENTIFIER", "STRING", "NUMBER", "LINE_COMMENT",
		"WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "IDENTIFIER",
		"STRING", "NUMBER", "LINE_COMMENT", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 86, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 5, 7, 42, 8, 7, 10, 7,
		12, 7, 45, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 51, 8, 8, 10, 8, 12, 8,
		54, 9, 8, 1, 8, 1, 8, 1, 9, 4, 9, 59, 8, 9, 11, 9, 12, 9, 60, 1, 9, 1,
		9, 4, 9, 65, 8, 9, 11, 9, 12, 9, 66, 3, 9, 69, 8, 9, 1, 10, 1, 10, 5, 10,
		73, 8, 10, 10, 10, 12, 10, 76, 9, 10, 1, 10, 1, 10, 1, 11, 4, 11, 81, 8,
		11, 11, 11, 12, 11, 82, 1, 11, 1, 11, 0, 0, 12, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 1, 0, 6, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 34,
		34, 92, 92, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32,
		32, 93, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 27, 1, 0, 0, 0, 5, 29, 1, 0, 0, 0,
		7, 31, 1, 0, 0, 0, 9, 33, 1, 0, 0, 0, 11, 35, 1, 0, 0, 0, 13, 37, 1, 0,
		0, 0, 15, 39, 1, 0, 0, 0, 17, 46, 1, 0, 0, 0, 19, 58, 1, 0, 0, 0, 21, 70,
		1, 0, 0, 0, 23, 80, 1, 0, 0, 0, 25, 26, 5, 123, 0, 0, 26, 2, 1, 0, 0, 0,
		27, 28, 5, 125, 0, 0, 28, 4, 1, 0, 0, 0, 29, 30, 5, 61, 0, 0, 30, 6, 1,
		0, 0, 0, 31, 32, 5, 91, 0, 0, 32, 8, 1, 0, 0, 0, 33, 34, 5, 93, 0, 0, 34,
		10, 1, 0, 0, 0, 35, 36, 5, 46, 0, 0, 36, 12, 1, 0, 0, 0, 37, 38, 5, 44,
		0, 0, 38, 14, 1, 0, 0, 0, 39, 43, 7, 0, 0, 0, 40, 42, 7, 1, 0, 0, 41, 40,
		1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0,
		44, 16, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 52, 5, 34, 0, 0, 47, 51, 8,
		2, 0, 0, 48, 49, 5, 92, 0, 0, 49, 51, 9, 0, 0, 0, 50, 47, 1, 0, 0, 0, 50,
		48, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0,
		0, 53, 55, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 56, 5, 34, 0, 0, 56, 18,
		1, 0, 0, 0, 57, 59, 7, 3, 0, 0, 58, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0,
		60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 68, 1, 0, 0, 0, 62, 64, 5,
		46, 0, 0, 63, 65, 7, 3, 0, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66,
		64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 69, 1, 0, 0, 0, 68, 62, 1, 0, 0,
		0, 68, 69, 1, 0, 0, 0, 69, 20, 1, 0, 0, 0, 70, 74, 5, 35, 0, 0, 71, 73,
		8, 4, 0, 0, 72, 71, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 78, 6,
		10, 0, 0, 78, 22, 1, 0, 0, 0, 79, 81, 7, 5, 0, 0, 80, 79, 1, 0, 0, 0, 81,
		82, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0,
		0, 84, 85, 6, 11, 0, 0, 85, 24, 1, 0, 0, 0, 9, 0, 43, 50, 52, 60, 66, 68,
		74, 82, 1, 6, 0, 0,
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
	EphaLexerIDENTIFIER   = 8
	EphaLexerSTRING       = 9
	EphaLexerNUMBER       = 10
	EphaLexerLINE_COMMENT = 11
	EphaLexerWHITESPACE   = 12
)
