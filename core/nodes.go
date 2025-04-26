package core

import "fmt"

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
type BestiaryNode struct {
	Filename interface{}
}

type ShelterNode struct {
	Symbols interface{}
}

// TRY/CATCH/THROW
type TrySymbolicNode struct {
	TryBody   interface{}
	CatchBody interface{}
	ErrorName string
}

type ThrowSymbolicNode struct {
	ErrorValue interface{}
}

// WHIMPER/HISS (break, continue)
type WhimperNode struct{}
type HissNode struct{}

// MIMIC (switch)
type MimicNode struct {
	Expr      interface{}
	Cases     []MimicCase
	Otherwise interface{}
}

type MimicCase struct {
	MatchValue interface{}
	Body       interface{}
}

// SNIFFBACK (return)
type SniffbackNode struct {
	Value interface{}
}

// GROWLNODE (if, else if, else)
type GrowlNode struct {
	Cases    []ConditionBlock // List of conditions and bodies
	ElseCase interface{}      // Optional else case
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

// __init__ (self, Tok)
type StringNode struct {
	Tok Token
}

type BoolNode struct {
	Tok Token
}

type NumberNode struct {
	Tok Token
}

type BinOpNode struct {
	Left_Node  interface{}
	Op_Tok     Token
	Right_Node interface{}
}

type UnaryOpNode struct {
	Op_Tok Token
	Node   interface{}
}

type VarAccessNode struct {
	Var_Name_Tok Token
}

type VarAssignNode struct {
	Var_Name_Tok Token
	TypeName     *Token
	Value_Node   interface{}
}

// Statement Node for multiple top-level expressions
type StatementsNode struct {
	Statements []interface{}
}

// ListNode
type ListNode struct {
	Elements []interface{}
}

// ListAccessNode
type ListAccessNode struct {
	Target interface{}
	Index  interface{}
}

// File I/O Nodes
type FetchJSONNode struct {
	Filename interface{}
}

type FetchCSVNode struct {
	Filename  interface{}
	Separator interface{}
	Header    interface{}
}

type SniffFileNode struct {
	Filename interface{}
}
type FetchNode struct {
	Filename interface{}
}
type DropNode struct {
	Filename interface{}
	Content  interface{}
}

type DropAppendNode struct {
	Filename interface{}
	Content  interface{}
}

// DotCall

type DotCallNode struct {
	Target interface{}
	Method string
	Args   []interface{}
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

// DEBUG
type DebugNode struct {
	Value interface{}
}

// STRING OUTPUTS

func (n SniffbackNode) String() string {
	return fmt.Sprintf("(SNIFFBACK %v)", n.Value)
}

func (n GrowlNode) String() string {
	return fmt.Sprintf("(GROWL %v ELSE %v)", n.Cases, n.ElseCase)
}

func (n RoarNode) String() string {
	return fmt.Sprintf("(ROAR %v)", n.Value)
}

func (n StringNode) String() string {
	return fmt.Sprintf("(%s: %s)", n.Tok.Type, n.Tok.Value)
}

func (n BoolNode) String() string {
	return fmt.Sprintf("(%s: %s)", n.Tok.Type, n.Tok.Value)
}

func (n NumberNode) String() string {
	return fmt.Sprintf("(%s: %s)", n.Tok.Type, n.Tok.Value)
}

func (b BinOpNode) String() string { // __repr__
	return fmt.Sprintf("(%s %s %s)", b.Left_Node, b.Op_Tok.Value, b.Right_Node)
}

func (u UnaryOpNode) String() string { // __repr__
	return fmt.Sprintf("(%s %s)", u.Op_Tok.Type, u.Node)
}

func (n VarAssignNode) String() string {
	return fmt.Sprintf("(VarAssignNode %s -> %v)", n.Var_Name_Tok.Value, n.Value_Node)
}

func (s StatementsNode) String() string {
	return fmt.Sprintf("(%s)", s.Statements)
}
