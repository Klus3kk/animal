package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CONSTANTS //

const DIGITS string = "0123456789"
const LETTERS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LETTERS_DIGITS string = LETTERS + DIGITS

// TOKENS //

const (
	TT_INT      = "INT"    //
	TT_FLOAT    = "FLOAT"  //
	TT_BOOL     = "BOOL"   //
	TT_STRING   = "STRING" //
	TT_IDEN     = "IDENTIFIER"
	TT_KEY      = "KEYWORD"
	TT_PLUS     = "PLUS"  //
	TT_MINUS    = "MINUS" //
	TT_NEG      = "NEG"   //
	TT_POS      = "POS"   //
	TT_MUL      = "MUL"   //
	TT_DIV      = "DIV"   //
	TT_MOD      = "MOD"   //
	TT_EXP      = "EXP"   //
	TT_CONC     = "CONC"  //
	TT_EQ       = "EQ"
	TT_GT       = "GT"
	TT_LT       = "LT"
	TT_GTE      = "GTE"
	TT_LTE      = "LTE"
	TT_EQEQ     = "EQEQ"
	TT_NEQ      = "NEQ"
	TT_DOT      = "DOT"
	TT_COMMA    = "COMMA"
	TT_LROUNDBR = "LROUNDBR" //
	TT_RROUNDBR = "RROUNDBR" //
	TT_RSQRBR   = "RSQRBR"   //
	TT_LSQRBR   = "LSQRBR"   //
	TT_RCURLBR  = "RCURLBR"  //
	TT_LCURLBR  = "LCURLBR"  //
	TT_EOF      = "EOF"      //
	TT_AND      = "AND"
	TT_OR       = "OR"
)

var KEYWORDS = []string{
	"int", "float", "bool", "string", // types
	"growl", "sniff", "wag", // if, elif, else
	"roar",           // print
	"pounce", "leap", // while, for
	"howl",      // function
	"nest",      // data structure
	"listen",    // user input
	"sniffback", // return
}

// Token represents a token with its type and value
type Token struct {
	Type      string
	Value     string
	Pos_Start *Position
	Pos_End   *Position
}

// String returns a string representation of the Token
func (t Token) String() string {
	if t.Value != "" {
		return fmt.Sprintf("(%s): [%s]", t.Type, t.Value)
	}
	return fmt.Sprintf("(%s)", t.Type)
}

func (t Token) matches(type_ string, value string) bool {
	return string(t.Type) == string(type_) && t.Value == value
}

// LEXER //

type Lexer struct {
	Text        string
	Pos         *Position
	CurrentChar byte
	Fn          string
}

func NewLexer(fn, text string) *Lexer {
	lexer := &Lexer{Text: text, Pos: NewPosition(0, 0, 0, fn, text), Fn: fn}
	if len(text) > 0 {
		lexer.CurrentChar = text[0]
	} else {
		lexer.CurrentChar = 0
	}
	return lexer
}

func (l *Lexer) MakeTokens() ([]Token, error) {
	return l.make_tokens()
}

func (l *Lexer) advance() {
	l.Pos.advance(l.CurrentChar)
	if l.Pos.Idx < len(l.Text) {
		l.CurrentChar = l.Text[l.Pos.Idx]
	} else {
		l.CurrentChar = 0 // None
	}
}

func (l *Lexer) peek(length int) string {
	endPos := l.Pos.Idx + length
	if endPos >= len(l.Text) {
		endPos = len(l.Text)
	}
	return l.Text[l.Pos.Idx:endPos]
}

