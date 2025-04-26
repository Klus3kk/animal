package core

// CONSTANTS AND VARIABLES //

const DIGITS string = "0123456789"
const LETTERS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LETTERS_DIGITS string = LETTERS + DIGITS

var importedFiles = map[string]bool{}
var Debug = false

// TOKENS //

const (
	TT_INT         = "INT"    //
	TT_FLOAT       = "FLOAT"  //
	TT_BOOL        = "BOOL"   //
	TT_STRING      = "STRING" //
	TT_IDEN        = "IDENTIFIER"
	TT_KEY         = "KEYWORD"
	TT_PLUS        = "PLUS"  //
	TT_MINUS       = "MINUS" //
	TT_NEG         = "NEG"   //
	TT_POS         = "POS"   //
	TT_MUL         = "MUL"   //
	TT_DIV         = "DIV"   //
	TT_MOD         = "MOD"   //
	TT_EXP         = "EXP"   //
	TT_CONC        = "CONC"  //
	TT_EQ          = "EQ"
	TT_GT          = "GT"
	TT_LT          = "LT"
	TT_GTE         = "GTE"
	TT_LTE         = "LTE"
	TT_EQEQ        = "EQEQ"
	TT_NEQ         = "NEQ"
	TT_DOT         = "DOT"
	TT_COMMA       = "COMMA"
	TT_LROUNDBR    = "LROUNDBR" //
	TT_RROUNDBR    = "RROUNDBR" //
	TT_RSQRBR      = "RSQRBR"   //
	TT_LSQRBR      = "LSQRBR"   //
	TT_RCURLBR     = "RCURLBR"  //
	TT_LCURLBR     = "LCURLBR"  //
	TT_EOF         = "EOF"      //
	TT_AND         = "AND"
	TT_OR          = "OR"
	TT_TRY_START   = "TT_TRY_START"
	TT_TRY_END     = "TT_TRY_END"
	TT_CATCH_START = "TT_CATCH_START"
	TT_CATCH_END   = "TT_CATCH_END"
	TT_THROW_START = "TT_THROW_START"
	TT_THROW_END   = "TT_THROW_END"
)

var KEYWORDS = []string{
	"int", "float", "bool", "string", // types
	"growl", "sniff", "wag", // if, elif, else
	"roar",           // print
	"pounce", "leap", // while, for
	"howl",                                       // function
	"nest",                                       // data structure
	"listen",                                     // user input
	"sniffback",                                  // return
	"fetch", "drop", "drop_append", "sniff_file", // file i/o
	"fetch_json", "fetch_csv", // Fetching json, csv
	"mimic", "_", // mimic
	"whimper", "hiss", // break, continue
}
