// animal.go
package main

import (
	"fmt"
)

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