func (l *Lexer) make_tokens() ([]Token, error) {
	tokens := []Token{}
	var err error

	for l.CurrentChar != 0 { // while current character isn't None
		if l.CurrentChar == ' ' || l.CurrentChar == '\t' || l.CurrentChar == '\r' || l.CurrentChar == '\n' {
			l.advance()
		} else if strings.IndexByte(DIGITS, l.CurrentChar) != -1 {
			tokens = append(tokens, l.make_number()) // Tokenize number
		} else if l.CurrentChar == '"' {
			tokens = append(tokens, l.make_string()) // Tokenize string
		} else if l.peek(4) == "true" || l.peek(5) == "false" {
			tokens = append(tokens, l.make_boolean()) // Tokenize boolean
		} else if l.peek(4) == "meow" {
			posStart := l.Pos.copy() // Copy the current position as the token's start
			l.advanceBy(4)
			posEnd := l.Pos.copy() // After advancing, copy for the token's end
			tokens = append(tokens, Token{Type: TT_PLUS, Value: "PLUS", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(4) == "woof" {
			posStart := l.Pos.copy()
			l.advanceBy(4)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_MINUS, Value: "MINUS", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(3) == "moo" {
			posStart := l.Pos.copy()
			l.advanceBy(3)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_MUL, Value: "MUL", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(5) == "drone" {
			posStart := l.Pos.copy()
			l.advanceBy(5)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_DIV, Value: "DIV", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(6) == "squeak" {
			posStart := l.Pos.copy()
			l.advanceBy(6)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_MOD, Value: "MOD", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(4) == "soar" {
			posStart := l.Pos.copy()
			l.advanceBy(4)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_EXP, Value: "EXP", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(4) == "purr" {
			posStart := l.Pos.copy()
			l.advanceBy(4)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_CONC, Value: "CONC", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(6) == "pounce" {
			posStart := l.Pos.copy()
			l.advanceBy(6)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_KEY, Value: "pounce", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(4) == "leap" {
			posStart := l.Pos.copy()
			l.advanceBy(4)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_KEY, Value: "leap", Pos_Start: posStart, Pos_End: posEnd})
		} else if strings.IndexByte(LETTERS, l.CurrentChar) != -1 {
			tokens = append(tokens, l.make_identifier()) // Tokenize letter
		} else if l.peek(2) == "->" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_EQ, Value: "EQ", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '(' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_LROUNDBR, Value: string(l.CurrentChar), Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '.' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_DOT, Value: ".", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == ')' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_RROUNDBR, Value: string(l.CurrentChar), Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '[' {
			posStart := l.Pos.copy()
			l.advance()
			tokens = append(tokens, Token{Type: TT_LSQRBR, Value: "[", Pos_Start: posStart, Pos_End: l.Pos.copy()})
		} else if l.CurrentChar == ']' {
			posStart := l.Pos.copy()
			l.advance()
			tokens = append(tokens, Token{Type: TT_RSQRBR, Value: "]", Pos_Start: posStart, Pos_End: l.Pos.copy()})
		} else if l.CurrentChar == '{' {
			posStart := l.Pos.copy()
			l.advance()
			tokens = append(tokens, Token{Type: TT_LCURLBR, Value: "{", Pos_Start: posStart, Pos_End: l.Pos.copy()})
		} else if l.CurrentChar == '}' {
			posStart := l.Pos.copy()
			l.advance()
			tokens = append(tokens, Token{Type: TT_RCURLBR, Value: "}", Pos_Start: posStart, Pos_End: l.Pos.copy()})
		} else if l.peek(2) == ">=" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_GTE, Value: ">=", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "<=" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_LTE, Value: "<=", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "==" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_EQEQ, Value: "==", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "!=" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_NEQ, Value: "!=", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '>' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_GT, Value: ">", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '<' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_LT, Value: "<", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '-' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_NEG, Value: string(l.CurrentChar), Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '+' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_POS, Value: string(l.CurrentChar), Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == ',' {
			posStart := l.Pos.copy()
			l.advance()
			tokens = append(tokens, Token{Type: TT_COMMA, Value: ",", Pos_Start: posStart, Pos_End: l.Pos.copy()})
		} else if l.peek(4) == "roar" {
			posStart := l.Pos.copy()
			l.advanceBy(4)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_KEY, Value: "roar", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(4) == "howl" {
			posStart := l.Pos.copy()
			l.advanceBy(4)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_KEY, Value: "howl", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(3) == "and" {
			posStart := l.Pos.copy()
			l.advanceBy(3)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_AND, Value: "and", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "or" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_OR, Value: "or", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "::" {
			l.advanceBy(2)
			for l.CurrentChar != 0 && l.CurrentChar != '\n' {
				l.advance() // Skip until end of line
			}
		} else {
			posStart := l.Pos.copy()
			char := string(l.CurrentChar)
			fmt.Printf("Invalid char encountered: [%q] Unicode: %U\n", char, l.CurrentChar)
			l.advance()
			err = fmt.Errorf("%s: Unexpected Character: %q", posStart.asString(), char)
			break
		}

	}
	tokens = append(tokens, Token{Type: TT_EOF, Value: "EOF", Pos_Start: l.Pos.copy(), Pos_End: l.Pos.copy()})
	for _, tok := range tokens { // debug
		fmt.Println("TOKEN:", tok)
	}

	return tokens, err
}

func (l *Lexer) advanceBy(count int) {
	for i := 0; i < count; i++ {
		l.advance()
	}
}

func (l *Lexer) make_identifier() Token {
	idStr := ""
	Pos_Start := l.Pos.copy()

	for l.CurrentChar != 0 && strings.ContainsAny(string(l.CurrentChar), LETTERS_DIGITS+"_") {
		idStr += string(l.CurrentChar)
		l.advance()
	}
	tok_Type := TT_IDEN

	// Check if identifier is a keyword only if not immediately after '.'
	if !(l.Pos.Idx > 0 && l.Text[l.Pos.Idx-1] == '.') {
		for _, keyword := range KEYWORDS {
			if idStr == keyword {
				tok_Type = TT_KEY
				break
			}
		}
	}

	return Token{Type: tok_Type, Value: idStr, Pos_Start: Pos_Start, Pos_End: l.Pos}
}

func (l *Lexer) make_number() Token {
	numStr := ""
	dotCount := 0
	Pos_Start := l.Pos.copy()

	if l.CurrentChar == '-' || l.CurrentChar == '+' {
		numStr += string(l.CurrentChar)
		l.advance()
	}

	for l.CurrentChar != 0 && strings.ContainsAny(string(l.CurrentChar), DIGITS+".") {
		if l.CurrentChar == '.' {
			if dotCount == 1 {
				break
			}
			dotCount++
			numStr += "."
		} else {
			numStr += string(l.CurrentChar)
		}
		l.advance()
	}

	if dotCount == 0 {
		intValue, _ := strconv.Atoi(numStr) // Convert numStr to int
		return Token{Type: TT_INT, Value: strconv.Itoa(intValue), Pos_Start: Pos_Start, Pos_End: l.Pos}
	} else {
		floatValue, _ := strconv.ParseFloat(numStr, 64) // Convert numStr to float64
		return Token{Type: TT_FLOAT, Value: strconv.FormatFloat(floatValue, 'f', -1, 64), Pos_Start: Pos_Start, Pos_End: l.Pos}
	}
}

// Make a string token
func (l *Lexer) make_string() Token {
	posStart := l.Pos.copy()
	l.advance() // Skip opening quote
	strVal := ""

	for l.CurrentChar != 0 && l.CurrentChar != '"' {
		strVal += string(l.CurrentChar)
		l.advance()
	}

	l.advance() // Skip closing quote
	return Token{Type: TT_STRING, Value: strVal, Pos_Start: posStart, Pos_End: l.Pos.copy()}
}

// Make a boolean token
func (l *Lexer) make_boolean() Token {
	posStart := l.Pos.copy()
	if l.peek(4) == "true" {
		l.advanceBy(4)
		return Token{Type: TT_BOOL, Value: "true", Pos_Start: posStart, Pos_End: l.Pos.copy()}
	} else if l.peek(5) == "false" {
		l.advanceBy(5)
		return Token{Type: TT_BOOL, Value: "false", Pos_Start: posStart, Pos_End: l.Pos.copy()}
	}
	return Token{}
}

// SNIFFBACK (return)
type SniffbackNode struct {
	Value interface{}
}

func (n SniffbackNode) String() string {
	return fmt.Sprintf("(SNIFFBACK %v)", n.Value)
}

// GROWLNODE (if, else if, else)
type GrowlNode struct {
	Cases    []ConditionBlock // List of conditions and bodies
	ElseCase interface{}      // Optional else case
}

func (n GrowlNode) String() string {
	return fmt.Sprintf("(GROWL %v ELSE %v)", n.Cases, n.ElseCase)
}

// WHILE/FOR LOOPS
// Node for pounce (while) loops
type PounceNode struct {
	Condition interface{}
	Body      []interface{}
}

// Node for leap (for) loops
type LeapNode struct {
	VarName   Token
	StartExpr interface{}
	EndExpr   interface{}
	Body      interface{}
}

// For one condition-body pair
type ConditionBlock struct {
	Condition interface{}
	Body      interface{}
}

// ROARNODE (print)
type RoarNode struct {
	Value interface{}
}

func (n RoarNode) String() string {
	return fmt.Sprintf("(ROAR %v)", n.Value)
}

// RTRESULT
// RTResult is a type that represents the result of an operation
type RTResult struct {
	Value interface{}
	Error error
}

// NewRTResult creates a new RTResult instance
func NewRTResult() *RTResult {
	return &RTResult{}
}

// Success sets the result value and returns it
func (r *RTResult) success(value interface{}) *RTResult {
	r.Value = value
	return r
}

// Failure sets the error and returns it
func (r *RTResult) failure(err error) *RTResult {
	r.Error = err
	return r
}

// Register handles the result of another RTResult, propagating errors if necessary
func (r *RTResult) register(res *RTResult) interface{} {
	if res.Error != nil {
		r.Error = res.Error
	}
	return res.Value
}

// ERRORS //

type Error struct {
	ErrorName string
	Details   string
	PosStart  *Position
	PosEnd    *Position
}

func (e Error) asString() string {
	result := fmt.Sprintf("%s: %s", e.ErrorName, e.Details)
	result += fmt.Sprintf(" File %s, line %d", e.PosStart.Fn, e.PosStart.Ln+1)
	return result
}

// IllegalCharError struct inheriting from Error
type IllegalCharError struct {
	Error
}

// test
func NewIllegalCharError(posStart, posEnd *Position, details string) *IllegalCharError {
	return &IllegalCharError{
		Error: Error{
			PosStart:  posStart,
			PosEnd:    posEnd,
			ErrorName: "Illegal Character",
			Details:   details,
		},
	}
}

// InvalidSyntaxError struct inheriting from Error
type InvalidSyntaxError struct {
	Error
}

func NewInvalidSyntaxError(posStart, posEnd *Position, details string) *InvalidSyntaxError {
	return &InvalidSyntaxError{
		Error: Error{
			PosStart:  posStart,
			PosEnd:    posEnd,
			ErrorName: "Invalid Syntax",
			Details:   details,
		},
	}
}

// RTERROR //
type RTError struct {
	Error
	Context *Context
}

func (e *RTError) AsString() string {
	result := e.Context.GenerateTraceback()
	result += fmt.Sprintf("%s: %s\n\n", e.ErrorName, e.Details)
	return result
}

// POSITION //

type Position struct {
	Idx  int
	Ln   int
	Col  int
	Fn   string
	Ftxt string
}

func NewPosition(idx, ln, col int, fn, ftxt string) *Position {
	return &Position{Idx: idx, Ln: ln, Col: col, Fn: fn, Ftxt: ftxt}
}

func (p *Position) asString() string {
	return fmt.Sprintf("File %s, line %d", p.Fn, p.Ln+1)
}

func (p *Position) advance(currentChar byte) {
	p.Idx++
	p.Col++

	if currentChar == '\n' {
		p.Ln++
		p.Col = 0
	}
}

func (p *Position) copy() *Position {
	return NewPosition(p.Idx, p.Ln, p.Col, p.Fn, p.Ftxt)
}

// NODES //
// __init__ (self, Tok)
type StringNode struct {
	Tok Token
}

func (n StringNode) String() string {
	return fmt.Sprintf("(%s: %s)", n.Tok.Type, n.Tok.Value)
}

type BoolNode struct {
	Tok Token
}

func (n BoolNode) String() string {
	return fmt.Sprintf("(%s: %s)", n.Tok.Type, n.Tok.Value)
}

type NumberNode struct {
	Tok Token
}

func (n NumberNode) String() string {
	return fmt.Sprintf("(%s: %s)", n.Tok.Type, n.Tok.Value)
}

type BinOpNode struct {
	Left_Node  interface{}
	Op_Tok     Token
	Right_Node interface{}
}

func (b BinOpNode) String() string { // __repr__
	return fmt.Sprintf("(%s %s %s)", b.Left_Node, b.Op_Tok.Value, b.Right_Node)
}

type UnaryOpNode struct {
	Op_Tok Token
	Node   interface{}
}

func (u UnaryOpNode) String() string { // __repr__
	return fmt.Sprintf("(%s %s)", u.Op_Tok.Type, u.Node)
}

type VarAccessNode struct {
	Var_Name_Tok Token
}

type VarAssignNode struct {
	Var_Name_Tok Token
	Value_Node   interface{}
}

func (n VarAssignNode) String() string {
	return fmt.Sprintf("(VarAssignNode %s -> %v)", n.Var_Name_Tok.Value, n.Value_Node)
}

// PARSE RESULT //

type ParseResult struct {
	Error string
	Node  interface{}
}

// Registers a ParseResult, storing its error and node
func (pr *ParseResult) register(res *ParseResult) interface{} {
	if res.Error != "" {
		pr.Error = res.Error
	}
	return res.Node
}

// Registers a successful parsing result
func (pr *ParseResult) success(node interface{}) *ParseResult {
	pr.Node = node
	return pr
}

// Registers a failure with an error message
func (pr *ParseResult) failure(error string) *ParseResult {
	pr.Error = error
	return pr
}

func (p *Parser) Parse() *ParseResult {
	return p.parse()
}

func (p *Parser) peek() Token {
	if p.Tok_Idx+1 < len(p.Tokens) {
		return p.Tokens[p.Tok_Idx+1]
	}
	return Token{Type: TT_EOF}
}

func (p *Parser) parseDotCalls(base interface{}) *ParseResult {
	res := &ParseResult{}
	node := base

	for p.Current_Tok.Type == TT_DOT {
		p.advance()

		if p.Current_Tok.Type != TT_IDEN && p.Current_Tok.Type != TT_KEY {
			return res.failure("Expected property or method name after '.'")
		}

		method := p.Current_Tok.Value
		p.Current_Tok.Type = TT_IDEN
		p.advance()

		// 👇 Check if it's a method call (with parentheses)
		if p.Current_Tok.Type == TT_LROUNDBR {
			p.advance()

			args := []interface{}{}
			if p.Current_Tok.Type != TT_RROUNDBR {
				for {
					arg := res.register(p.expr())
					if res.Error != "" {
						return res
					}
					args = append(args, arg)
					if p.Current_Tok.Type == TT_COMMA {
						p.advance()
					} else {
						break
					}
				}
			}

			if p.Current_Tok.Type != TT_RROUNDBR {
				return res.failure("Expected ')' after method arguments")
			}
			p.advance()

			node = DotCallNode{
				Target: node,
				Method: method,
				Args:   args, // method call
			}
		} else {
			node = DotCallNode{
				Target: node,
				Method: method,
				Args:   nil, // property get or set
			}
		}
	}

	return res.success(node)
}

func (p *Parser) nest_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "nest") {
		return res.failure("Expected 'nest'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_IDEN {
		return res.failure("Expected name after 'nest'")
	}
	nestName := p.Current_Tok.Value
	p.advance()

	if p.Current_Tok.Type != TT_LCURLBR {
		return res.failure("Expected '{' after nest name")
	}
	p.advance()

	fields := []string{}
	methods := map[string]*FunctionDefNode{}

	for p.Current_Tok.Type != TT_RCURLBR && p.Current_Tok.Type != TT_EOF {
		if p.Current_Tok.matches(TT_KEY, "howl") {
			fn := res.register(p.howl_expr())
			if res.Error != "" {
				return res
			}
			funcNode := fn.(FunctionDefNode)
			methods[funcNode.Name] = &funcNode
		} else if p.Current_Tok.Type == TT_IDEN {
			fields = append(fields, p.Current_Tok.Value)
			p.advance()
		} else {
			return res.failure("Expected property or method in nest body")
		}
	}

	if p.Current_Tok.Type != TT_RCURLBR {
		return res.failure("Expected '}' to close nest block")
	}
	p.advance()

	return res.success(NestDefNode{
		Name:     nestName,
		Fields:   fields,
		Methods:  methods,
		PosStart: nil,
		PosEnd:   nil,
	})
}

func (p *Parser) howl_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "howl") {
		return res.failure("Expected 'howl'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_IDEN {
		return res.failure("Expected function name after 'howl'")
	}
	funcName := p.Current_Tok.Value
	p.advance()

	if p.Current_Tok.Type != TT_LROUNDBR {
		return res.failure("Expected '(' after function name")
	}
	p.advance()

	argNames := []string{}
	for p.Current_Tok.Type == TT_IDEN {
		argNames = append(argNames, p.Current_Tok.Value)
		p.advance()
		if p.Current_Tok.Type == TT_COMMA {
			p.advance()
		} else {
			break
		}
	}

	if p.Current_Tok.Type != TT_RROUNDBR {
		return res.failure("Expected ')' after argument list")
	}
	p.advance()

	if p.Current_Tok.Type != TT_LCURLBR {
		return res.failure("Expected '{' before function body")
	}
	p.advance()

	statements := []interface{}{}
	for p.Current_Tok.Type != TT_RCURLBR && p.Current_Tok.Type != TT_EOF {
		stmt := res.register(p.expr())
		if res.Error != "" {
			return res
		}
		statements = append(statements, stmt)
	}
	body := StatementsNode{Statements: statements}

	if p.Current_Tok.Type != TT_RCURLBR {
		return res.failure("Expected '}' after function body")
	}
	p.advance()

	return res.success(FunctionDefNode{
		Name:     funcName,
		ArgNames: argNames,
		Body:     body,
	})
}

func (p *Parser) roar_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "roar") {
		return res.failure("Expected 'roar'")
	}
	p.advance()

	values := []interface{}{}

	// If it's an empty roar like `roar`, just print newline
	if p.Current_Tok.Type == TT_EOF || p.Current_Tok.Type == TT_KEY || p.Current_Tok.Type == TT_RCURLBR {
		return res.success(RoarNode{Value: nil})
	}

	// Parse one or more comma-separated expressions
	expr := res.register(p.expr())
	if res.Error != "" {
		return res
	}
	values = append(values, expr)

	for p.Current_Tok.Type == TT_COMMA {
		p.advance()

		expr := res.register(p.expr())
		if res.Error != "" {
			return res
		}
		values = append(values, expr)
	}

	return res.success(RoarNode{Value: values})
}

func (p *Parser) comp_expr() *ParseResult {
	// Comparison ops: ==, !=, >, <, >=, <=
	return p.bin_op(p.term, []string{TT_EQEQ, TT_NEQ, TT_GT, TT_LT, TT_GTE, TT_LTE})
}

func (p *Parser) growl_expr() *ParseResult {
	res := &ParseResult{}
	cases := []ConditionBlock{}
	var elseCase interface{}

	// Handle "growl" (if)
	if !p.Current_Tok.matches(TT_KEY, "growl") {
		return res.failure("Expected 'growl'")
	}
	p.advance()
	fmt.Println("Advanced token:", p.Current_Tok)

	condition := res.register(p.bin_op(p.comp_expr, []string{TT_AND, TT_OR}))

	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_LCURLBR { // Expecting '{'
		return res.failure("Expected '{' after growl condition")
	}
	p.advance()
	fmt.Println("Advanced token:", p.Current_Tok)

	stmts := []interface{}{}
	for p.Current_Tok.Type != TT_RCURLBR && p.Current_Tok.Type != TT_EOF {
		stmt := res.register(p.expr())
		if res.Error != "" {
			return res
		}
		stmts = append(stmts, stmt)
	}
	body := StatementsNode{Statements: stmts}

	if p.Current_Tok.Type != TT_RCURLBR { // Expecting '}'
		return res.failure("Expected '}' after growl body")
	}
	p.advance()
	fmt.Println("Advanced token:", p.Current_Tok)

	cases = append(cases, ConditionBlock{Condition: condition, Body: body})

	// Handle "sniff" (else-if)
	for p.Current_Tok.matches(TT_KEY, "sniff") {
		p.advance()
		fmt.Println("Advanced token:", p.Current_Tok)

		condition := res.register(p.bin_op(p.comp_expr, []string{TT_AND, TT_OR}))

		if res.Error != "" {
			return res
		}

		if p.Current_Tok.Type != TT_LCURLBR {
			return res.failure("Expected '{' after sniff condition")
		}
		p.advance()
		fmt.Println("Advanced token:", p.Current_Tok)

		stmts := []interface{}{}
		for p.Current_Tok.Type != TT_RCURLBR && p.Current_Tok.Type != TT_EOF {
			stmt := res.register(p.expr())
			if res.Error != "" {
				return res
			}
			stmts = append(stmts, stmt)
		}
		body := StatementsNode{Statements: stmts}

		if p.Current_Tok.Type != TT_RCURLBR {
			return res.failure("Expected '}' after sniff body")
		}
		p.advance()
		fmt.Println("Advanced token:", p.Current_Tok)

		cases = append(cases, ConditionBlock{Condition: condition, Body: body})
	}

	// Handle "wag" (else)
	if p.Current_Tok.matches(TT_KEY, "wag") {
		p.advance()
		fmt.Println("Advanced token:", p.Current_Tok)

		if p.Current_Tok.Type != TT_LCURLBR {
			return res.failure("Expected '{' after wag")
		}
		p.advance()
		fmt.Println("Advanced token:", p.Current_Tok)

		stmts := []interface{}{}
		for p.Current_Tok.Type != TT_RCURLBR && p.Current_Tok.Type != TT_EOF {
			stmt := res.register(p.expr())
			if res.Error != "" {
				return res
			}
			stmts = append(stmts, stmt)
		}

		elseCase = StatementsNode{Statements: stmts}

		if p.Current_Tok.Type != TT_RCURLBR {
			return res.failure("Expected '}' after wag body")
		}
		p.advance()
		fmt.Println("Advanced token:", p.Current_Tok)

	}

	return res.success(GrowlNode{Cases: cases, ElseCase: elseCase})
}

// PARSER //

type Parser struct {
	Tokens      []Token
	Tok_Idx     int
	Current_Tok Token
}

func NewParser(tokens []Token) *Parser {
	parser := &Parser{Tokens: tokens, Tok_Idx: -1}
	parser.advance()
	return parser
}

func (p *Parser) advance() Token {
	p.Tok_Idx++

	if p.Tok_Idx < len(p.Tokens) {
		p.Current_Tok = p.Tokens[p.Tok_Idx]
	} else {
		p.Current_Tok = Token{Type: TT_EOF}
	}
	return p.Current_Tok
}

func (p *Parser) parse() *ParseResult {
	res := &ParseResult{}
	statements := []interface{}{}

	for p.Current_Tok.Type != TT_EOF {
		fmt.Println("Parsing statement, current token:", p.Current_Tok)

		stmt := res.register(p.expr())
		if res.Error != "" {
			return res
		}

		fmt.Println("Parsed statement:", stmt)
		statements = append(statements, stmt)
	}

	return res.success(StatementsNode{Statements: statements})
}

// Statement Node for multiple top-level expressions
type StatementsNode struct {
	Statements []interface{}
}

func (s StatementsNode) String() string {
	return fmt.Sprintf("(%s)", s.Statements)
}

// ListNode
type ListNode struct {
	Elements []interface{}
}

// HIERARCHY //

// Exponentiation, highest precedence
func (p *Parser) power() *ParseResult {
	return p.bin_op(p.atom, []string{TT_EXP})
}

// For recurssion
func (p *Parser) call_expr() *ParseResult {
	res := &ParseResult{}
	node := res.register(p.atom())
	if res.Error != "" {
		return res
	}
	// debug
	fmt.Println("call_expr: building call for", node)

	for p.Current_Tok.Type == TT_LROUNDBR {
		p.advance()
		args := []interface{}{}
		if p.Current_Tok.Type != TT_RROUNDBR {
			for {
				arg := res.register(p.expr()) // <-- parse full expressions
				if res.Error != "" {
					return res
				}
				args = append(args, arg)
				if p.Current_Tok.Type == TT_COMMA {
					p.advance()
				} else {
					break
				}
			}
		}
		if p.Current_Tok.Type != TT_RROUNDBR {
			return res.failure("Expected ')' after function arguments")
		}
		p.advance()

		// Attach call
		if varAccess, ok := node.(VarAccessNode); ok {
			node = FunctionCallNode{
				FuncName: varAccess.Var_Name_Tok.Value,
				Args:     args,
			}
		} else {
			return res.failure("Cannot call non-function")
		}
	}

	return res.success(node)
}

// Handles parentheses and atoms (integers, booleans, strings)
func (p *Parser) atom() *ParseResult {
	res := &ParseResult{}
	tok := p.Current_Tok
	if p.Current_Tok.Type == TT_EOF {
		return res.failure("Unexpected end of file")
	}

	if tok.Type == TT_INT || tok.Type == TT_FLOAT {
		p.advance()
		return p.parseDotCalls(NumberNode{Tok: tok})
	} else if tok.Type == TT_BOOL {
		p.advance()
		return p.parseDotCalls(BoolNode{Tok: tok})
	} else if tok.Type == TT_STRING {
		p.advance()
		return p.parseDotCalls(StringNode{Tok: tok})
	} else if tok.Type == TT_IDEN {
		p.advance()

		// Function call
		if p.Current_Tok.Type == TT_LROUNDBR {
			p.advance()
			args := []interface{}{}
			if p.Current_Tok.Type != TT_RROUNDBR {
				for {
					arg := res.register(p.expr())
					if res.Error != "" {
						return res
					}
					args = append(args, arg)
					if p.Current_Tok.Type == TT_COMMA {
						p.advance()
					} else {
						break
					}
				}
			}
			if p.Current_Tok.Type != TT_RROUNDBR {
				return res.failure("Expected ')' after function arguments")
			}
			p.advance()

			node := FunctionCallNode{
				FuncName: tok.Value,
				Args:     args,
			}
			return p.parseDotCalls(node)
		}

		// Variable access
		node := VarAccessNode{Var_Name_Tok: tok}
		return p.parseDotCalls(node)
	} else if tok.Type == TT_LROUNDBR {
		p.advance()
		expr := res.register(p.expr())
		if res.Error != "" {
			return res
		}
		if p.Current_Tok.Type != TT_RROUNDBR {
			return res.failure("Expected ')' after expression")
		}
		p.advance()
		return p.parseDotCalls(expr)
	} else if tok.Type == TT_LSQRBR {
		p.advance()
		elements := []interface{}{}
		if p.Current_Tok.Type != TT_RSQRBR {
			for {
				elem := res.register(p.expr())
				if res.Error != "" {
					return res
				}
				elements = append(elements, elem)
				if p.Current_Tok.Type == TT_COMMA {
					p.advance()
				} else {
					break
				}
			}
		}
		if p.Current_Tok.Type != TT_RSQRBR {
			return res.failure("Expected ']' after list")
		}
		p.advance()
		return p.parseDotCalls(ListNode{Elements: elements})
	} else if tok.Type == TT_LCURLBR {
		p.advance()
		expr := res.register(p.expr())
		if res.Error != "" {
			return res
		}
		if p.Current_Tok.Type != TT_RCURLBR {
			return res.failure("Expected '}'")
		}
		p.advance()
		return p.parseDotCalls(expr)
	}

	return res.failure("Expected int, float, boolean, string, identifier, or bracketed expression")
}

// Unary operations: + and -
func (p *Parser) factor() *ParseResult {
	res := &ParseResult{}
	tok := p.Current_Tok

	if tok.Type == TT_POS || tok.Type == TT_NEG {
		p.advance()
		factor := res.register(p.factor())
		if res.Error != "" {
			return res
		}
		return res.success(UnaryOpNode{Op_Tok: tok, Node: factor})
	}

	// Proceed to the next higher precedence operation
	return p.bin_op(p.call_expr, []string{TT_EXP})
}

// Multiplication, division, modulo, and concatenation
func (p *Parser) term() *ParseResult {
	return p.bin_op(p.factor, []string{TT_MUL, TT_DIV, TT_MOD, TT_CONC, TT_PLUS, TT_MINUS})
}

// Modified expr function to handle variable assignments with types and operations
func (p *Parser) expr() *ParseResult {
	res := &ParseResult{}

	// Print statement
	if p.Current_Tok.matches(TT_KEY, "roar") {
		return p.roar_expr()
	}

	// Conditional statements
	if p.Current_Tok.matches(TT_KEY, "growl") {
		return p.growl_expr()
	}

	// Loops
	if p.Current_Tok.matches(TT_KEY, "pounce") {
		return p.pounce_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "leap") {
		return p.leap_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "howl") {
		return p.howl_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "nest") {
		return p.nest_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "listen") {
		p.advance()
		return res.success(ListenNode{})
	}

	// Handle variable access and assignment
	// If next token is EQ (->), parse as assignment
	if p.Current_Tok.Type == TT_IDEN && p.peek().Type == TT_EQ {
		target := res.register(p.atom())
		if res.Error != "" {
			return res
		}

		p.advance() // consume ->
		value_expr := res.register(p.expr())
		if res.Error != "" {
			return res
		}

		if dotNode, ok := target.(DotCallNode); ok {
			dotNode.Args = []interface{}{value_expr}
			return res.success(dotNode)
		}
		if varNode, ok := target.(VarAccessNode); ok {
			return res.success(VarAssignNode{Var_Name_Tok: varNode.Var_Name_Tok, Value_Node: value_expr})
		}

		return res.failure("Invalid assignment target")
	}

	// Else parse the full binary expression
	node := res.register(p.bin_op(p.comp_expr, []string{TT_AND, TT_OR}))
	if res.Error != "" {
		return res
	}

	// sniffback handling
	if p.Current_Tok.matches(TT_KEY, "sniffback") {
		p.advance()
		return res.success(SniffbackNode{Value: node})
	}

	// ✅ dot-assignment support
	if p.Current_Tok.Type == TT_EQ {
		p.advance()
		value_expr := res.register(p.expr())
		if res.Error != "" {
			return res
		}

		// Must be a dot-access target
		if dotNode, ok := node.(DotCallNode); ok {
			dotNode.Args = []interface{}{value_expr}
			return res.success(dotNode)
		}
		return res.failure("Assignment target must be a variable or object property")
	}

	return res.success(node)

}

func (p *Parser) pounce_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "pounce") {
		return res.failure("Expected 'pounce'")
	}
	p.advance()

	// Debug
	// fmt.Println("After 'pounce', token:", p.Current_Tok)

	// The condition is actually a comparison (binary operation)
	// First get the left side (variable or value)
	left := res.register(p.atom())
	if res.Error != "" {
		return res
	}

	// Debug
	// fmt.Println("Left side parsed:", left)
	// fmt.Println("Current token after left side:", p.Current_Tok)

	// Now expect a comparison operator
	if !contains([]string{TT_GT, TT_LT, TT_GTE, TT_LTE, TT_EQEQ, TT_NEQ}, p.Current_Tok.Type) {
		return res.failure(fmt.Sprintf("Expected comparison operator, got %s", p.Current_Tok.Type))
	}

	// Save the operator
	op_tok := p.Current_Tok
	p.advance()

	// Debug
	// fmt.Println("Operator parsed:", op_tok)
	// fmt.Println("Current token after operator:", p.Current_Tok)

	// Parse the right side of the comparison
	right := res.register(p.atom())
	if res.Error != "" {
		return res
	}

	// Debug
	// fmt.Println("Right side parsed:", right)
	// fmt.Println("Current token after right side:", p.Current_Tok)

	// Build the condition as a binary operation
	condition := BinOpNode{
		Left_Node:  left,
		Op_Tok:     op_tok,
		Right_Node: right,
	}

	// Now expect the opening brace
	if p.Current_Tok.Type != TT_LCURLBR {
		return res.failure(fmt.Sprintf("Expected '{' after pounce condition, got %s", p.Current_Tok.Type))
	}
	p.advance()

	// Debug
	// fmt.Println("Opening brace found, parsing body")

	// Parse body
	body := []interface{}{}
	for p.Current_Tok.Type != TT_RCURLBR && p.Current_Tok.Type != TT_EOF {
		stmt := res.register(p.expr())
		if res.Error != "" {
			return res
		}
		body = append(body, stmt)

		// Debug
		// fmt.Println("Added statement to body:", stmt)
	}

	// Check for closing brace
	if p.Current_Tok.Type != TT_RCURLBR {
		return res.failure("Expected '}' at end of pounce body")
	}
	p.advance()

	// Debug
	// fmt.Println("Successfully parsed pounce loop")

	return res.success(&PounceNode{
		Condition: condition,
		Body:      body,
	})
}

func (p *Parser) leap_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "leap") {
		return res.failure("Expected 'leap'")
	}
	p.advance()

	// Expect a variable name (identifier)
	if p.Current_Tok.Type != TT_IDEN {
		return res.failure("Expected loop variable name after 'leap'")
	}
	varName := p.Current_Tok
	p.advance()

	// Expect 'from'
	if p.Current_Tok.Value != "from" {
		return res.failure("Expected 'from' after loop variable")
	}
	p.advance()

	// Parse start expression
	startExpr := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	// Expect 'to'
	if p.Current_Tok.Value != "to" {
		return res.failure("Expected 'to' after start value")
	}
	p.advance()

	// Parse end expression
	endExpr := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	// Expect '{'
	if p.Current_Tok.Type != TT_LCURLBR {
		return res.failure("Expected '{' before loop body")
	}
	p.advance()

	// Parse loop body
	body := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	// Expect '}'
	if p.Current_Tok.Type != TT_RCURLBR {
		return res.failure("Expected '}' after loop body")
	}
	p.advance()

	return res.success(LeapNode{
		VarName:   varName,
		StartExpr: startExpr,
		EndExpr:   endExpr,
		Body:      body,
	})
}

func (p *Parser) bin_op(funcToCall func() *ParseResult, ops []string) *ParseResult {
	res := &ParseResult{}

	left := res.register(funcToCall())
	if res.Error != "" {
		return res
	}

	for p.Current_Tok.Type != TT_EOF && contains(ops, p.Current_Tok.Type) {
		// Don't consume further if next token is '{' — that's for blocks like growl/wag
		if p.peek().Type == TT_LCURLBR || p.Current_Tok.Type == TT_LCURLBR {
			break
		}

		fmt.Println("Found operator:", p.Current_Tok)
		op_tok := p.Current_Tok
		p.advance()

		fmt.Println("Current token after operator:", p.Current_Tok)
		right := res.register(funcToCall())
		if res.Error != "" {
			return res
		}

		fmt.Println("Right side parsed:", right)
		left = BinOpNode{Left_Node: left, Op_Tok: op_tok, Right_Node: right}
	}

	return res.success(left)
}

// Utility function to check if a TokenType is in the list
func contains(ops []string, op string) bool {
	for _, val := range ops {
		if val == op {
			return true
		}
	}
	return false
}

// DotCall

type DotCallNode struct {
	Target interface{}
	Method string
	Args   []interface{}
}

// CONTEXT //
type Context struct {
	DisplayName    string
	Parent         *Context
	ParentEntryPos *Position
	Symbol_Table   *SymbolTable
}

func (c *Context) GenerateTraceback() string {
	result := ""
	pos := c.ParentEntryPos
	ctx := c

	for ctx != nil {
		result = fmt.Sprintf("File %s, line %d, in %s\n", pos.Fn, pos.Ln+1, ctx.DisplayName) + result // !!!!
		pos = ctx.ParentEntryPos
		ctx = ctx.Parent
	}

	return "Traceback (most recent call last): \n" + result
}

// LISTEN NODE
type ListenNode struct{}

// FUNCTION NODES
type FunctionDefNode struct {
	Name     string
	ArgNames []string
	Body     interface{}
}

type FunctionCallNode struct {
	FuncName string
	Args     []interface{}
}

// NEST NODES - CUSTOM DATA STRUCTURE
type NestDefNode struct {
	Name     string
	Fields   []string
	Methods  map[string]*FunctionDefNode
	PosStart *Position
	PosEnd   *Position
}

// SYMBOL TABLE
type SymbolTable struct {
	symbols map[string]interface{} // Dictionary to store symbols
	parent  *SymbolTable           // Pointer to parent symbol table
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]interface{}), // Initialize symbols map
		parent:  nil,
	}
}

