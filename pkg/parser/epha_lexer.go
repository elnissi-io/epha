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
		"", "'import'", "'from'", "','", "'='", "'helm'", "'{'", "'}'", "'k8s'",
		"':'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER", "STRING",
		"NUMBER", "BOOLEAN", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "IDENTIFIER", "STRING", "NUMBER", "BOOLEAN", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 107, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 5, 11, 71, 8, 11, 10, 11, 12, 11, 74, 9, 11, 1, 12, 1, 12,
		5, 12, 78, 8, 12, 10, 12, 12, 12, 81, 9, 12, 1, 12, 1, 12, 1, 13, 4, 13,
		86, 8, 13, 11, 13, 12, 13, 87, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 99, 8, 14, 1, 15, 4, 15, 102, 8, 15, 11, 15,
		12, 15, 103, 1, 15, 1, 15, 1, 79, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 111, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 1, 33, 1, 0, 0, 0, 3, 40, 1, 0, 0, 0, 5, 45, 1, 0, 0, 0, 7, 47, 1, 0,
		0, 0, 9, 49, 1, 0, 0, 0, 11, 54, 1, 0, 0, 0, 13, 56, 1, 0, 0, 0, 15, 58,
		1, 0, 0, 0, 17, 62, 1, 0, 0, 0, 19, 64, 1, 0, 0, 0, 21, 66, 1, 0, 0, 0,
		23, 68, 1, 0, 0, 0, 25, 75, 1, 0, 0, 0, 27, 85, 1, 0, 0, 0, 29, 98, 1,
		0, 0, 0, 31, 101, 1, 0, 0, 0, 33, 34, 5, 105, 0, 0, 34, 35, 5, 109, 0,
		0, 35, 36, 5, 112, 0, 0, 36, 37, 5, 111, 0, 0, 37, 38, 5, 114, 0, 0, 38,
		39, 5, 116, 0, 0, 39, 2, 1, 0, 0, 0, 40, 41, 5, 102, 0, 0, 41, 42, 5, 114,
		0, 0, 42, 43, 5, 111, 0, 0, 43, 44, 5, 109, 0, 0, 44, 4, 1, 0, 0, 0, 45,
		46, 5, 44, 0, 0, 46, 6, 1, 0, 0, 0, 47, 48, 5, 61, 0, 0, 48, 8, 1, 0, 0,
		0, 49, 50, 5, 104, 0, 0, 50, 51, 5, 101, 0, 0, 51, 52, 5, 108, 0, 0, 52,
		53, 5, 109, 0, 0, 53, 10, 1, 0, 0, 0, 54, 55, 5, 123, 0, 0, 55, 12, 1,
		0, 0, 0, 56, 57, 5, 125, 0, 0, 57, 14, 1, 0, 0, 0, 58, 59, 5, 107, 0, 0,
		59, 60, 5, 56, 0, 0, 60, 61, 5, 115, 0, 0, 61, 16, 1, 0, 0, 0, 62, 63,
		5, 58, 0, 0, 63, 18, 1, 0, 0, 0, 64, 65, 5, 91, 0, 0, 65, 20, 1, 0, 0,
		0, 66, 67, 5, 93, 0, 0, 67, 22, 1, 0, 0, 0, 68, 72, 7, 0, 0, 0, 69, 71,
		7, 1, 0, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0,
		72, 73, 1, 0, 0, 0, 73, 24, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 79, 5,
		34, 0, 0, 76, 78, 9, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79,
		80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0,
		0, 82, 83, 5, 34, 0, 0, 83, 26, 1, 0, 0, 0, 84, 86, 7, 2, 0, 0, 85, 84,
		1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0,
		88, 28, 1, 0, 0, 0, 89, 90, 5, 116, 0, 0, 90, 91, 5, 114, 0, 0, 91, 92,
		5, 117, 0, 0, 92, 99, 5, 101, 0, 0, 93, 94, 5, 102, 0, 0, 94, 95, 5, 97,
		0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 115, 0, 0, 97, 99, 5, 101, 0, 0,
		98, 89, 1, 0, 0, 0, 98, 93, 1, 0, 0, 0, 99, 30, 1, 0, 0, 0, 100, 102, 7,
		3, 0, 0, 101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0,
		0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 6, 15, 0, 0, 106,
		32, 1, 0, 0, 0, 6, 0, 72, 79, 87, 98, 103, 1, 6, 0, 0,
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
	EphaLexerT__0       = 1
	EphaLexerT__1       = 2
	EphaLexerT__2       = 3
	EphaLexerT__3       = 4
	EphaLexerT__4       = 5
	EphaLexerT__5       = 6
	EphaLexerT__6       = 7
	EphaLexerT__7       = 8
	EphaLexerT__8       = 9
	EphaLexerT__9       = 10
	EphaLexerT__10      = 11
	EphaLexerIDENTIFIER = 12
	EphaLexerSTRING     = 13
	EphaLexerNUMBER     = 14
	EphaLexerBOOLEAN    = 15
	EphaLexerWS         = 16
)
