package core

import (
	"fmt"
	"strconv"
	"strings"
)

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
		} else if l.CurrentChar == '"' || l.CurrentChar == '\'' {
			tokens = append(tokens, l.make_string()) // Tokenize string
		} else if l.peek(4) == "true" || l.peek(5) == "false" {
			tokens = append(tokens, l.make_boolean()) // Tokenize boolean
		} else if l.peek(4) == "meow" {
			posStart := l.Pos.copy() // Copy the current position as the token's start
			l.advanceBy(4)
			posEnd := l.Pos.copy() // After advancing, copy for the token's end
			tokens = append(tokens, Token{Type: TT_PLUS, Value: "PLUS", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(9) == "%bestiary" {
			posStart := l.Pos.copy()
			l.advanceBy(9)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_KEY, Value: "%bestiary", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(6) == "%debug" {
			posStart := l.Pos.copy()
			l.advanceBy(6)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_KEY, Value: "%debug", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(8) == "!shelter" {
			posStart := l.Pos.copy()
			l.advanceBy(8)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_KEY, Value: "!shelter", Pos_Start: posStart, Pos_End: posEnd})
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
		} else if strings.ContainsAny(string(l.CurrentChar), LETTERS+"_") {
			tokens = append(tokens, l.make_identifier()) // Tokenize letter
		} else if l.peek(2) == "->" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_EQ, Value: "EQ", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '.' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_DOT, Value: ".", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "*{" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: "TT_THROW_START", Value: "*{", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "}*" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: "TT_THROW_END", Value: "}*", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "*[" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: "TT_TRY_START", Value: "*[", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "]*" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: "TT_TRY_END", Value: "]*", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == "*(" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: "TT_CATCH_START", Value: "*(", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.peek(2) == ")*" {
			posStart := l.Pos.copy()
			l.advanceBy(2)
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: "TT_CATCH_END", Value: ")*", Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '*' {
			posStart := l.Pos.copy()
			l.advance()
			err = fmt.Errorf("%s: Unexpected standalone '*' - did you mean '*[', '*{', etc.?", posStart.asString())
			break
		} else if l.CurrentChar == ')' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_RROUNDBR, Value: string(l.CurrentChar), Pos_Start: posStart, Pos_End: posEnd})
		} else if l.CurrentChar == '(' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_LROUNDBR, Value: string(l.CurrentChar), Pos_Start: posStart, Pos_End: posEnd})
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
		} else if l.CurrentChar == ':' {
			posStart := l.Pos.copy()
			l.advance()
			posEnd := l.Pos.copy()
			tokens = append(tokens, Token{Type: TT_COLON, Value: ":", Pos_Start: posStart, Pos_End: posEnd})
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
	if Debug {
		for _, tok := range tokens {
			fmt.Println("TOKEN:", tok)
		}
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

	// Force underscore to be treated as a keyword
	if idStr == "_" {
		return Token{Type: TT_KEY, Value: "_", Pos_Start: Pos_Start, Pos_End: l.Pos}
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
	quoteChar := l.CurrentChar
	l.advance() // Skip opening quote
	strVal := ""

	//strVal = strings.ReplaceAll(strVal, `\n`, "\n")
	//strVal = strings.ReplaceAll(strVal, `\t`, "\t")
	//strVal = strings.ReplaceAll(strVal, `\\`, `\`)

	for l.CurrentChar != 0 && l.CurrentChar != quoteChar {
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