// Set a value in the symbol table
func (s *SymbolTable) set(name string, value interface{}) {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return
	}
	fmt.Printf("Setting variable %s to %v in context\n", name, value)
	s.symbols[name] = value
}

// Get a value from the symbol table
func (s *SymbolTable) get(name string) interface{} {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return nil
	}
	if value, exists := s.symbols[name]; exists {
		fmt.Printf("Getting variable %s with value %v from context\n", name, value)
		return value
	} else if s.parent != nil {
		fmt.Printf("Looking up variable %s in parent context\n", name) // Added debug
		return s.parent.get(name)
	}
	fmt.Printf("Variable %s not found\n", name) // Added debug
	return nil
}

// Remove a symbol from the table
func (s *SymbolTable) remove(name string) {
	delete(s.symbols, name)
}

// INTERPRETER //
type Interpreter struct{}

func (i *Interpreter) visit(node interface{}, context *Context) *RTResult {
	switch node := node.(type) {
	case FunctionDefNode:
		return i.visitFunctionDefNode(node, context)
	case DotCallNode:
		return i.visitDotCallNode(node, context)
	case NestDefNode:
		return i.visitNestDefNode(node, context)
	case ListNode:
		return i.visitListNode(node, context)
	case FunctionCallNode:
		return i.visitFunctionCallNode(node, context)
	case ListenNode:
		return i.visitListenNode(node, context)
	case RoarNode:
		return i.visitRoarNode(node, context)
	case NumberNode:
		return i.visitNumberNode(node)
	case BinOpNode:
		return i.visitBinOpNode(node, context)
	case UnaryOpNode:
		return i.visitUnaryOpNode(node, context)
	case StringNode:
		return i.visitStringNode(node)
	case BoolNode:
		return i.visitBoolNode(node)
	case VarAccessNode:
		return i.visitVarAccessNode(node, context)
	case VarAssignNode:
		return i.visitVarAssignNode(node, context)
	case GrowlNode:
		return i.visitGrowlNode(node, context)
	case *PounceNode:
		return i.visitPounceNode(*node, context)
	case LeapNode:
		return i.visitLeapNode(node, context)
	case StatementsNode:
		return i.visitStatementsNode(node, context)
	case SniffbackNode: // return node
		val := i.visit(node.Value, context)
		res := NewRTResult()
		if val.Error != nil {
			return res.failure(val.Error)
		}
		return res.successWithReturn(val.Value)
	default:
		res := NewRTResult()
		return res.failure(fmt.Errorf("No visit method for node type %T", node))
	}
}

