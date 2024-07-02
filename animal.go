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
	Error_Name string
	Details    string
}

func (e Error) as_string() string {
	return fmt.Sprintf("(%s): (%s)", e.Error_Name, e.Details)
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
	Pos         int
	CurrentChar byte
}

func NewLexer(text string) *Lexer {
	lexer := &Lexer{Text: text, Pos: 0}
	if len(text) > 0 {
		lexer.CurrentChar = text[0]
	} else {
		lexer.CurrentChar = 0
	}
	return lexer
}

func (l *Lexer) advance() {
	l.Pos++
	if l.Pos < len(l.Text) {
		l.CurrentChar = l.Text[l.Pos]
	} else {
		l.CurrentChar = 0 // None
	}
}

func (l *Lexer) make_tokens() ([]Token, error) { // Updated signature to include return types
	tokens := []Token{}
	var err error

	for l.CurrentChar != 0 {
		if l.CurrentChar == ' ' || l.CurrentChar == '\t' {
			l.advance()
		} else if strings.IndexByte(DIGITS, l.CurrentChar) != -1 {
			tokens = append(tokens, l.make_number())
		} else if l.CurrentChar == '+' {
			tokens = append(tokens, Token{Type: TT_PLUS, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == '-' {
			tokens = append(tokens, Token{Type: TT_MINUS, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == '*' {
			tokens = append(tokens, Token{Type: TT_MUL, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == '/' {
			tokens = append(tokens, Token{Type: TT_DIV, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == '(' {
			tokens = append(tokens, Token{Type: TT_LPAREN, Value: string(l.CurrentChar)})
			l.advance()
		} else if l.CurrentChar == ')' {
			tokens = append(tokens, Token{Type: TT_RPAREN, Value: string(l.CurrentChar)})
			l.advance()
		} else {
			char := string(l.CurrentChar)
			l.advance()
			err = fmt.Errorf("Unexpected Character: '%s'", char)
			break
		}
	}
	return tokens, err
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

// RUN //

func run(text string) ([]Token, error) {
	lexer := NewLexer(text)
	tokens, err := lexer.make_tokens()
	return tokens, err
}
