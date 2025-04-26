package core

import "fmt"

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
		if Debug {
			fmt.Println("Parsing statement, current token:", p.Current_Tok)
		}

		stmt := res.register(p.statement())
		if res.Error != "" {
			return res
		}

		if Debug {
			fmt.Println("Parsed statement:", stmt)
		}

		statements = append(statements, stmt)
	}

	return res.success(StatementsNode{Statements: statements})
}

func (p *Parser) peek() Token {
	if p.Tok_Idx+1 < len(p.Tokens) {
		return p.Tokens[p.Tok_Idx+1]
	}
	return Token{Type: TT_EOF}
}

func (p *Parser) Parse() *ParseResult {
	return p.parse()
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

		// Check if it's a method call (with parentheses)
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
				Args:   nil, // property get or Set
			}
		}
	}

	return res.success(node)
}

func (p *Parser) fetch_csv_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "fetch_csv") {
		return res.failure("Expected 'fetch_csv'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_LROUNDBR {
		return res.failure("Expected '(' after 'fetch_csv'")
	}
	p.advance()

	filename := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	var sep interface{} = nil
	var header interface{} = nil

	if p.Current_Tok.Type == TT_COMMA {
		p.advance()
		sep = res.register(p.expr())
		if res.Error != "" {
			return res
		}

		if p.Current_Tok.Type == TT_COMMA {
			p.advance()
			header = res.register(p.expr())
			if res.Error != "" {
				return res
			}
		}
	}

	if p.Current_Tok.Type != TT_RROUNDBR {
		return res.failure("Expected ')' after arguments")
	}
	p.advance()

	return res.success(FetchCSVNode{
		Filename:  filename,
		Separator: sep,
		Header:    header,
	})
}

func (p *Parser) fetch_json_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "fetch_json") {
		return res.failure("Expected 'fetch_json'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_LROUNDBR {
		return res.failure("Expected '(' after 'fetch_json'")
	}
	p.advance()

	filename := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_RROUNDBR {
		return res.failure("Expected ')' after filename")
	}
	p.advance()

	return res.success(FetchJSONNode{Filename: filename})
}

func (p *Parser) sniff_file_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "sniff_file") {
		return res.failure("Expected 'sniff_file'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_LROUNDBR {
		return res.failure("Expected '(' after 'sniff_file'")
	}
	p.advance()

	filename := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_RROUNDBR {
		return res.failure("Expected ')' after filename")
	}
	p.advance()

	return res.success(SniffFileNode{Filename: filename})
}