func (r *RTResult) successWithReturn(value interface{}) *RTResult {
	r.Value = value
	r.Error = nil
	return r // Acts as a return carrier
}

func (i *Interpreter) visitListNode(node ListNode, context *Context) *RTResult {
	res := NewRTResult()
	evaluated := []interface{}{}
	for _, el := range node.Elements {
		val := res.register(i.visit(el, context))
		if res.Error != nil {
			return res
		}
		evaluated = append(evaluated, val)
	}
	return res.success(evaluated)
}

func (i *Interpreter) visitNestDefNode(node NestDefNode, context *Context) *RTResult {
	res := NewRTResult()
	context.Symbol_Table.set(node.Name, node)
	return res.success(nil)
}

func (i *Interpreter) visitDotCallNode(node DotCallNode, context *Context) *RTResult {
	res := NewRTResult()
	var varName string

	if nodeTarget, ok := node.Target.(VarAccessNode); ok {
		varName = nodeTarget.Var_Name_Tok.Value
	}

	targetVal := res.register(i.visit(node.Target, context))
	if res.Error != nil {
		return res
	}

	// LIST METHODS
	if listVal, ok := targetVal.([]interface{}); ok {
		switch node.Method {
		case "sniff":
			val := res.register(i.visit(node.Args[0], context))
			if res.Error != nil {
				return res
			}
			listVal = append(listVal, val)

		case "howl":
			idxRaw := res.register(i.visit(node.Args[0], context))
			if res.Error != nil {
				return res
			}
			idxFloat, ok := idxRaw.(float64)
			if !ok {
				return res.failure(fmt.Errorf("Expected index to be a number"))
			}
			idx := int(idxFloat)
			if idx >= 0 && idx < len(listVal) {
				listVal = append(listVal[:idx], listVal[idx+1:]...)
			}

		case "wag":
			return res.success(float64(len(listVal)))

		case "prowl":
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			r.Shuffle(len(listVal), func(i, j int) {
				listVal[i], listVal[j] = listVal[j], listVal[i]
			})

		case "snarl":
			for i := 0; i < len(listVal)/2; i++ {
				listVal[i], listVal[len(listVal)-1-i] = listVal[len(listVal)-1-i], listVal[i]
			}

		case "lick":
			flat := []interface{}{}
			for _, v := range listVal {
				if nested, ok := v.([]interface{}); ok {
					flat = append(flat, nested...)
				} else {
					flat = append(flat, v)
				}
			}
			return res.success(flat)

		case "howl_at":
			thresholdRaw := res.register(i.visit(node.Args[0], context))
			if res.Error != nil {
				return res
			}
			threshold, ok := thresholdRaw.(float64)
			if !ok {
				return res.failure(fmt.Errorf("Expected threshold to be a number"))
			}
			filtered := []interface{}{}
			for _, v := range listVal {
				if val, ok := v.(float64); ok && val > threshold {
					filtered = append(filtered, val)
				}
			}
			return res.success(filtered)

		case "nest":
			sizeRaw := res.register(i.visit(node.Args[0], context))
			if res.Error != nil {
				return res
			}
			size, ok := sizeRaw.(float64)
			if !ok {
				return res.failure(fmt.Errorf("Expected chunk size to be a number"))
			}
			chunks := []interface{}{}
			sz := int(size)
			for i := 0; i < len(listVal); i += sz {
				end := i + sz
				if end > len(listVal) {
					end = len(listVal)
				}
				chunks = append(chunks, listVal[i:end])
			}
			return res.success(chunks)

		default:
			return res.failure(fmt.Errorf("Unknown list method: %s", node.Method))
		}

		if varName != "" {
			context.Symbol_Table.set(varName, listVal)
		}
		return res.success(nil)
	}

	// 🐚 NEST METHODS & FIELDS
	if inst, ok := targetVal.(map[string]interface{}); ok {
		// METHOD CALL: d.speak()
		if methods, ok := inst["__methods__"].(map[string]*FunctionDefNode); ok {
			if fn, ok := methods[node.Method]; ok {
				if len(fn.ArgNames) != len(node.Args) {
					return res.failure(fmt.Errorf("Method '%s' expects %d args, got %d", node.Method, len(fn.ArgNames), len(node.Args)))
				}

				methodCtx := &Context{
					DisplayName:  fn.Name,
					Parent:       context,
					Symbol_Table: NewSymbolTable(),
				}
				methodCtx.Symbol_Table.parent = context.Symbol_Table
				methodCtx.Symbol_Table.set("this", inst)

				for idx := 0; idx < len(fn.ArgNames); idx++ {
					argVal := res.register(i.visit(node.Args[idx], context))
					if res.Error != nil {
						return res
					}
					methodCtx.Symbol_Table.set(fn.ArgNames[idx], argVal)
				}

				val := res.register(i.visit(fn.Body, methodCtx))
				if res.Error != nil {
					return res
				}
				return res.success(val)
			}
		}

		// PROPERTY SET: d.name -> "Buddy"
		if len(node.Args) == 1 {
			val := res.register(i.visit(node.Args[0], context))
			if res.Error != nil {
				return res
			}
			inst[node.Method] = val
			return res.success(val)
		}

		// PROPERTY GET: roar d.name
		if val, exists := inst[node.Method]; exists {
			return res.success(val)
		}

		return res.failure(fmt.Errorf("Unknown property or method '%s'", node.Method))
	}

	return res.failure(fmt.Errorf("Only lists or nest instances support dot-access"))
}

