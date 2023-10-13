// Code generated from /app/OpenFGALexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type OpenFGALexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var OpenFGALexerLexerStaticData struct {
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

func openfgalexerLexerInit() {
  staticData := &OpenFGALexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "", "'model'", "'schema'", "'1.1'", "'type'", "'condition'", "'relations'", 
    "'define'", "'with'", "'#'", "':'", "'*'", "'['", "']'", "'('", "')'", 
    "'{'", "'}'", "','", "", "", "'and'", "'or'", "'but not'", "'from'",
  }
  staticData.SymbolicNames = []string{
    "", "INDENT", "MODEL", "SCHEMA", "SCHEMA_VERSION", "TYPE", "CONDITION", 
    "RELATIONS", "DEFINE", "WTH", "HASH", "COLON", "WILDCARD", "L_SQUARE", 
    "R_SQUARE", "L_PARANTHESES", "R_PARANTHESES", "L_BRACES", "R_BRACES", 
    "COMMA", "CONDITION_PARAM_TYPE", "CONDITION_SYMBOL", "AND", "OR", "BUT_NOT", 
    "FROM", "ALPHA_NUMERIC", "NEWLINE", "WS",
  }
  staticData.RuleNames = []string{
    "SINGLE_INDENT", "DOUBLE_INDENT", "BOL", "INDENT", "MODEL", "SCHEMA", 
    "SCHEMA_VERSION", "TYPE", "CONDITION", "RELATIONS", "DEFINE", "WTH", 
    "HASH", "COLON", "WILDCARD", "L_SQUARE", "R_SQUARE", "L_PARANTHESES", 
    "R_PARANTHESES", "L_BRACES", "R_BRACES", "COMMA", "CONDITION_PARAM_TYPE", 
    "CONDITION_SYMBOL", "AND", "OR", "BUT_NOT", "FROM", "ALPHA_NUMERIC", 
    "NEWLINE", "WS",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 28, 291, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 
	2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 
	0, 1, 0, 1, 0, 3, 0, 67, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 
	75, 8, 1, 1, 2, 4, 2, 78, 8, 2, 11, 2, 12, 2, 79, 1, 3, 1, 3, 1, 3, 3, 
	3, 85, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 
	1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 
	10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 
	1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 
	18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 
	1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 
	22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 
	1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 
	22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 
	1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 
	22, 3, 22, 217, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 
	1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 
	23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 
	1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 
	23, 1, 23, 1, 23, 1, 23, 3, 23, 261, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 
	1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 
	26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 4, 28, 284, 8, 28, 11, 28, 
	12, 28, 285, 1, 29, 1, 29, 1, 30, 1, 30, 0, 0, 31, 1, 0, 3, 0, 5, 0, 7, 
	1, 9, 2, 11, 3, 13, 4, 15, 5, 17, 6, 19, 7, 21, 8, 23, 9, 25, 10, 27, 11, 
	29, 12, 31, 13, 33, 14, 35, 15, 37, 16, 39, 17, 41, 18, 43, 19, 45, 20, 
	47, 21, 49, 22, 51, 23, 53, 24, 55, 25, 57, 26, 59, 27, 61, 28, 1, 0, 6, 
	2, 0, 10, 10, 12, 13, 3, 0, 33, 33, 45, 45, 63, 63, 2, 0, 37, 37, 47, 47, 
	5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 2, 
	0, 9, 9, 32, 32, 325, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 
	0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 
	0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 
	1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 
	35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 
	0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 
	0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 
	0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 66, 1, 0, 0, 0, 3, 74, 1, 
	0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 81, 1, 0, 0, 0, 9, 86, 1, 0, 0, 0, 11, 92, 
	1, 0, 0, 0, 13, 99, 1, 0, 0, 0, 15, 103, 1, 0, 0, 0, 17, 108, 1, 0, 0, 
	0, 19, 118, 1, 0, 0, 0, 21, 128, 1, 0, 0, 0, 23, 135, 1, 0, 0, 0, 25, 140, 
	1, 0, 0, 0, 27, 142, 1, 0, 0, 0, 29, 144, 1, 0, 0, 0, 31, 146, 1, 0, 0, 
	0, 33, 148, 1, 0, 0, 0, 35, 150, 1, 0, 0, 0, 37, 152, 1, 0, 0, 0, 39, 154, 
	1, 0, 0, 0, 41, 156, 1, 0, 0, 0, 43, 158, 1, 0, 0, 0, 45, 216, 1, 0, 0, 
	0, 47, 260, 1, 0, 0, 0, 49, 262, 1, 0, 0, 0, 51, 266, 1, 0, 0, 0, 53, 269, 
	1, 0, 0, 0, 55, 277, 1, 0, 0, 0, 57, 283, 1, 0, 0, 0, 59, 287, 1, 0, 0, 
	0, 61, 289, 1, 0, 0, 0, 63, 64, 5, 32, 0, 0, 64, 67, 5, 32, 0, 0, 65, 67, 
	5, 9, 0, 0, 66, 63, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0, 67, 2, 1, 0, 0, 0, 
	68, 69, 5, 32, 0, 0, 69, 70, 5, 32, 0, 0, 70, 71, 5, 32, 0, 0, 71, 75, 
	5, 32, 0, 0, 72, 73, 5, 9, 0, 0, 73, 75, 5, 9, 0, 0, 74, 68, 1, 0, 0, 0, 
	74, 72, 1, 0, 0, 0, 75, 4, 1, 0, 0, 0, 76, 78, 7, 0, 0, 0, 77, 76, 1, 0, 
	0, 0, 78, 79, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 6, 
	1, 0, 0, 0, 81, 84, 3, 5, 2, 0, 82, 85, 3, 3, 1, 0, 83, 85, 3, 1, 0, 0, 
	84, 82, 1, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 8, 1, 0, 0, 0, 86, 87, 5, 109, 
	0, 0, 87, 88, 5, 111, 0, 0, 88, 89, 5, 100, 0, 0, 89, 90, 5, 101, 0, 0, 
	90, 91, 5, 108, 0, 0, 91, 10, 1, 0, 0, 0, 92, 93, 5, 115, 0, 0, 93, 94, 
	5, 99, 0, 0, 94, 95, 5, 104, 0, 0, 95, 96, 5, 101, 0, 0, 96, 97, 5, 109, 
	0, 0, 97, 98, 5, 97, 0, 0, 98, 12, 1, 0, 0, 0, 99, 100, 5, 49, 0, 0, 100, 
	101, 5, 46, 0, 0, 101, 102, 5, 49, 0, 0, 102, 14, 1, 0, 0, 0, 103, 104, 
	5, 116, 0, 0, 104, 105, 5, 121, 0, 0, 105, 106, 5, 112, 0, 0, 106, 107, 
	5, 101, 0, 0, 107, 16, 1, 0, 0, 0, 108, 109, 5, 99, 0, 0, 109, 110, 5, 
	111, 0, 0, 110, 111, 5, 110, 0, 0, 111, 112, 5, 100, 0, 0, 112, 113, 5, 
	105, 0, 0, 113, 114, 5, 116, 0, 0, 114, 115, 5, 105, 0, 0, 115, 116, 5, 
	111, 0, 0, 116, 117, 5, 110, 0, 0, 117, 18, 1, 0, 0, 0, 118, 119, 5, 114, 
	0, 0, 119, 120, 5, 101, 0, 0, 120, 121, 5, 108, 0, 0, 121, 122, 5, 97, 
	0, 0, 122, 123, 5, 116, 0, 0, 123, 124, 5, 105, 0, 0, 124, 125, 5, 111, 
	0, 0, 125, 126, 5, 110, 0, 0, 126, 127, 5, 115, 0, 0, 127, 20, 1, 0, 0, 
	0, 128, 129, 5, 100, 0, 0, 129, 130, 5, 101, 0, 0, 130, 131, 5, 102, 0, 
	0, 131, 132, 5, 105, 0, 0, 132, 133, 5, 110, 0, 0, 133, 134, 5, 101, 0, 
	0, 134, 22, 1, 0, 0, 0, 135, 136, 5, 119, 0, 0, 136, 137, 5, 105, 0, 0, 
	137, 138, 5, 116, 0, 0, 138, 139, 5, 104, 0, 0, 139, 24, 1, 0, 0, 0, 140, 
	141, 5, 35, 0, 0, 141, 26, 1, 0, 0, 0, 142, 143, 5, 58, 0, 0, 143, 28, 
	1, 0, 0, 0, 144, 145, 5, 42, 0, 0, 145, 30, 1, 0, 0, 0, 146, 147, 5, 91, 
	0, 0, 147, 32, 1, 0, 0, 0, 148, 149, 5, 93, 0, 0, 149, 34, 1, 0, 0, 0, 
	150, 151, 5, 40, 0, 0, 151, 36, 1, 0, 0, 0, 152, 153, 5, 41, 0, 0, 153, 
	38, 1, 0, 0, 0, 154, 155, 5, 123, 0, 0, 155, 40, 1, 0, 0, 0, 156, 157, 
	5, 125, 0, 0, 157, 42, 1, 0, 0, 0, 158, 159, 5, 44, 0, 0, 159, 44, 1, 0, 
	0, 0, 160, 161, 5, 98, 0, 0, 161, 162, 5, 111, 0, 0, 162, 163, 5, 111, 
	0, 0, 163, 217, 5, 108, 0, 0, 164, 165, 5, 115, 0, 0, 165, 166, 5, 116, 
	0, 0, 166, 167, 5, 114, 0, 0, 167, 168, 5, 105, 0, 0, 168, 169, 5, 110, 
	0, 0, 169, 217, 5, 103, 0, 0, 170, 171, 5, 105, 0, 0, 171, 172, 5, 110, 
	0, 0, 172, 217, 5, 116, 0, 0, 173, 174, 5, 117, 0, 0, 174, 175, 5, 105, 
	0, 0, 175, 176, 5, 110, 0, 0, 176, 217, 5, 116, 0, 0, 177, 178, 5, 100, 
	0, 0, 178, 179, 5, 111, 0, 0, 179, 180, 5, 117, 0, 0, 180, 181, 5, 98, 
	0, 0, 181, 182, 5, 108, 0, 0, 182, 217, 5, 101, 0, 0, 183, 184, 5, 100, 
	0, 0, 184, 185, 5, 117, 0, 0, 185, 186, 5, 114, 0, 0, 186, 187, 5, 97, 
	0, 0, 187, 188, 5, 116, 0, 0, 188, 189, 5, 105, 0, 0, 189, 190, 5, 111, 
	0, 0, 190, 217, 5, 110, 0, 0, 191, 192, 5, 116, 0, 0, 192, 193, 5, 105, 
	0, 0, 193, 194, 5, 109, 0, 0, 194, 195, 5, 101, 0, 0, 195, 196, 5, 115, 
	0, 0, 196, 197, 5, 116, 0, 0, 197, 198, 5, 97, 0, 0, 198, 199, 5, 109, 
	0, 0, 199, 217, 5, 112, 0, 0, 200, 201, 5, 109, 0, 0, 201, 202, 5, 97, 
	0, 0, 202, 217, 5, 112, 0, 0, 203, 204, 5, 108, 0, 0, 204, 205, 5, 105, 
	0, 0, 205, 206, 5, 115, 0, 0, 206, 217, 5, 116, 0, 0, 207, 208, 5, 105, 
	0, 0, 208, 209, 5, 112, 0, 0, 209, 210, 5, 97, 0, 0, 210, 211, 5, 100, 
	0, 0, 211, 212, 5, 100, 0, 0, 212, 213, 5, 114, 0, 0, 213, 214, 5, 101, 
	0, 0, 214, 215, 5, 115, 0, 0, 215, 217, 5, 115, 0, 0, 216, 160, 1, 0, 0, 
	0, 216, 164, 1, 0, 0, 0, 216, 170, 1, 0, 0, 0, 216, 173, 1, 0, 0, 0, 216, 
	177, 1, 0, 0, 0, 216, 183, 1, 0, 0, 0, 216, 191, 1, 0, 0, 0, 216, 200, 
	1, 0, 0, 0, 216, 203, 1, 0, 0, 0, 216, 207, 1, 0, 0, 0, 217, 46, 1, 0, 
	0, 0, 218, 219, 5, 61, 0, 0, 219, 261, 5, 61, 0, 0, 220, 221, 5, 33, 0, 
	0, 221, 261, 5, 61, 0, 0, 222, 223, 5, 105, 0, 0, 223, 261, 5, 110, 0, 
	0, 224, 261, 5, 60, 0, 0, 225, 226, 5, 60, 0, 0, 226, 261, 5, 61, 0, 0, 
	227, 228, 5, 62, 0, 0, 228, 261, 5, 61, 0, 0, 229, 261, 5, 62, 0, 0, 230, 
	231, 5, 38, 0, 0, 231, 261, 5, 38, 0, 0, 232, 233, 5, 124, 0, 0, 233, 261, 
	5, 124, 0, 0, 234, 261, 3, 31, 15, 0, 235, 261, 3, 33, 16, 0, 236, 261, 
	3, 39, 19, 0, 237, 261, 3, 35, 17, 0, 238, 261, 3, 37, 18, 0, 239, 261, 
	5, 46, 0, 0, 240, 261, 3, 43, 21, 0, 241, 261, 7, 1, 0, 0, 242, 261, 3, 
	27, 13, 0, 243, 261, 5, 43, 0, 0, 244, 261, 3, 29, 14, 0, 245, 261, 7, 
	2, 0, 0, 246, 247, 5, 116, 0, 0, 247, 248, 5, 114, 0, 0, 248, 249, 5, 117, 
	0, 0, 249, 261, 5, 101, 0, 0, 250, 251, 5, 102, 0, 0, 251, 252, 5, 97, 
	0, 0, 252, 253, 5, 108, 0, 0, 253, 254, 5, 115, 0, 0, 254, 261, 5, 101, 
	0, 0, 255, 256, 5, 110, 0, 0, 256, 257, 5, 117, 0, 0, 257, 258, 5, 108, 
	0, 0, 258, 261, 5, 108, 0, 0, 259, 261, 5, 34, 0, 0, 260, 218, 1, 0, 0, 
	0, 260, 220, 1, 0, 0, 0, 260, 222, 1, 0, 0, 0, 260, 224, 1, 0, 0, 0, 260, 
	225, 1, 0, 0, 0, 260, 227, 1, 0, 0, 0, 260, 229, 1, 0, 0, 0, 260, 230, 
	1, 0, 0, 0, 260, 232, 1, 0, 0, 0, 260, 234, 1, 0, 0, 0, 260, 235, 1, 0, 
	0, 0, 260, 236, 1, 0, 0, 0, 260, 237, 1, 0, 0, 0, 260, 238, 1, 0, 0, 0, 
	260, 239, 1, 0, 0, 0, 260, 240, 1, 0, 0, 0, 260, 241, 1, 0, 0, 0, 260, 
	242, 1, 0, 0, 0, 260, 243, 1, 0, 0, 0, 260, 244, 1, 0, 0, 0, 260, 245, 
	1, 0, 0, 0, 260, 246, 1, 0, 0, 0, 260, 250, 1, 0, 0, 0, 260, 255, 1, 0, 
	0, 0, 260, 259, 1, 0, 0, 0, 261, 48, 1, 0, 0, 0, 262, 263, 5, 97, 0, 0, 
	263, 264, 5, 110, 0, 0, 264, 265, 5, 100, 0, 0, 265, 50, 1, 0, 0, 0, 266, 
	267, 5, 111, 0, 0, 267, 268, 5, 114, 0, 0, 268, 52, 1, 0, 0, 0, 269, 270, 
	5, 98, 0, 0, 270, 271, 5, 117, 0, 0, 271, 272, 5, 116, 0, 0, 272, 273, 
	5, 32, 0, 0, 273, 274, 5, 110, 0, 0, 274, 275, 5, 111, 0, 0, 275, 276, 
	5, 116, 0, 0, 276, 54, 1, 0, 0, 0, 277, 278, 5, 102, 0, 0, 278, 279, 5, 
	114, 0, 0, 279, 280, 5, 111, 0, 0, 280, 281, 5, 109, 0, 0, 281, 56, 1, 
	0, 0, 0, 282, 284, 7, 3, 0, 0, 283, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 
	0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 58, 1, 0, 0, 0, 287, 
	288, 7, 4, 0, 0, 288, 60, 1, 0, 0, 0, 289, 290, 7, 5, 0, 0, 290, 62, 1, 
	0, 0, 0, 8, 0, 66, 74, 79, 84, 216, 260, 285, 0,
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

// OpenFGALexerInit initializes any static state used to implement OpenFGALexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewOpenFGALexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func OpenFGALexerInit() {
  staticData := &OpenFGALexerLexerStaticData
  staticData.once.Do(openfgalexerLexerInit)
}