func (p *Parser) drop_append_expr() *ParseResult {
	res := &ParseResult{}
	if !p.Current_Tok.matches(TT_KEY, "drop_append") {
		return res.failure("Expected 'drop_append'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_LROUNDBR {
		return res.failure("Expected '(' after 'drop_append'")
	}
	p.advance()

	filename := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_COMMA {
		return res.failure("Expected ',' after filename")
	}
	p.advance()

	content := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_RROUNDBR {
		return res.failure("Expected ')' after content")
	}
	p.advance()

	return res.success(DropAppendNode{Filename: filename, Content: content})
}

func (p *Parser) fetch_expr() *ParseResult {
	res := &ParseResult{}

	// fetch (
	if !p.Current_Tok.matches(TT_KEY, "fetch") {
		return res.failure("Expected 'fetch'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_LROUNDBR {
		return res.failure("Expected '(' after 'fetch'")
	}
	p.advance()

	arg := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_RROUNDBR {
		return res.failure("Expected ')' after fetch argument")
	}
	p.advance()

	return res.success(FetchNode{Filename: arg})
}

func (p *Parser) drop_expr() *ParseResult {
	res := &ParseResult{}

	// drop (
	if !p.Current_Tok.matches(TT_KEY, "drop") {
		return res.failure("Expected 'drop'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_LROUNDBR {
		return res.failure("Expected '(' after 'drop'")
	}
	p.advance()

	filename := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_COMMA {
		return res.failure("Expected ',' after filename")
	}
	p.advance()

	content := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_RROUNDBR {
		return res.failure("Expected ')' after content")
	}
	p.advance()

	return res.success(DropNode{Filename: filename, Content: content})
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

func (p *Parser) mimic_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "mimic") {
		return res.failure("Expected 'mimic'")
	}
	p.advance()

	expr := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_LCURLBR {
		return res.failure("Expected '{' after mimic expression")
	}
	p.advance()

	cases := []MimicCase{}
	var otherwise interface{}

	for p.Current_Tok.Type != TT_RCURLBR && p.Current_Tok.Type != TT_EOF {
		if p.Current_Tok.Type == TT_KEY && p.Current_Tok.Value == "_" {
			p.advance()
			if p.Current_Tok.Type != TT_EQ {
				return res.failure("Expected '->' after '_'")
			}
			p.advance()

			otherwise = res.register(p.expr())
			if res.Error != "" {
				return res
			}
			continue
		}

		match := res.register(p.atom())
		if res.Error != "" {
			return res
		}

		if p.Current_Tok.Type != TT_EQ {
			return res.failure("Expected '->' after match value")
		}
		p.advance()

		body := res.register(p.expr())
		if res.Error != "" {
			return res
		}

		cases = append(cases, MimicCase{
			MatchValue: match,
			Body:       body,
		})
	}

	if p.Current_Tok.Type != TT_RCURLBR {
		return res.failure("Expected '}' at end of mimic block")
	}
	p.advance()

	return res.success(MimicNode{
		Expr:      expr,
		Cases:     cases,
		Otherwise: otherwise,
	})
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
	if Debug {
		fmt.Println("Advanced token:", p.Current_Tok)
	}
	condition := res.register(p.bin_op(p.comp_expr, []string{TT_AND, TT_OR}))

	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != TT_LCURLBR { // Expecting '{'
		return res.failure("Expected '{' after growl condition")
	}
	p.advance()
	if Debug {
		fmt.Println("Advanced token:", p.Current_Tok)
	}

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
	if Debug {
		fmt.Println("Advanced token:", p.Current_Tok)
	}

	cases = append(cases, ConditionBlock{Condition: condition, Body: body})

	// Handle "sniff" (else-if)
	for p.Current_Tok.matches(TT_KEY, "sniff") {
		p.advance()
		if Debug {
			fmt.Println("Advanced token:", p.Current_Tok)
		}

		condition := res.register(p.bin_op(p.comp_expr, []string{TT_AND, TT_OR}))

		if res.Error != "" {
			return res
		}

		if p.Current_Tok.Type != TT_LCURLBR {
			return res.failure("Expected '{' after sniff condition")
		}
		p.advance()
		if Debug {
			fmt.Println("Advanced token:", p.Current_Tok)
		}

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
		if Debug {
			fmt.Println("Advanced token:", p.Current_Tok)
		}

		cases = append(cases, ConditionBlock{Condition: condition, Body: body})
	}

	// Handle "wag" (else)
	if p.Current_Tok.matches(TT_KEY, "wag") {
		p.advance()
		if Debug {
			fmt.Println("Advanced token:", p.Current_Tok)
		}

		if p.Current_Tok.Type != TT_LCURLBR {
			return res.failure("Expected '{' after wag")
		}
		p.advance()
		if Debug {
			fmt.Println("Advanced token:", p.Current_Tok)
		}

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
		if Debug {
			fmt.Println("Advanced token:", p.Current_Tok)
		}

	}

	return res.success(GrowlNode{Cases: cases, ElseCase: elseCase})
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
	// fmt.Println("call_expr: building call for", node)

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

	// Protect against block markers appearing inside expressions
	//if p.Current_Tok.Type == TT_TRY_START ||
	//	p.Current_Tok.Type == TT_TRY_END ||
	//	p.Current_Tok.Type == TT_CATCH_START ||
	//	p.Current_Tok.Type == TT_CATCH_END ||
	//	p.Current_Tok.Type == TT_THROW_START ||
	//	p.Current_Tok.Type == TT_THROW_END {
	//	return res.failure("Unexpected block marker in expression")
	//}

	if p.Current_Tok.Type == TT_EOF {
		return res.failure("Unexpected end of file")
	}

	// Handle special case keywords like roar, growl, etc.
	if tok.Type == TT_KEY {
		return p.expr()
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

		// Variable access + list index support
		node := VarAccessNode{Var_Name_Tok: tok}
		resultNode := interface{}(node)

		for p.Current_Tok.Type == TT_LSQRBR {
			p.advance()
			index := res.register(p.expr())
			if res.Error != "" {
				return res
			}
			if p.Current_Tok.Type != TT_RSQRBR {
				return res.failure("Expected ']' after index")
			}
			p.advance()

			resultNode = ListAccessNode{
				Target: resultNode,
				Index:  index,
			}
		}

		return p.parseDotCalls(resultNode)
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
			return res.failure("Expected '}' after expression")
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

func (p *Parser) statement() *ParseResult {
	// Handle block-level statements first
	if p.Current_Tok.Type == TT_TRY_START {
		return p.try_symbolic_expr()
	}
	if p.Current_Tok.Type == TT_THROW_START {
		return p.throw_symbolic_expr()
	}

	// Otherwise, fallback to normal expr (which includes roar, growl, fetch, drop, etc.)
	return p.expr()
}

// Modified expr function to handle variable assignments with types and operations
func (p *Parser) expr() *ParseResult {
	res := &ParseResult{}
	if p.Current_Tok.Type == TT_TRY_START {
		return p.try_symbolic_expr()
	}
	if p.Current_Tok.Type == TT_THROW_START {
		return p.throw_symbolic_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "%bestiary") {
		return p.bestiary_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "!shelter") {
		return p.shelter_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "%debug") {
		res := &ParseResult{}
		p.advance()
		return res.success(DebugNode{})
	}

	// Print statement
	if p.Current_Tok.matches(TT_KEY, "roar") {
		return p.roar_expr()
	}

	// Conditional statements
	if p.Current_Tok.matches(TT_KEY, "growl") {
		return p.growl_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "mimic") {
		return p.mimic_expr()
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
	if p.Current_Tok.matches(TT_KEY, "fetch") {
		return p.fetch_expr()
	}
	if p.Current_Tok.matches(TT_KEY, "drop") {
		return p.drop_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "drop_append") {
		return p.drop_append_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "sniff_file") {
		return p.sniff_file_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "fetch_json") {
		return p.fetch_json_expr()
	}
	if p.Current_Tok.matches(TT_KEY, "fetch_csv") {
		return p.fetch_csv_expr()
	}

	if p.Current_Tok.matches(TT_KEY, "whimper") {
		p.advance()
		return res.success(WhimperNode{})
	}
	if p.Current_Tok.matches(TT_KEY, "hiss") {
		p.advance()
		return res.success(HissNode{})
	}

	if p.Current_Tok.matches(TT_KEY, "howl_fail") {
		return p.howl_fail_expr()
	}

	// Handle variable access and assignment (with optional type hint)
	if p.Current_Tok.Type == TT_IDEN && (p.peek().Type == TT_COLON || p.peek().Type == TT_EQ) {
		varName := p.Current_Tok
		p.advance()

		var typeName *Token = nil

		if p.Current_Tok.Type == TT_COLON {
			p.advance()

			if p.Current_Tok.Type != TT_IDEN && p.Current_Tok.Type != TT_KEY {
				return res.failure("Expected type name after ':'")
			}
			tmp := p.Current_Tok
			typeName = &tmp
			p.advance()
		}

		if p.Current_Tok.Type != TT_EQ {
			return res.failure("Expected '->' after variable name (and optional type)")
		}
		p.advance()

		value_expr := res.register(p.expr())
		if res.Error != "" {
			return res
		}

		return res.success(VarAssignNode{
			Var_Name_Tok: varName,
			TypeName:     typeName,
			Value_Node:   value_expr,
		})
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

	// dot-assignment support
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

func (p *Parser) debug_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "%debug") {
		return res.failure("Expected '%debug'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_EQ {
		return res.failure("Expected '->' after '%debug'")
	}
	p.advance()

	value := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	return res.success(ShelterNode{Symbols: value})
}

func (p *Parser) shelter_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "!shelter") {
		return res.failure("Expected '!shelter'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_EQ {
		return res.failure("Expected '->' after '!shelter'")
	}
	p.advance()

	shelterList := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	return res.success(ShelterNode{Symbols: shelterList})
}

func (p *Parser) bestiary_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "%bestiary") {
		return res.failure("Expected '%bestiary'")
	}
	p.advance()

	if p.Current_Tok.Type != TT_STRING {
		return res.failure("Expected filename as string after %bestiary")
	}

	filename := StringNode{Tok: p.Current_Tok}
	p.advance()

	return res.success(BestiaryNode{Filename: filename})
}

func (p *Parser) howl_fail_expr() *ParseResult {
	res := &ParseResult{}

	if !p.Current_Tok.matches(TT_KEY, "howl_fail") {
		return res.failure("Expected 'howl_fail'")
	}
	p.advance()

	msg := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	return res.success(ThrowSymbolicNode{ErrorValue: msg})
}

func (p *Parser) try_symbolic_expr() *ParseResult {
	res := &ParseResult{}

	if p.Current_Tok.Type != TT_TRY_START {
		return res.failure("Expected '*[' to begin try block")
	}
	p.advance()

	for p.Current_Tok.Type == TT_EOF || (p.Current_Tok.Type == TT_KEY && p.Current_Tok.Value == "") {
		p.advance()
	}
	// COLLECT MULTIPLE STATEMENTS
	statements := []interface{}{}
	for p.Current_Tok.Type != TT_TRY_END && p.Current_Tok.Type != TT_EOF {
		stmt := res.register(p.statement())
		if res.Error != "" {
			return res
		}
		if stmt != nil { // <<< only append non-nil
			statements = append(statements, stmt)
		}
	}

	tryBody := StatementsNode{Statements: statements}

	if p.Current_Tok.Type != TT_TRY_END {
		return res.failure("Expected ']*' to close try block")
	}
	p.advance()

	if p.Current_Tok.Type != TT_CATCH_START {
		return res.failure("Expected '*(' to begin catch block")
	}
	p.advance()

	// NO errorName support, always default to "_error"
	errorName := "_error"

	// 🛠 COLLECT MULTIPLE STATEMENTS IN CATCH TOO
	catchStatements := []interface{}{}
	for p.Current_Tok.Type != TT_CATCH_END && p.Current_Tok.Type != TT_EOF {
		stmt := res.register(p.statement())
		if res.Error != "" {
			return res
		}
		if stmt != nil { // <<< only append non-nil
			catchStatements = append(catchStatements, stmt)
		}
	}

	catchBody := StatementsNode{Statements: catchStatements}

	if p.Current_Tok.Type != TT_CATCH_END {
		return res.failure("Expected ')*' to close catch block")
	}
	p.advance()

	return res.success(TrySymbolicNode{
		TryBody:   tryBody,
		CatchBody: catchBody,
		ErrorName: errorName,
	})
}

func (p *Parser) throw_symbolic_expr() *ParseResult {
	res := &ParseResult{}

	if p.Current_Tok.Type != "TT_THROW_START" {
		return res.failure("Expected '*{' to begin throw block")
	}
	p.advance()

	expr := res.register(p.expr())
	if res.Error != "" {
		return res
	}

	if p.Current_Tok.Type != "TT_THROW_END" {
		return res.failure("Expected '}*' to close throw block")
	}
	p.advance()

	return res.success(ThrowSymbolicNode{ErrorValue: expr})
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
	statements := []interface{}{}
	for p.Current_Tok.Type != TT_RCURLBR && p.Current_Tok.Type != TT_EOF {
		stmt := res.register(p.expr())
		if res.Error != "" {
			return res
		}
		statements = append(statements, stmt)
	}
	body := StatementsNode{Statements: statements}

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
		if Debug {
			fmt.Println("Found operator:", p.Current_Tok)
		}

		op_tok := p.Current_Tok
		p.advance()

		if Debug {
			fmt.Println("Current token after operator:", p.Current_Tok)
		}

		right := res.register(funcToCall())
		if res.Error != "" {
			return res
		}

		if Debug {
			fmt.Println("Right side parsed:", right)
		}

		left = BinOpNode{Left_Node: left, Op_Tok: op_tok, Right_Node: right}
	}

	return res.success(left)
}