func (i *Interpreter) visitStatementsNode(node StatementsNode, context *Context) *RTResult {
	res := NewRTResult()
	var lastValue interface{}

	for _, stmt := range node.Statements {
		result := i.visit(stmt, context)
		if result.Error != nil {
			return res.failure(result.Error)
		}

		// Only store the result if it's not sniffback or nil
		if result.Value != nil {
			lastValue = result.Value
		}
	}

	return res.success(lastValue)
}

func (i *Interpreter) visitRoarNode(node RoarNode, context *Context) *RTResult {
	res := NewRTResult()

	// If empty roar, print newline
	if node.Value == nil {
		fmt.Println()
		return res.success(nil)
	}

	values, ok := node.Value.([]interface{})
	if !ok {
		// Single value case (backward compatibility)
		value := res.register(i.visit(node.Value, context))
		if res.Error != nil {
			return res
		}
		fmt.Println(value)
		return res.success(nil)
	}

	outputs := []string{}
	for _, expr := range values {
		val := res.register(i.visit(expr, context))
		if res.Error != nil {
			return res
		}
		outputs = append(outputs, fmt.Sprint(val))
	}

	fmt.Println(strings.Join(outputs, " "))
	return res.success(nil)
}

func (i *Interpreter) visitGrowlNode(node GrowlNode, context *Context) *RTResult {
	res := NewRTResult()

	for _, caseBlock := range node.Cases {
		condition := res.register(i.visit(caseBlock.Condition, context))
		if res.Error != nil {
			return res
		}

		if condition.(bool) {
			return res.success(res.register(i.visit(caseBlock.Body, context)))
		}
	}

	if node.ElseCase != nil {
		return res.success(res.register(i.visit(node.ElseCase, context)))
	}

	return res.success(nil)
}

