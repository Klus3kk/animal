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
	Type  TokenType
	Value string
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
			tokens = append(tokens, l.make_number())
		} else if l.peek(4) == "meow" {
			tokens = append(tokens, Token{Type: TT_PLUS, Value: "PLUS"})
			l.advanceBy(4)
		} else if l.peek(4) == "woof" {
			tokens = append(tokens, Token{Type: TT_MINUS, Value: "MINUS"})
			l.advanceBy(4)
		} else if l.peek(3) == "moo" {
			tokens = append(tokens, Token{Type: TT_MUL, Value: "MUL"})
			l.advanceBy(3)
		} else if l.peek(5) == "drone" {
			tokens = append(tokens, Token{Type: TT_DIV, Value: "DIV"})
			l.advanceBy(5)
		} else if l.peek(6) == "squeak" {
			tokens = append(tokens, Token{Type: TT_MOD, Value: "MOD"})
			l.advanceBy(6)
		} else if l.peek(4) == "soar" {
			tokens = append(tokens, Token{Type: TT_EXP, Value: "EXP"})
			l.advanceBy(4)
		} else if l.CurrentChar == '(' {
			tokens = append(tokens, Token{Type: TT_LROUNDBR, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == ')' {
			tokens = append(tokens, Token{Type: TT_RROUNDBR, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == '[' {
			tokens = append(tokens, Token{Type: TT_LSQRBR, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == ']' {
			tokens = append(tokens, Token{Type: TT_RSQRBR, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == '{' {
			tokens = append(tokens, Token{Type: TT_LCURLBR, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == '}' {
			tokens = append(tokens, Token{Type: TT_RCURLBR, Value: string(l.CurrentChar)})
			l.advance()
		} else {
			posStart := l.Pos.copy()
			char := string(l.CurrentChar)
			l.advance()
			err = fmt.Errorf("%s: Unexpected Character: '%s'", posStart.asString(), char)
			break
		}
	}
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
		return Token{Type: TT_INT, Value: strconv.Itoa(intValue)}
	} else {
		floatValue, _ := strconv.ParseFloat(numStr, 64) // Convert numStr to float64
		return Token{Type: TT_FLOAT, Value: strconv.FormatFloat(floatValue, 'f', -1, 64)}
	}
}

// NODES //
// __init__ (self, tok)

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

func (b BinOpNode) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left_Node, b.Op_Tok.Value, b.Right_Node)
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

////////////

func (p *Parser) parse() interface{} {
	res := p.expr()
	return res
}

func (p *Parser) factor() interface{} {
	tok := p.Current_Tok

	if tok.Type == TT_INT || tok.Type == TT_FLOAT {
		p.advance()
		return NumberNode{Tok: tok}
	} else if tok.Type == TT_LROUNDBR {
		p.advance()
		expr := p.expr()
		if p.Current_Tok.Type == TT_RROUNDBR {
			p.advance()
			return expr
		} else {
			panic("Expected ')'")
		}
	}
	panic("Expected int or float or (")
}

func (p *Parser) term() interface{} {
	return p.bin_op(p.factor, []TokenType{TT_MUL, TT_DIV})
}

func (p *Parser) expr() interface{} {
	return p.bin_op(p.term, []TokenType{TT_PLUS, TT_MINUS})
}

////////////

func (p *Parser) bin_op(Func func() interface{}, ops []TokenType) interface{} {
	left := Func()

	for p.Current_Tok.Type != TT_EOF && contains(ops, p.Current_Tok.Type) {
		op_tok := p.Current_Tok
		p.advance()
		right := Func()
		left = BinOpNode{Left_Node: left, Op_Tok: op_tok, Right_Node: right}
	}

	return left
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
	if err != nil {
		return nil, err
	}

	// Generate AST
	parser := NewParser(tokens)
	ast := parser.expr()

	return ast, nil
}
