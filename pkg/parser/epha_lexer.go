// Code generated from antlr/Epha.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 109,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 7, 13, 73, 10, 13, 12, 13, 14, 13, 76, 11, 13, 3, 14, 3, 14,
	7, 14, 80, 10, 14, 12, 14, 14, 14, 83, 11, 14, 3, 14, 3, 14, 3, 15, 6,
	15, 88, 10, 15, 13, 15, 14, 15, 89, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 101, 10, 16, 3, 17, 6, 17, 104, 10,
	17, 13, 17, 14, 17, 105, 3, 17, 3, 17, 3, 81, 2, 18, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 3, 2, 6, 5, 2, 67, 92, 97, 97, 99, 124, 6,
	2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 113, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 42, 3, 2, 2, 2, 7, 47, 3,
	2, 2, 2, 9, 49, 3, 2, 2, 2, 11, 51, 3, 2, 2, 2, 13, 56, 3, 2, 2, 2, 15,
	58, 3, 2, 2, 2, 17, 60, 3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 66, 3, 2, 2,
	2, 23, 68, 3, 2, 2, 2, 25, 70, 3, 2, 2, 2, 27, 77, 3, 2, 2, 2, 29, 87,
	3, 2, 2, 2, 31, 100, 3, 2, 2, 2, 33, 103, 3, 2, 2, 2, 35, 36, 7, 107, 2,
	2, 36, 37, 7, 111, 2, 2, 37, 38, 7, 114, 2, 2, 38, 39, 7, 113, 2, 2, 39,
	40, 7, 116, 2, 2, 40, 41, 7, 118, 2, 2, 41, 4, 3, 2, 2, 2, 42, 43, 7, 104,
	2, 2, 43, 44, 7, 116, 2, 2, 44, 45, 7, 113, 2, 2, 45, 46, 7, 111, 2, 2,
	46, 6, 3, 2, 2, 2, 47, 48, 7, 46, 2, 2, 48, 8, 3, 2, 2, 2, 49, 50, 7, 63,
	2, 2, 50, 10, 3, 2, 2, 2, 51, 52, 7, 106, 2, 2, 52, 53, 7, 103, 2, 2, 53,
	54, 7, 110, 2, 2, 54, 55, 7, 111, 2, 2, 55, 12, 3, 2, 2, 2, 56, 57, 7,
	125, 2, 2, 57, 14, 3, 2, 2, 2, 58, 59, 7, 127, 2, 2, 59, 16, 3, 2, 2, 2,
	60, 61, 7, 109, 2, 2, 61, 62, 7, 58, 2, 2, 62, 63, 7, 117, 2, 2, 63, 18,
	3, 2, 2, 2, 64, 65, 7, 60, 2, 2, 65, 20, 3, 2, 2, 2, 66, 67, 7, 93, 2,
	2, 67, 22, 3, 2, 2, 2, 68, 69, 7, 95, 2, 2, 69, 24, 3, 2, 2, 2, 70, 74,
	9, 2, 2, 2, 71, 73, 9, 3, 2, 2, 72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 26, 3, 2, 2, 2, 76, 74, 3,
	2, 2, 2, 77, 81, 7, 36, 2, 2, 78, 80, 11, 2, 2, 2, 79, 78, 3, 2, 2, 2,
	80, 83, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 84, 3,
	2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 7, 36, 2, 2, 85, 28, 3, 2, 2, 2, 86,
	88, 9, 4, 2, 2, 87, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 87, 3, 2, 2,
	2, 89, 90, 3, 2, 2, 2, 90, 30, 3, 2, 2, 2, 91, 92, 7, 118, 2, 2, 92, 93,
	7, 116, 2, 2, 93, 94, 7, 119, 2, 2, 94, 101, 7, 103, 2, 2, 95, 96, 7, 104,
	2, 2, 96, 97, 7, 99, 2, 2, 97, 98, 7, 110, 2, 2, 98, 99, 7, 117, 2, 2,
	99, 101, 7, 103, 2, 2, 100, 91, 3, 2, 2, 2, 100, 95, 3, 2, 2, 2, 101, 32,
	3, 2, 2, 2, 102, 104, 9, 5, 2, 2, 103, 102, 3, 2, 2, 2, 104, 105, 3, 2,
	2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2,
	107, 108, 8, 17, 2, 2, 108, 34, 3, 2, 2, 2, 8, 2, 74, 81, 89, 100, 105,
	3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'import'", "'from'", "','", "'='", "'helm'", "'{'", "'}'", "'k8s'",
	"':'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER", "STRING",
	"NUMBER", "BOOLEAN", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "IDENTIFIER", "STRING", "NUMBER", "BOOLEAN", "WS",
}

type EphaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewEphaLexer(input antlr.CharStream) *EphaLexer {

	l := new(EphaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
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