// Visit methods
func (i *Interpreter) visitVarAccessNode(node VarAccessNode, context *Context) *RTResult {
	res := NewRTResult()
	varName := node.Var_Name_Tok.Value
	value := context.Symbol_Table.get(varName)
	// Debug
	// fmt.Printf("Looking for variable: %s in context\n", varName)
	// fmt.Printf("Available variables: %v\n", context.Symbol_Table.symbols)

	if value == nil {
		return res.failure(fmt.Errorf("'%s' is not defined", varName))
	}
	return res.success(value)
}

func (i *Interpreter) visitVarAssignNode(node VarAssignNode, context *Context) *RTResult {
	res := NewRTResult()
	varName := node.Var_Name_Tok.Value
	value := res.register(i.visit(node.Value_Node, context))
	if res.Error != nil {
		return res
	}

	fmt.Printf("Assigning variable: %s -> %v\n", varName, value) // Debug output

	context.Symbol_Table.set(varName, value)
	return res.success(value)
}

func (i *Interpreter) visitListenNode(node ListenNode, context *Context) *RTResult {
	res := NewRTResult()
	fmt.Print("> ") // Optional: display a prompt
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return res.failure(fmt.Errorf("Failed to read input"))
	}
	return res.success(input)
}

// Using howl functions
func (i *Interpreter) visitFunctionDefNode(node FunctionDefNode, context *Context) *RTResult {
	res := NewRTResult()
	context.Symbol_Table.set(node.Name, node)
	return res.success(nil)
}

