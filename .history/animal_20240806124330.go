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
	TT_INT    TokenType = "INT"
	TT_FLOAT  TokenType = "FLOAT"
	TT_PLUS   TokenType = "PLUS"
	TT_MINUS  TokenType = "MINUS"
	TT_MUL    TokenType = "MUL"
	TT_DIV    TokenType = "DIV"
	TT_MOD    TokenType = "MOD"
	TT_LPAREN TokenType = "LPAREN"
	TT_RPAREN TokenType = "RPAREN"
	TT_EOF    TokenType = "EOF" // End of file
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

	for l.CurrentChar != 0 {
		if l.CurrentChar == ' ' || l.CurrentChar == '\t' {
			l.advance()
		} else if strings.IndexByte(DIGITS, l.CurrentChar) != -1 {
			tokens = append(tokens, l.make_number())
		} else if l.peek(4) == "meow" {
			tokens = append(tokens, Token{Type: TT_PLUS, Value: "meow"})
			l.advanceBy(4)
		} else if l.peek(4) == "woof" {
			tokens = append(tokens, Token{Type: TT_MINUS, Value: "woof"})
			l.advanceBy(4)
		} else if l.peek(3) == "moo" {
			tokens = append(tokens, Token{Type: TT_MUL, Value: "moo"})
			l.advanceBy(3)
		} else if l.peek(5) == "drone" {
			tokens = append(tokens, Token{Type: TT_DIV, Value: "drone"})
			l.advanceBy(5)
		} else if l.peek(5) == "squeak" {
			tokens = append(tokens, Token{Type: TT_MOD, Value: "squeak"})
			l.advanceBy(5)
		} else if l.CurrentChar == '(' {
			tokens = append(tokens, Token{Type: TT_LPAREN, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == ')' {
			tokens = append(tokens, Token{Type: TT_RPAREN, Value: string(l.CurrentChar)})
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
// __init__

type NumberNode struct {
	Tok Token
}

// __repr__
func (n NumberNode) String() string {
	return fmt.Sprintf("(%s)", n.Tok)
}

type BinOpNode struct {
	Left_Node  int
	Op_Tok     Token
	Right_Node int
}

func (b BinOpNode) String() string {
	return fmt.Sprintf("(%d %s %d)", b.Left_Node, b.Op_Tok, b.Right_Node)
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

// func (p *Parser) factor() NumberNode {
// 	tok := p.Current_Tok

// 	if tok.Type == TT_INT || tok.Type == TT_FLOAT {
// 		p.advance()
// 		return NumberNode{Tok: tok}
// 	} else if tok.Type == TT_LPAREN {
// 		p.advance()
// 		expr := p.expr()
// 		if p.Current_Tok.Type == TT_RPAREN {
// 			p.advance()
// 			return expr
// 		} else {
// 			panic("Expected ')'")
// 		}
// 	}
// 	panic("Expected int or float or (")
// }

// func (p *Parser) term() NumberNode {
// 	left := p.factor()

// 	for p.Current_Tok.Type == TT_MUL || p.Current_Tok.Type == TT_DIV {
// 		op := p.Current_Tok
// 		p.advance()
// 		right := p.factor()
// 		left = BinOpNode{Left_Node: left, Op_Tok: op, Right_Node: right}
// 	}

// 	return left
// }

// func (p *Parser) expr() NumberNode {
// 	left := p.term()

// 	for p.Current_Tok.Type == TT_PLUS || p.Current_Tok.Type == TT_MINUS {
// 		op := p.Current_Tok
// 		p.advance()
// 		right := p.term()
// 		left = BinOpNode{Left_Node: left, Op_Tok: op, Right_Node: right}
// 	}

// 	return left
// }

// RUN //

func run(text string, fn string) ([]Token, error) {
	lexer := NewLexer(fn, text)
	tokens, err := lexer.make_tokens()
	return tokens, err
}