// NewOpenFGALexer produces a new lexer instance for the optional input antlr.CharStream.
func NewOpenFGALexer(input antlr.CharStream) *OpenFGALexer {
  OpenFGALexerInit()
	l := new(OpenFGALexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &OpenFGALexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "OpenFGALexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// OpenFGALexer tokens.
const (
	OpenFGALexerINDENT = 1
	OpenFGALexerMODEL = 2
	OpenFGALexerSCHEMA = 3
	OpenFGALexerSCHEMA_VERSION = 4
	OpenFGALexerTYPE = 5
	OpenFGALexerCONDITION = 6
	OpenFGALexerRELATIONS = 7
	OpenFGALexerDEFINE = 8
	OpenFGALexerWTH = 9
	OpenFGALexerHASH = 10
	OpenFGALexerCOLON = 11
	OpenFGALexerWILDCARD = 12
	OpenFGALexerL_SQUARE = 13
	OpenFGALexerR_SQUARE = 14
	OpenFGALexerL_PARANTHESES = 15
	OpenFGALexerR_PARANTHESES = 16
	OpenFGALexerL_BRACES = 17
	OpenFGALexerR_BRACES = 18
	OpenFGALexerCOMMA = 19
	OpenFGALexerCONDITION_PARAM_TYPE = 20
	OpenFGALexerCONDITION_SYMBOL = 21
	OpenFGALexerAND = 22
	OpenFGALexerOR = 23
	OpenFGALexerBUT_NOT = 24
	OpenFGALexerFROM = 25
	OpenFGALexerALPHA_NUMERIC = 26
	OpenFGALexerNEWLINE = 27
	OpenFGALexerWS = 28
)