func (i *Interpreter) visitFunctionCallNode(node FunctionCallNode, context *Context) *RTResult {
	res := NewRTResult()

	fnVal := context.Symbol_Table.get(node.FuncName)
	if fnVal == nil {
		return res.failure(fmt.Errorf("Function or nest '%s' is not defined", node.FuncName))
	}

	// If it's a nest, instantiate it
	if nestDef, isNest := fnVal.(NestDefNode); isNest {
		instance := map[string]interface{}{}

		// Initialize all fields to nil
		for _, field := range nestDef.Fields {
			instance[field] = nil
		}

		// Attach method map and __type__
		instance["__methods__"] = nestDef.Methods
		instance["__type__"] = nestDef.Name

		return res.success(instance)
	}

	// 🐾 Otherwise, it's a normal function
	fnNode, ok := fnVal.(FunctionDefNode)
	if !ok {
		return res.failure(fmt.Errorf("'%s' is not a function or nest", node.FuncName))
	}

	if len(fnNode.ArgNames) != len(node.Args) {
		return res.failure(fmt.Errorf("Function '%s' expects %d arguments, got %d", node.FuncName, len(fnNode.ArgNames), len(node.Args)))
	}

	funcContext := &Context{
		DisplayName:  fnNode.Name,
		Parent:       context,
		Symbol_Table: NewSymbolTable(),
	}
	funcContext.Symbol_Table.parent = context.Symbol_Table

	for idx := 0; idx < len(fnNode.ArgNames); idx++ {
		argVal := res.register(i.visit(node.Args[idx], context))
		if res.Error != nil {
			return res
		}
		funcContext.Symbol_Table.set(fnNode.ArgNames[idx], argVal)
	}

	bodyRes := i.visit(fnNode.Body, funcContext)
	if bodyRes.Error != nil {
		return res.failure(bodyRes.Error)
	}
	return res.success(bodyRes.Value)
}

