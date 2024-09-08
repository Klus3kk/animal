package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CONSTANTS //

const DIGITS string = "0123456789"

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

// TOKENS //

// TokenType defines the type of token
type TokenType string

const (
	TT_INT      TokenType = "INT"
	TT_FLOAT    TokenType = "FLOAT"
	TT_BOOL     TokenType = "BOOL"
	TT_STRING   TokenType = "STRING"
	TT_PLUS     TokenType = "PLUS"
	TT_MINUS    TokenType = "MINUS"
	TT_NEG      TokenType = "NEG"
	TT_POS      TokenType = "POS"
	TT_MUL      TokenType = "MUL"
	TT_DIV      TokenType = "DIV"
	TT_MOD      TokenType = "MOD"
	TT_EXP      TokenType = "EXP"
	TT_LROUNDBR TokenType = "LROUNDBR"
	TT_RROUNDBR TokenType = "RROUNDBR"
	TT_RSQRBR   TokenType = "RSQRBR"
	TT_LSQRBR   TokenType = "LSQRBR"
	TT_RCURLBR  TokenType = "RCURLBR"
	TT_LCURLBR  TokenType = "LCURLBR"
	TT_EOF      TokenType = "EOF" // End of file
)

// Token represents a token with its type and value
type Token struct {
	Type      TokenType
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
		if l.CurrentChar == ' ' || l.CurrentChar == '\t' {
			l.advance()
		} else if strings.IndexByte(DIGITS, l.CurrentChar) != -1 {
			tokens = append(tokens, l.make_number()) // Tokenize number
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
		} else if l.CurrentChar == '(' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_LROUNDBR, Value: string(l.CurrentChar), Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == ')' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_RROUNDBR, Value: string(l.CurrentChar), Pos_Start: posStart, Pos_End: posEnd})
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
		} else {
			// Handling illegal characters
			posStart := l.Pos.copy()
			char := string(l.CurrentChar)
			l.advance()
			err = fmt.Errorf("%s: Unexpected Character: '%s'", posStart.asString(), char)
			break
		}
	}
	tokens = append(tokens, Token{Type: TT_EOF, Value: "EOF", Pos_Start: l.Pos.copy(), Pos_End: l.Pos.copy()})
	return tokens, err
}

func (l *Lexer) advanceBy(count int) {
	for i := 0; i < count; i++ {
		l.advance()
	}
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

// NODES //
// __init__ (self, Tok)

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
	res := p.expr()
	if res.Error == "" && p.Current_Tok.Type != TT_EOF {
		return res.failure(NewInvalidSyntaxError(
			p.Current_Tok.Pos_Start, p.Current_Tok.Pos_End,
			"Expected 'meow', 'woof', 'moo', 'drone'",
		).asString())
	}
	return res
}

////////////

func (p *Parser) factor() *ParseResult {
	res := &ParseResult{}
	tok := p.Current_Tok

	if tok.Type == TT_POS || tok.Type == TT_NEG {
		// Advance the token, but don't register it as a *ParseResult
		p.advance()
		factor := res.register(p.factor())
		if res.Error != "" {
			return res
		}
		return res.success(UnaryOpNode{Op_Tok: tok, Node: factor})
	} else if tok.Type == TT_INT || tok.Type == TT_FLOAT {
		p.advance() // No need to register, just advance
		return res.success(NumberNode{Tok: tok})
	} else if tok.Type == TT_LROUNDBR {
		p.advance() // same here
		expr := p.expr()
		if expr.Error != "" {
			return res.failure(expr.Error)
		}
		if p.Current_Tok.Type == TT_RROUNDBR {
			p.advance() // -||-
			return res.success(expr.Node)
		} else {
			return res.failure("Expected ')'")
		}
	}
	return res.failure("Expected int, float, '+', '-', or '('")
}

func (p *Parser) term() *ParseResult {
	return p.bin_op(p.factor, []TokenType{TT_MUL, TT_DIV})
}

func (p *Parser) expr() *ParseResult {
	res := p.bin_op(p.term, []TokenType{TT_PLUS, TT_MINUS})

	// Check if there are multiple consecutive numbers without operators
	if res.Error == "" && p.Current_Tok.Type != TT_EOF {
		// If we have not reached EOF, check for invalid syntax
		if p.Current_Tok.Type == TT_INT || p.Current_Tok.Type == TT_FLOAT {
			return res.failure(NewInvalidSyntaxError(
				p.Current_Tok.Pos_Start, p.Current_Tok.Pos_End,
				"Unexpected consecutive numbers, expected operator",
			).asString())
		}
	}
	return res
}

////////////

func (p *Parser) bin_op(funcToCall func() *ParseResult, ops []TokenType) *ParseResult {
	res := &ParseResult{}
	left := res.register(funcToCall())
	if res.Error != "" {
		return res
	}

	for p.Current_Tok.Type != TT_EOF && contains(ops, p.Current_Tok.Type) {
		op_tok := p.Current_Tok
		p.advance()
		right := res.register(funcToCall())
		if res.Error != "" {
			return res
		}
		left = BinOpNode{Left_Node: left, Op_Tok: op_tok, Right_Node: right}
	}

	return res.success(left)
}

// Utility function to check if a TokenType is in the list
func contains(ops []TokenType, op TokenType) bool {
	for _, val := range ops {
		if val == op {
			return true
		}
	}
	return false
}

// RUN //

func run(text string, fn string) (interface{}, error) {
	// Generate Tokens
	lexer := NewLexer(fn, text)
	tokens, err := lexer.make_tokens()
	if err != nil { // if error: return None, error
		return nil, err
	}

	// Generate AST
	parser := NewParser(tokens)
	res := parser.expr()

	if res.Error != "" {
		return nil, fmt.Errorf(res.Error)
	}
	return res.Node, nil
}
