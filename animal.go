package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// CONSTANTS //

const DIGITS string = "0123456789"
const LETTERS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LETTERS_DIGITS string = LETTERS + DIGITS

// TOKENS //

// TokenType defines the type of token
type TokenType string

const (
	TT_INT      TokenType = "INT"    //
	TT_FLOAT    TokenType = "FLOAT"  //
	TT_BOOL     TokenType = "BOOL"   //
	TT_STRING   TokenType = "STRING" //
	TT_IDEN     TokenType = "IDENTIFIER"
	TT_KEY      TokenType = "KEYWORD"
	TT_PLUS     TokenType = "PLUS"  //
	TT_MINUS    TokenType = "MINUS" //
	TT_NEG      TokenType = "NEG"   //
	TT_POS      TokenType = "POS"   //
	TT_MUL      TokenType = "MUL"   //
	TT_DIV      TokenType = "DIV"   //
	TT_MOD      TokenType = "MOD"   //
	TT_EXP      TokenType = "EXP"   //
	TT_CONC     TokenType = "CONC"  //
	TT_EQ       TokenType = "EQ"
	TT_LROUNDBR TokenType = "LROUNDBR" //
	TT_RROUNDBR TokenType = "RROUNDBR" //
	TT_RSQRBR   TokenType = "RSQRBR"   //
	TT_LSQRBR   TokenType = "LSQRBR"   //
	TT_RCURLBR  TokenType = "RCURLBR"  //
	TT_LCURLBR  TokenType = "LCURLBR"  //
	TT_EOF      TokenType = "EOF"      //
)

var KEYWORDS = []string{
	"int", "float", "bool", "string",
}

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