// Updated functions to use RTResult
func (i *Interpreter) visitNumberNode(node NumberNode) *RTResult {
	res := NewRTResult()

	// Evaluate the number based on its type
	if node.Tok.Type == TT_INT {
		val, err := strconv.Atoi(node.Tok.Value)
		if err != nil {
			return res.failure(err)
		}
		return res.success(float64(val))
	} else if node.Tok.Type == TT_FLOAT {
		val, err := strconv.ParseFloat(node.Tok.Value, 64)
		if err != nil {
			return res.failure(err)
		}
		return res.success(val)
	}
	return res.failure(fmt.Errorf("Invalid number type: %s", node.Tok.Type))
}

func (i *Interpreter) visitBinOpNode(node BinOpNode, context *Context) *RTResult {
	res := NewRTResult()

	// Evaluate the left and right nodes
	leftResult := i.visit(node.Left_Node, context)
	if leftResult.Error != nil {
		return res.failure(leftResult.Error)
	}

	rightResult := i.visit(node.Right_Node, context)
	if rightResult.Error != nil {
		return res.failure(rightResult.Error)
	}

	leftValue := leftResult.Value
	rightValue := rightResult.Value

	switch node.Op_Tok.Type {

	// Arithmetic Operators
	case TT_PLUS, TT_MINUS, TT_MUL, TT_DIV, TT_MOD, TT_EXP:
		leftFloat, okLeft := leftValue.(float64)
		rightFloat, okRight := rightValue.(float64)
		if !okLeft || !okRight {
			return res.failure(fmt.Errorf("Expected numbers for arithmetic operations"))
		}

		switch node.Op_Tok.Type {
		case TT_PLUS:
			return res.success(leftFloat + rightFloat)
		case TT_MINUS:
			return res.success(leftFloat - rightFloat)
		case TT_MUL:
			return res.success(leftFloat * rightFloat)
		case TT_DIV:
			if rightFloat == 0 {
				return res.failure(fmt.Errorf("Division by zero"))
			}
			return res.success(leftFloat / rightFloat)
		case TT_MOD:
			return res.success(math.Mod(leftFloat, rightFloat))
		case TT_EXP:
			return res.success(math.Pow(leftFloat, rightFloat))
		}

	// String Concatenation
	case TT_CONC:
		leftStr, okLeft := leftValue.(string)
		rightStr, okRight := rightValue.(string)
		if okLeft && okRight {
			return res.success(leftStr + rightStr)
		}
		return res.failure(fmt.Errorf("Cannot concatenate non-string types"))

	// Comparison Operators
	case TT_GT, TT_LT, TT_GTE, TT_LTE, TT_EQEQ, TT_NEQ:
		leftFloat, okLeft := leftValue.(float64)
		rightFloat, okRight := rightValue.(float64)
		if !okLeft || !okRight {
			return res.failure(fmt.Errorf("Expected numbers for comparison operations"))
		}

		switch node.Op_Tok.Type {
		case TT_GT:
			return res.success(leftFloat > rightFloat)
		case TT_LT:
			return res.success(leftFloat < rightFloat)
		case TT_GTE:
			return res.success(leftFloat >= rightFloat)
		case TT_LTE:
			return res.success(leftFloat <= rightFloat)
		case TT_EQEQ:
			return res.success(leftFloat == rightFloat)
		case TT_NEQ:
			return res.success(leftFloat != rightFloat)
		}

	// Logical Operators
	case TT_AND, TT_OR:
		leftBool, okLeft := leftValue.(bool)
		rightBool, okRight := rightValue.(bool)
		if !okLeft || !okRight {
			return res.failure(fmt.Errorf("Expected booleans for logical operations"))
		}

		if node.Op_Tok.Type == TT_AND {
			return res.success(leftBool && rightBool)
		} else {
			return res.success(leftBool || rightBool)
		}

	default:
		return res.failure(fmt.Errorf("Unknown operator: %s", node.Op_Tok.Value))
	}

	return res.failure(fmt.Errorf("Unexpected fallthrough in visitBinOpNode"))
}

func (i *Interpreter) visitUnaryOpNode(node UnaryOpNode, context *Context) *RTResult {
	res := NewRTResult()

	// Evaluate the operand
	valResult := i.visit(node.Node, context)
	if valResult.Error != nil {
		return res.failure(valResult.Error)
	}

	// Apply the unary operator
	switch node.Op_Tok.Type {
	case TT_POS:
		if val, ok := valResult.Value.(float64); ok {
			return res.success(+val)
		}
		return res.failure(fmt.Errorf("Unary positive operation requires a number"))
	case TT_NEG:
		if val, ok := valResult.Value.(float64); ok {
			return res.success(-val)
		}
		return res.failure(fmt.Errorf("Unary negation operation requires a number"))
	default:
		return res.failure(fmt.Errorf("Unknown unary operator: %s", node.Op_Tok.Type))
	}
}

func (i *Interpreter) visitStringNode(node StringNode) *RTResult {
	res := NewRTResult()
	return res.success(node.Tok.Value)
}

func (i *Interpreter) visitBoolNode(node BoolNode) *RTResult {
	res := NewRTResult()
	if node.Tok.Value == "true" {
		return res.success(true)
	}
	return res.success(false)
}

func (i *Interpreter) visitPounceNode(node PounceNode, context *Context) *RTResult {
	res := NewRTResult()

	for {
		// Evaluate loop condition
		condResult := res.register(i.visit(node.Condition, context))
		if res.Error != nil {
			return res
		}

		// Ensure condition is boolean
		condBool, ok := condResult.(bool) // FIX: No pointer
		if !ok {
			return res.failure(fmt.Errorf("Pounce condition must be a boolean, got %T", condResult))
		}

		// Stop if condition is false
		if !condBool {
			break
		}

		// Execute loop body
		for _, stmt := range node.Body {
			res.register(i.visit(stmt, context))
			if res.Error != nil {
				return res
			}
		}
	}

	return res.success(nil)
}

func (i *Interpreter) visitLeapNode(node LeapNode, context *Context) *RTResult {
	res := NewRTResult()

	startResult := res.register(i.visit(node.StartExpr, context))
	if res.Error != nil {
		return res
	}

	endResult := res.register(i.visit(node.EndExpr, context))
	if res.Error != nil {
		return res
	}

	// Ensure start and end values are numbers
	start, ok1 := startResult.(float64)
	end, ok2 := endResult.(float64)

	if !ok1 || !ok2 {
		return res.failure(fmt.Errorf("Expected numbers in leap range, got %T and %T", startResult, endResult))
	}

	for iter := int(start); iter < int(end); iter++ { // FIXED: Changed `i` to `iter`
		context.Symbol_Table.set(node.VarName.Value, float64(iter))

		res.register(i.visit(node.Body, context)) // FIXED: i.visit works now!
		if res.Error != nil {
			return res
		}
	}

	return res.success(nil)
}

func Run(text string, fn string) (interface{}, error) {
	return run(text, fn)
}

// RUN //
func run(text string, fn string) (interface{}, error) {
	// Create a new global symbol table
	globalSymbolTable := NewSymbolTable()

	// Make the context
	context := &Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	// Initialize the lexer and generate tokens
	lexer := NewLexer(fn, text)
	tokens, err := lexer.make_tokens()
	if err != nil {
		return nil, err
	}

	// Parse the tokens to generate the AST
	parser := NewParser(tokens)
	parseResult := parser.parse()

	if parseResult.Error != "" {
		return nil, fmt.Errorf(parseResult.Error)
	}

	// Create an interpreter with the SAME context
	interpreter := Interpreter{}
	result := interpreter.visit(parseResult.Node, context)

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Value, nil
}
