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
    "'{'", "'}'", "'<'", "'>'", "','", "", "", "", "'and'", "'or'", "'but not'", 
    "'from'",
  }
  staticData.SymbolicNames = []string{
    "", "INDENT", "MODEL", "SCHEMA", "SCHEMA_VERSION", "TYPE", "CONDITION", 
    "RELATIONS", "DEFINE", "WTH", "HASH", "COLON", "WILDCARD", "L_SQUARE", 
    "R_SQUARE", "L_PARANTHESES", "R_PARANTHESES", "L_BRACES", "R_BRACES", 
    "L_ANGLE_BRACKET", "R_ANGLE_BRACKET", "COMMA", "CONDITION_PARAM_CONTAINER", 
    "CONDITION_PARAM_TYPE", "CONDITION_SYMBOL", "AND", "OR", "BUT_NOT", 
    "FROM", "ALPHA_NUMERIC", "NEWLINE", "WS",
  }
  staticData.RuleNames = []string{
    "SINGLE_INDENT", "DOUBLE_INDENT", "BOL", "INDENT", "MODEL", "SCHEMA", 
    "SCHEMA_VERSION", "TYPE", "CONDITION", "RELATIONS", "DEFINE", "WTH", 
    "HASH", "COLON", "WILDCARD", "L_SQUARE", "R_SQUARE", "L_PARANTHESES", 
    "R_PARANTHESES", "L_BRACES", "R_BRACES", "L_ANGLE_BRACKET", "R_ANGLE_BRACKET", 
    "COMMA", "CONDITION_PARAM_CONTAINER", "CONDITION_PARAM_TYPE", "CONDITION_SYMBOL", 
    "AND", "OR", "BUT_NOT", "FROM", "ALPHA_NUMERIC", "NEWLINE", "WS",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 31, 303, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 
	2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 
	31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 3, 0, 73, 8, 0, 
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 81, 8, 1, 1, 2, 4, 2, 84, 8, 
	2, 11, 2, 12, 2, 85, 1, 3, 1, 3, 1, 3, 3, 3, 91, 8, 3, 1, 4, 1, 4, 1, 4, 
	1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 
	1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 
	1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 
	1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 
	15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 
	1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 
	24, 1, 24, 1, 24, 1, 24, 3, 24, 178, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 
	1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 
	25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 
	1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 
	25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 
	1, 25, 1, 25, 1, 25, 3, 25, 229, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 
	26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 
	1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 
	26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 
	1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 273, 8, 26, 1, 27, 1, 
	27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 
	1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 4, 31, 296, 
	8, 31, 11, 31, 12, 31, 297, 1, 32, 1, 32, 1, 33, 1, 33, 0, 0, 34, 1, 0, 
	3, 0, 5, 0, 7, 1, 9, 2, 11, 3, 13, 4, 15, 5, 17, 6, 19, 7, 21, 8, 23, 9, 
	25, 10, 27, 11, 29, 12, 31, 13, 33, 14, 35, 15, 37, 16, 39, 17, 41, 18, 
	43, 19, 45, 20, 47, 21, 49, 22, 51, 23, 53, 24, 55, 25, 57, 26, 59, 27, 
	61, 28, 63, 29, 65, 30, 67, 31, 1, 0, 6, 2, 0, 10, 10, 12, 13, 3, 0, 33, 
	33, 45, 45, 63, 63, 2, 0, 37, 37, 47, 47, 5, 0, 45, 45, 48, 57, 65, 90, 
	95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 336, 0, 7, 1, 
	0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 
	1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 
	23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 
	0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 
	0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 
	0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 
	0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 
	1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 1, 
	72, 1, 0, 0, 0, 3, 80, 1, 0, 0, 0, 5, 83, 1, 0, 0, 0, 7, 87, 1, 0, 0, 0, 
	9, 92, 1, 0, 0, 0, 11, 98, 1, 0, 0, 0, 13, 105, 1, 0, 0, 0, 15, 109, 1, 
	0, 0, 0, 17, 114, 1, 0, 0, 0, 19, 124, 1, 0, 0, 0, 21, 134, 1, 0, 0, 0, 
	23, 141, 1, 0, 0, 0, 25, 146, 1, 0, 0, 0, 27, 148, 1, 0, 0, 0, 29, 150, 
	1, 0, 0, 0, 31, 152, 1, 0, 0, 0, 33, 154, 1, 0, 0, 0, 35, 156, 1, 0, 0, 
	0, 37, 158, 1, 0, 0, 0, 39, 160, 1, 0, 0, 0, 41, 162, 1, 0, 0, 0, 43, 164, 
	1, 0, 0, 0, 45, 166, 1, 0, 0, 0, 47, 168, 1, 0, 0, 0, 49, 177, 1, 0, 0, 
	0, 51, 228, 1, 0, 0, 0, 53, 272, 1, 0, 0, 0, 55, 274, 1, 0, 0, 0, 57, 278, 
	1, 0, 0, 0, 59, 281, 1, 0, 0, 0, 61, 289, 1, 0, 0, 0, 63, 295, 1, 0, 0, 
	0, 65, 299, 1, 0, 0, 0, 67, 301, 1, 0, 0, 0, 69, 70, 5, 32, 0, 0, 70, 73, 
	5, 32, 0, 0, 71, 73, 5, 9, 0, 0, 72, 69, 1, 0, 0, 0, 72, 71, 1, 0, 0, 0, 
	73, 2, 1, 0, 0, 0, 74, 75, 5, 32, 0, 0, 75, 76, 5, 32, 0, 0, 76, 77, 5, 
	32, 0, 0, 77, 81, 5, 32, 0, 0, 78, 79, 5, 9, 0, 0, 79, 81, 5, 9, 0, 0, 
	80, 74, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 4, 1, 0, 0, 0, 82, 84, 7, 0, 
	0, 0, 83, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 
	1, 0, 0, 0, 86, 6, 1, 0, 0, 0, 87, 90, 3, 5, 2, 0, 88, 91, 3, 3, 1, 0, 
	89, 91, 3, 1, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 8, 1, 0, 
	0, 0, 92, 93, 5, 109, 0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 100, 0, 0, 
	95, 96, 5, 101, 0, 0, 96, 97, 5, 108, 0, 0, 97, 10, 1, 0, 0, 0, 98, 99, 
	5, 115, 0, 0, 99, 100, 5, 99, 0, 0, 100, 101, 5, 104, 0, 0, 101, 102, 5, 
	101, 0, 0, 102, 103, 5, 109, 0, 0, 103, 104, 5, 97, 0, 0, 104, 12, 1, 0, 
	0, 0, 105, 106, 5, 49, 0, 0, 106, 107, 5, 46, 0, 0, 107, 108, 5, 49, 0, 
	0, 108, 14, 1, 0, 0, 0, 109, 110, 5, 116, 0, 0, 110, 111, 5, 121, 0, 0, 
	111, 112, 5, 112, 0, 0, 112, 113, 5, 101, 0, 0, 113, 16, 1, 0, 0, 0, 114, 
	115, 5, 99, 0, 0, 115, 116, 5, 111, 0, 0, 116, 117, 5, 110, 0, 0, 117, 
	118, 5, 100, 0, 0, 118, 119, 5, 105, 0, 0, 119, 120, 5, 116, 0, 0, 120, 
	121, 5, 105, 0, 0, 121, 122, 5, 111, 0, 0, 122, 123, 5, 110, 0, 0, 123, 
	18, 1, 0, 0, 0, 124, 125, 5, 114, 0, 0, 125, 126, 5, 101, 0, 0, 126, 127, 
	5, 108, 0, 0, 127, 128, 5, 97, 0, 0, 128, 129, 5, 116, 0, 0, 129, 130, 
	5, 105, 0, 0, 130, 131, 5, 111, 0, 0, 131, 132, 5, 110, 0, 0, 132, 133, 
	5, 115, 0, 0, 133, 20, 1, 0, 0, 0, 134, 135, 5, 100, 0, 0, 135, 136, 5, 
	101, 0, 0, 136, 137, 5, 102, 0, 0, 137, 138, 5, 105, 0, 0, 138, 139, 5, 
	110, 0, 0, 139, 140, 5, 101, 0, 0, 140, 22, 1, 0, 0, 0, 141, 142, 5, 119, 
	0, 0, 142, 143, 5, 105, 0, 0, 143, 144, 5, 116, 0, 0, 144, 145, 5, 104, 
	0, 0, 145, 24, 1, 0, 0, 0, 146, 147, 5, 35, 0, 0, 147, 26, 1, 0, 0, 0, 
	148, 149, 5, 58, 0, 0, 149, 28, 1, 0, 0, 0, 150, 151, 5, 42, 0, 0, 151, 
	30, 1, 0, 0, 0, 152, 153, 5, 91, 0, 0, 153, 32, 1, 0, 0, 0, 154, 155, 5, 
	93, 0, 0, 155, 34, 1, 0, 0, 0, 156, 157, 5, 40, 0, 0, 157, 36, 1, 0, 0, 
	0, 158, 159, 5, 41, 0, 0, 159, 38, 1, 0, 0, 0, 160, 161, 5, 123, 0, 0, 
	161, 40, 1, 0, 0, 0, 162, 163, 5, 125, 0, 0, 163, 42, 1, 0, 0, 0, 164, 
	165, 5, 60, 0, 0, 165, 44, 1, 0, 0, 0, 166, 167, 5, 62, 0, 0, 167, 46, 
	1, 0, 0, 0, 168, 169, 5, 44, 0, 0, 169, 48, 1, 0, 0, 0, 170, 171, 5, 109, 
	0, 0, 171, 172, 5, 97, 0, 0, 172, 178, 5, 112, 0, 0, 173, 174, 5, 108, 
	0, 0, 174, 175, 5, 105, 0, 0, 175, 176, 5, 115, 0, 0, 176, 178, 5, 116, 
	0, 0, 177, 170, 1, 0, 0, 0, 177, 173, 1, 0, 0, 0, 178, 50, 1, 0, 0, 0, 
	179, 180, 5, 98, 0, 0, 180, 181, 5, 111, 0, 0, 181, 182, 5, 111, 0, 0, 
	182, 229, 5, 108, 0, 0, 183, 184, 5, 115, 0, 0, 184, 185, 5, 116, 0, 0, 
	185, 186, 5, 114, 0, 0, 186, 187, 5, 105, 0, 0, 187, 188, 5, 110, 0, 0, 
	188, 229, 5, 103, 0, 0, 189, 190, 5, 105, 0, 0, 190, 191, 5, 110, 0, 0, 
	191, 229, 5, 116, 0, 0, 192, 193, 5, 117, 0, 0, 193, 194, 5, 105, 0, 0, 
	194, 195, 5, 110, 0, 0, 195, 229, 5, 116, 0, 0, 196, 197, 5, 100, 0, 0, 
	197, 198, 5, 111, 0, 0, 198, 199, 5, 117, 0, 0, 199, 200, 5, 98, 0, 0, 
	200, 201, 5, 108, 0, 0, 201, 229, 5, 101, 0, 0, 202, 203, 5, 100, 0, 0, 
	203, 204, 5, 117, 0, 0, 204, 205, 5, 114, 0, 0, 205, 206, 5, 97, 0, 0, 
	206, 207, 5, 116, 0, 0, 207, 208, 5, 105, 0, 0, 208, 209, 5, 111, 0, 0, 
	209, 229, 5, 110, 0, 0, 210, 211, 5, 116, 0, 0, 211, 212, 5, 105, 0, 0, 
	212, 213, 5, 109, 0, 0, 213, 214, 5, 101, 0, 0, 214, 215, 5, 115, 0, 0, 
	215, 216, 5, 116, 0, 0, 216, 217, 5, 97, 0, 0, 217, 218, 5, 109, 0, 0, 
	218, 229, 5, 112, 0, 0, 219, 220, 5, 105, 0, 0, 220, 221, 5, 112, 0, 0, 
	221, 222, 5, 97, 0, 0, 222, 223, 5, 100, 0, 0, 223, 224, 5, 100, 0, 0, 
	224, 225, 5, 114, 0, 0, 225, 226, 5, 101, 0, 0, 226, 227, 5, 115, 0, 0, 
	227, 229, 5, 115, 0, 0, 228, 179, 1, 0, 0, 0, 228, 183, 1, 0, 0, 0, 228, 
	189, 1, 0, 0, 0, 228, 192, 1, 0, 0, 0, 228, 196, 1, 0, 0, 0, 228, 202, 
	1, 0, 0, 0, 228, 210, 1, 0, 0, 0, 228, 219, 1, 0, 0, 0, 229, 52, 1, 0, 
	0, 0, 230, 231, 5, 61, 0, 0, 231, 273, 5, 61, 0, 0, 232, 233, 5, 33, 0, 
	0, 233, 273, 5, 61, 0, 0, 234, 235, 5, 105, 0, 0, 235, 273, 5, 110, 0, 
	0, 236, 273, 3, 43, 21, 0, 237, 238, 5, 60, 0, 0, 238, 273, 5, 61, 0, 0, 
	239, 240, 5, 62, 0, 0, 240, 273, 5, 61, 0, 0, 241, 273, 3, 45, 22, 0, 242, 
	243, 5, 38, 0, 0, 243, 273, 5, 38, 0, 0, 244, 245, 5, 124, 0, 0, 245, 273, 
	5, 124, 0, 0, 246, 273, 3, 31, 15, 0, 247, 273, 3, 33, 16, 0, 248, 273, 
	3, 39, 19, 0, 249, 273, 3, 35, 17, 0, 250, 273, 3, 37, 18, 0, 251, 273, 
	5, 46, 0, 0, 252, 273, 3, 47, 23, 0, 253, 273, 7, 1, 0, 0, 254, 273, 3, 
	27, 13, 0, 255, 273, 5, 43, 0, 0, 256, 273, 3, 29, 14, 0, 257, 273, 7, 
	2, 0, 0, 258, 259, 5, 116, 0, 0, 259, 260, 5, 114, 0, 0, 260, 261, 5, 117, 
	0, 0, 261, 273, 5, 101, 0, 0, 262, 263, 5, 102, 0, 0, 263, 264, 5, 97, 
	0, 0, 264, 265, 5, 108, 0, 0, 265, 266, 5, 115, 0, 0, 266, 273, 5, 101, 
	0, 0, 267, 268, 5, 110, 0, 0, 268, 269, 5, 117, 0, 0, 269, 270, 5, 108, 
	0, 0, 270, 273, 5, 108, 0, 0, 271, 273, 5, 34, 0, 0, 272, 230, 1, 0, 0, 
	0, 272, 232, 1, 0, 0, 0, 272, 234, 1, 0, 0, 0, 272, 236, 1, 0, 0, 0, 272, 
	237, 1, 0, 0, 0, 272, 239, 1, 0, 0, 0, 272, 241, 1, 0, 0, 0, 272, 242, 
	1, 0, 0, 0, 272, 244, 1, 0, 0, 0, 272, 246, 1, 0, 0, 0, 272, 247, 1, 0, 
	0, 0, 272, 248, 1, 0, 0, 0, 272, 249, 1, 0, 0, 0, 272, 250, 1, 0, 0, 0, 
	272, 251, 1, 0, 0, 0, 272, 252, 1, 0, 0, 0, 272, 253, 1, 0, 0, 0, 272, 
	254, 1, 0, 0, 0, 272, 255, 1, 0, 0, 0, 272, 256, 1, 0, 0, 0, 272, 257, 
	1, 0, 0, 0, 272, 258, 1, 0, 0, 0, 272, 262, 1, 0, 0, 0, 272, 267, 1, 0, 
	0, 0, 272, 271, 1, 0, 0, 0, 273, 54, 1, 0, 0, 0, 274, 275, 5, 97, 0, 0, 
	275, 276, 5, 110, 0, 0, 276, 277, 5, 100, 0, 0, 277, 56, 1, 0, 0, 0, 278, 
	279, 5, 111, 0, 0, 279, 280, 5, 114, 0, 0, 280, 58, 1, 0, 0, 0, 281, 282, 
	5, 98, 0, 0, 282, 283, 5, 117, 0, 0, 283, 284, 5, 116, 0, 0, 284, 285, 
	5, 32, 0, 0, 285, 286, 5, 110, 0, 0, 286, 287, 5, 111, 0, 0, 287, 288, 
	5, 116, 0, 0, 288, 60, 1, 0, 0, 0, 289, 290, 5, 102, 0, 0, 290, 291, 5, 
	114, 0, 0, 291, 292, 5, 111, 0, 0, 292, 293, 5, 109, 0, 0, 293, 62, 1, 
	0, 0, 0, 294, 296, 7, 3, 0, 0, 295, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 
	0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 64, 1, 0, 0, 0, 299, 
	300, 7, 4, 0, 0, 300, 66, 1, 0, 0, 0, 301, 302, 7, 5, 0, 0, 302, 68, 1, 
	0, 0, 0, 9, 0, 72, 80, 85, 90, 177, 228, 272, 297, 0,
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
	OpenFGALexerL_ANGLE_BRACKET = 19
	OpenFGALexerR_ANGLE_BRACKET = 20
	OpenFGALexerCOMMA = 21
	OpenFGALexerCONDITION_PARAM_CONTAINER = 22
	OpenFGALexerCONDITION_PARAM_TYPE = 23
	OpenFGALexerCONDITION_SYMBOL = 24
	OpenFGALexerAND = 25
	OpenFGALexerOR = 26
	OpenFGALexerBUT_NOT = 27
	OpenFGALexerFROM = 28
	OpenFGALexerALPHA_NUMERIC = 29
	OpenFGALexerNEWLINE = 30
	OpenFGALexerWS = 31
)