func (t Token) matches(type_ TokenType, value string) bool {
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

func (l *Lexer) make_identifier() Token {
	idStr := ""
	Pos_Start := l.Pos.copy()

	for l.CurrentChar != 0 && strings.ContainsAny(string(l.CurrentChar), LETTERS_DIGITS+"_") {
		idStr += string(l.CurrentChar)
		l.advance()
	}
	tok_Type := TT_IDEN

	for _, keyword := range KEYWORDS {
		if idStr == keyword {
			tok_Type = TT_KEY
			break
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

// HIERARCHY //

// Exponentiation, highest precedence
func (p *Parser) power() *ParseResult {
	return p.bin_op(p.atom, []TokenType{TT_EXP})
}

// Handles parentheses and atoms (integers, booleans, strings)
func (p *Parser) atom() *ParseResult {
	res := &ParseResult{}
	tok := p.Current_Tok

	if tok.Type == TT_INT || tok.Type == TT_FLOAT {
		p.advance()
		return res.success(NumberNode{Tok: tok})
	} else if tok.Type == TT_BOOL {
		p.advance()
		return res.success(BoolNode{Tok: tok})
	} else if tok.Type == TT_STRING {
		p.advance()
		return res.success(StringNode{Tok: tok})
	} else if tok.Type == TT_IDEN {
		p.advance()
		return res.success(VarAccessNode{Var_Name_Tok: tok})
	} else if tok.Type == TT_LROUNDBR {
		p.advance()
		expr := p.expr()
		if expr.Error != "" {
			return res.failure(expr.Error)
		}
		if p.Current_Tok.Type == TT_RROUNDBR {
			p.advance()
			return res.success(expr.Node)
		} else {
			return res.failure("Expected ')'")
		}
	} else if tok.Type == TT_LSQRBR {
		p.advance()
		expr := p.expr()
		if expr.Error != "" {
			return res.failure(expr.Error)
		}
		if p.Current_Tok.Type == TT_RSQRBR {
			p.advance()
			return res.success(expr.Node)
		} else {
			return res.failure("Expected ']'")
		}
	} else if tok.Type == TT_LCURLBR {
		p.advance()
		expr := p.expr()
		if expr.Error != "" {
			return res.failure(expr.Error)
		}
		if p.Current_Tok.Type == TT_RCURLBR {
			p.advance()
			return res.success(expr.Node)
		} else {
			return res.failure("Expected '}'")
		}
	}
	return res.failure("Expected int, float, boolean, string, '+', '-', '(', '[', '{'")
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
	return p.power()
}

// Multiplication, division, modulo, and concatenation
func (p *Parser) term() *ParseResult {
	return p.bin_op(p.factor, []TokenType{TT_MUL, TT_DIV, TT_MOD, TT_CONC})
}

// Modified expr function to handle variable assignments with types and operations
func (p *Parser) expr() *ParseResult {
	res := &ParseResult{}

	// Handle variable access
	if p.Current_Tok.Type == TT_IDEN {
		var_name := p.Current_Tok
		p.advance()

		// Check for assignment (e.g., int -> 10)
		if p.Current_Tok.Type == TT_KEY && p.Current_Tok.matches(TT_KEY, "int") {
			p.advance()

			if p.Current_Tok.Type != TT_EQ {
				return res.failure(NewInvalidSyntaxError(
					p.Current_Tok.Pos_Start, p.Current_Tok.Pos_End,
					"Expected '->' after type",
				).asString())
			}
			p.advance()

			value_expr := res.register(p.expr())
			if res.Error != "" {
				return res
			}

			return res.success(VarAssignNode{Var_Name_Tok: var_name, Value_Node: value_expr})
		}

		// If it is not an assignment, treat it as a variable access
		return res.success(VarAccessNode{Var_Name_Tok: var_name})
	}

	// Handle binary operations and other expressions
	return p.bin_op(p.term, []TokenType{TT_PLUS, TT_MINUS})
}

// The bin_op function ensures the correct precedence for binary operations
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

// Get a value from the symbol table
func (s *SymbolTable) set(name string, value interface{}) {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return
	}
	fmt.Printf("Setting variable %s to %v in context\n", name, value)
	s.symbols[name] = value
}

func (s *SymbolTable) get(name string) interface{} {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return nil
	}
	value, exists := s.symbols[name]
	if exists {
		fmt.Printf("Getting variable %s with value %v from context\n", name, value)
		return value
	} else if s.parent != nil {
		return s.parent.get(name)
	}
	return nil
}

// Remove a symbol from the table
func (s *SymbolTable) remove(name string) {
	delete(s.symbols, name)
}

// INTERPRETER //
type Interpreter struct{}

func (i Interpreter) visit(node interface{}, context *Context) (interface{}, error) {
	/*  dynamically dispatch based on node type (<3 switch, thanks you for existing)
	    in Python you could probably make it a bit easier with strings, but i don't care, this works
		kinda
	*/
	switch node := node.(type) {
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
	default:
		return nil, fmt.Errorf("No visit method for node type %T", node)
	}
}

// Visit methods
func (i Interpreter) visitVarAccessNode(node VarAccessNode, context *Context) (interface{}, error) {
	var_name := node.Var_Name_Tok.Value
	value := context.Symbol_Table.get(var_name)
	if value == nil {
		return nil, fmt.Errorf("%s is not defined", var_name)
	}
	fmt.Printf("Accessing variable %s with value %v\n", var_name, value)
	return value, nil
}

func (i Interpreter) visitVarAssignNode(node VarAssignNode, context *Context) (interface{}, error) {
	var_name := node.Var_Name_Tok.Value
	value, err := i.visit(node.Value_Node, context)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Assigning variable %s with value %v\n", var_name, value)
	context.Symbol_Table.set(var_name, value)
	return value, nil
}

func (i Interpreter) visitNumberNode(node NumberNode) (interface{}, error) {
	// it is recommended to return float64 to avoid type mismatch
	if node.Tok.Type == TT_INT {
		val, err := strconv.Atoi(node.Tok.Value)
		if err != nil {
			return nil, err
		}
		return float64(val), nil
	} else if node.Tok.Type == TT_FLOAT {
		val, err := strconv.ParseFloat(node.Tok.Value, 64)
		if err != nil {
			return nil, err
		}
		return val, nil
	}
	return nil, fmt.Errorf("Invalid number type: %s", node.Tok.Type)
}

func (i Interpreter) visitBinOpNode(node BinOpNode, context *Context) (interface{}, error) {
	// Evaluate the left and right nodes
	leftVal, err := i.visit(node.Left_Node, context)
	if err != nil {
		return nil, err
	}

	rightVal, err := i.visit(node.Right_Node, context)
	if err != nil {
		return nil, err
	}

	// Apply the operator based on node.Op_Tok
	switch node.Op_Tok.Type {
	case TT_PLUS, TT_MINUS, TT_MUL, TT_DIV, TT_MOD, TT_EXP:
		// Handle arithmetic operations
		leftFloat, okLeft := leftVal.(float64)
		rightFloat, okRight := rightVal.(float64)
		if !okLeft || !okRight {
			return nil, fmt.Errorf("Expected numbers for arithmetic operations")
		}

		switch node.Op_Tok.Type {
		case TT_PLUS:
			return leftFloat + rightFloat, nil
		case TT_MINUS:
			return leftFloat - rightFloat, nil
		case TT_MUL:
			return leftFloat * rightFloat, nil
		case TT_DIV:
			if rightFloat == 0 {
				return nil, fmt.Errorf("Division by zero")
			}
			return leftFloat / rightFloat, nil
		case TT_MOD:
			return math.Mod(leftFloat, rightFloat), nil
		case TT_EXP:
			return math.Pow(leftFloat, rightFloat), nil
		}

	case TT_CONC:
		leftStr, okLeft := leftVal.(string)
		rightStr, okRight := rightVal.(string)
		if okLeft && okRight {
			return leftStr + rightStr, nil
		}
		return nil, fmt.Errorf("Cannot concatenate non-string types")
	default:
		return nil, fmt.Errorf("Unknown operator: %s", node.Op_Tok.Value)
	}

	return nil, fmt.Errorf("Unknown error")
}

func (i Interpreter) visitUnaryOpNode(node UnaryOpNode, context *Context) (interface{}, error) {
	// Evaluate the operand
	val, err := i.visit(node.Node, context)
	if err != nil {
		return nil, err
	}

	// Apply the unary operator
	switch node.Op_Tok.Type {
	case TT_POS:
		return +val.(float64), nil
	case TT_NEG:
		return -val.(float64), nil
	default:
		return nil, fmt.Errorf("Unknown unary operator: %s", node.Op_Tok.Type)
	}
}

func (i Interpreter) visitStringNode(node StringNode) (interface{}, error) {
	return node.Tok.Value, nil
}

func (i Interpreter) visitBoolNode(node BoolNode) (interface{}, error) {
	if node.Tok.Value == "true" {
		return true, nil
	}
	return false, nil
}

// RUN //

func run(text string, fn string) (interface{}, error) {
	global_symbol_table := NewSymbolTable()

	lexer := NewLexer(fn, text)
	tokens, err := lexer.make_tokens()
	if err != nil {
		return nil, err
	}
	// Generate AST
	parser := NewParser(tokens)
	res := parser.parse()

	if res.Error != "" {
		return nil, fmt.Errorf(res.Error)
	}

	interpreter := Interpreter{}
	context := &Context{
		DisplayName:  "<program>",
		Symbol_Table: global_symbol_table,
	}

	result, err := interpreter.visit(res.Node, context)
	if err != nil {
		return nil, err
	}

	return result, nil
}
