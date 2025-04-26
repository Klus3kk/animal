package core

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// INTERPRETER //
type Interpreter struct{}

func (i *Interpreter) visit(node interface{}, context *Context) *RTResult {
	switch node := node.(type) {
	case BestiaryNode:
		return i.visitBestiaryNode(node, context)
	case DebugNode:
		Debug = true
		fmt.Println("Debug mode: true")
		return NewRTResult().success(nil)
	case ShelterNode:
		return i.visitShelterNode(node, context)
	case FunctionDefNode:
		return i.visitFunctionDefNode(node, context)
	case DotCallNode:
		return i.visitDotCallNode(node, context)
	case NestDefNode:
		return i.visitNestDefNode(node, context)
	case ListNode:
		return i.visitListNode(node, context)
	case ThrowSymbolicNode:
		return i.visitThrowSymbolicNode(node, context)
	case TrySymbolicNode:
		return i.visitTrySymbolicNode(node, context)
	case ListAccessNode:
		return i.visitListAccessNode(node, context)
	case FunctionCallNode:
		return i.visitFunctionCallNode(node, context)
	case DropNode:
		return i.visitDropNode(node, context)
	case FetchNode:
		return i.visitFetchNode(node, context)
	case DropAppendNode:
		return i.visitDropAppendNode(node, context)
	case SniffFileNode:
		return i.visitSniffFileNode(node, context)
	case FetchJSONNode:
		return i.visitFetchJSONNode(node, context)
	case FetchCSVNode:
		return i.visitFetchCSVNode(node, context)
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
	case MimicNode:
		return i.visitMimicNode(node, context)
	case WhimperNode:
		return i.visitWhimperNode(node, context)
	case HissNode:
		return i.visitHissNode(node, context)
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

func (i *Interpreter) visitDebugNode(node DebugNode, context *Context) *RTResult {
	res := NewRTResult()
	value := res.register(i.visit(node.Value, context))
	if res.Error != nil {
		return res
	}

	boolValue, ok := value.(bool)
	if !ok {
		return res.failure(fmt.Errorf("%debug expects a boolean value (true or false)"))
	}

	Debug = boolValue
	return res.success(nil)
}

func (i *Interpreter) visitShelterNode(node ShelterNode, context *Context) *RTResult {
	res := NewRTResult()

	// Special case: setting Debug mode
	if debugVal, ok := node.Symbols.(bool); ok {
		Debug = debugVal
		fmt.Println("Debug mode:", Debug)
		return res.success(nil)
	}

	shelterList := res.register(i.visit(node.Symbols, context))
	if res.Error != nil {
		return res
	}

	context.Symbol_Table.Set("__shelter__", shelterList)
	return res.success(nil)
}

func (i *Interpreter) visitMimicNode(node MimicNode, context *Context) *RTResult {
	res := NewRTResult()

	// Evaluate the expression being matched
	value := res.register(i.visit(node.Expr, context))
	if res.Error != nil {
		return res
	}

	// Try to match with each case
	for _, mc := range node.Cases {
		matchVal := res.register(i.visit(mc.MatchValue, context))
		if res.Error != nil {
			return res
		}

		if fmt.Sprint(matchVal) == fmt.Sprint(value) {
			return res.success(res.register(i.visit(mc.Body, context)))
		}
	}

	// Run default case if present
	if node.Otherwise != nil {
		return res.success(res.register(i.visit(node.Otherwise, context)))
	}

	return res.success(nil)
}

func (i *Interpreter) visitFetchJSONNode(node FetchJSONNode, context *Context) *RTResult {
	res := NewRTResult()

	filenameVal := res.register(i.visit(node.Filename, context))
	if res.Error != nil {
		return res
	}

	filename, ok := filenameVal.(string)
	if !ok {
		return res.failure(fmt.Errorf("fetch_json expects filename to be a string"))
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return res.failure(fmt.Errorf("Failed to read file: %v", err))
	}

	var parsed interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return res.failure(fmt.Errorf("Invalid JSON: %v", err))
	}

	return res.success(parsed)
}

func (i *Interpreter) visitFetchCSVNode(node FetchCSVNode, context *Context) *RTResult {
	res := NewRTResult()

	filenameVal := res.register(i.visit(node.Filename, context))
	if res.Error != nil {
		return res
	}
	filename, ok := filenameVal.(string)
	if !ok {
		return res.failure(fmt.Errorf("fetch_csv expects filename to be a string"))
	}

	// Default: comma, header = true
	sep := ','
	header := true

	if node.Separator != nil {
		sepRaw := res.register(i.visit(node.Separator, context))
		if res.Error != nil {
			return res
		}
		sepStr, ok := sepRaw.(string)
		if !ok || len(sepStr) != 1 {
			return res.failure(fmt.Errorf("Separator must be a single character string"))
		}
		sep = rune(sepStr[0])
	}

	if node.Header != nil {
		headerRaw := res.register(i.visit(node.Header, context))
		if res.Error != nil {
			return res
		}
		header, _ = headerRaw.(bool)
	}

	file, err := os.Open(filename)
	if err != nil {
		return res.failure(fmt.Errorf("Could not open CSV file: %v", err))
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = sep

	rows, err := reader.ReadAll()
	if err != nil {
		return res.failure(fmt.Errorf("Could not read CSV: %v", err))
	}

	var result []interface{}
	if header && len(rows) > 0 {
		keys := rows[0]
		for _, row := range rows[1:] {
			entry := map[string]interface{}{}
			for i := range keys {
				if i < len(row) {
					entry[keys[i]] = row[i]
				}
			}
			result = append(result, entry)
		}
	} else {
		for _, row := range rows {
			r := []interface{}{}
			for _, val := range row {
				r = append(r, val)
			}
			result = append(result, r)
		}
	}

	return res.success(result)
}

func (i *Interpreter) visitSniffFileNode(node SniffFileNode, context *Context) *RTResult {
	res := NewRTResult()
	filenameVal := res.register(i.visit(node.Filename, context))
	if res.Error != nil {
		return res
	}

	filename, ok := filenameVal.(string)
	if !ok {
		return res.failure(fmt.Errorf("sniff_file expects a string filename"))
	}

	if _, err := os.Stat(filename); err == nil {
		return res.success(true)
	}
	return res.success(false)
}

func (i *Interpreter) visitDropAppendNode(node DropAppendNode, context *Context) *RTResult {
	res := NewRTResult()
	filenameVal := res.register(i.visit(node.Filename, context))
	if res.Error != nil {
		return res
	}
	contentVal := res.register(i.visit(node.Content, context))
	if res.Error != nil {
		return res
	}
	filename, ok := filenameVal.(string)
	if !ok {
		return res.failure(fmt.Errorf("drop_append expects filename to be a string"))
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return res.failure(fmt.Errorf("could not open file for appending: %v", err))
	}
	defer file.Close()
	if _, err := file.WriteString(fmt.Sprint(contentVal)); err != nil {
		return res.failure(fmt.Errorf("could not append to file: %v", err))
	}
	return res.success(nil)
}

func (i *Interpreter) visitFetchNode(node FetchNode, context *Context) *RTResult {
	res := NewRTResult()
	filenameVal := res.register(i.visit(node.Filename, context))
	if res.Error != nil {
		return res
	}

	filename, ok := filenameVal.(string)
	if !ok {
		return res.failure(fmt.Errorf("fetch expects a string filename"))
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return res.failure(fmt.Errorf("could not read file: %v", err))
	}

	return res.success(string(content))
}

func (i *Interpreter) visitDropNode(node DropNode, context *Context) *RTResult {
	res := NewRTResult()
	filenameVal := res.register(i.visit(node.Filename, context))
	if res.Error != nil {
		return res
	}

	contentVal := res.register(i.visit(node.Content, context))
	if res.Error != nil {
		return res
	}

	filename, ok := filenameVal.(string)
	if !ok {
		return res.failure(fmt.Errorf("drop expects filename to be string"))
	}

	err := os.WriteFile(filename, []byte(fmt.Sprint(contentVal)), 0644)
	if err != nil {
		return res.failure(fmt.Errorf("could not write file: %v", err))
	}

	return res.success(nil)
}

func (i *Interpreter) visitListAccessNode(node ListAccessNode, context *Context) *RTResult {
	res := NewRTResult()
	target := res.register(i.visit(node.Target, context))
	if res.Error != nil {
		return res
	}
	index := res.register(i.visit(node.Index, context))
	if res.Error != nil {
		return res
	}

	idxFloat, ok := index.(float64)
	if !ok {
		return res.failure(fmt.Errorf("List index must be a number"))
	}
	idx := int(idxFloat)

	switch t := target.(type) {
	case []interface{}:
		if idx < 0 || idx >= len(t) {
			return res.failure(fmt.Errorf("List index out of bounds"))
		}
		return res.success(t[idx])
	default:
		return res.failure(fmt.Errorf("Target is not a list: %T", t))
	}
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
	context.Symbol_Table.Set(node.Name, node)
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
			context.Symbol_Table.Set(varName, listVal)
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
				methodCtx.Symbol_Table.Set("this", inst)

				for idx := 0; idx < len(fn.ArgNames); idx++ {
					argVal := res.register(i.visit(node.Args[idx], context))
					if res.Error != nil {
						return res
					}
					methodCtx.Symbol_Table.Set(fn.ArgNames[idx], argVal)
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

		// If it's a control signal, return it...
		if result.Value == "__whimper__" || result.Value == "__hiss__" {
			return res.success(result.Value)
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
			// New child scope for the body
			childCtx := &Context{
				DisplayName:  "growl-block",
				Parent:       context,
				Symbol_Table: NewSymbolTable(),
			}
			childCtx.Symbol_Table.parent = context.Symbol_Table

			return res.success(res.register(i.visit(caseBlock.Body, childCtx)))
		}
	}

	if node.ElseCase != nil {
		// New child scope for the else block
		childCtx := &Context{
			DisplayName:  "wag-block",
			Parent:       context,
			Symbol_Table: NewSymbolTable(),
		}
		childCtx.Symbol_Table.parent = context.Symbol_Table

		return res.success(res.register(i.visit(node.ElseCase, childCtx)))
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

	if Debug {
		fmt.Printf("Assigning variable: %s -> %v\n", varName, value)
	}

	context.Symbol_Table.Set(varName, value)
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
	context.Symbol_Table.Set(node.Name, node)
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
		funcContext.Symbol_Table.Set(fnNode.ArgNames[idx], argVal)
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
		condBool, ok := condResult.(bool)
		if !ok {
			return res.failure(fmt.Errorf("Pounce condition must be a boolean, got %T", condResult))
		}

		// Stop if condition is false
		if !condBool {
			break
		}

		// Scoped execution of loop body
		for _, stmt := range node.Body {
			childCtx := &Context{
				DisplayName:  "pounce-block",
				Parent:       context,
				Symbol_Table: NewSymbolTable(),
			}
			childCtx.Symbol_Table.parent = context.Symbol_Table

			result := i.visit(stmt, childCtx)
			if result.Error != nil {
				return result
			}
			if result.Value == "__whimper__" {
				return res.success(nil) // break the loop
			}
			if result.Value == "__hiss__" {
				break // skip rest of body
			}

		}
	}

	return res.success(nil)
}

func (i *Interpreter) visitWhimperNode(node WhimperNode, context *Context) *RTResult {
	res := NewRTResult()
	return res.success("__whimper__")
}

func (i *Interpreter) visitHissNode(node HissNode, context *Context) *RTResult {
	res := NewRTResult()
	return res.success("__hiss__")
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

	start, ok1 := startResult.(float64)
	endVal, ok2 := endResult.(float64)
	if !ok1 || !ok2 {
		return res.failure(fmt.Errorf("Expected numbers in leap range, got %T and %T", startResult, endResult))
	}

	for iter := int(start); iter < int(endVal); iter++ {
		childCtx := &Context{
			DisplayName:  "leap-iteration",
			Parent:       context,
			Symbol_Table: NewSymbolTable(),
		}
		childCtx.Symbol_Table.parent = context.Symbol_Table
		childCtx.Symbol_Table.Set(node.VarName.Value, float64(iter))

		result := i.visit(node.Body, childCtx)
		if result.Error != nil {
			return result
		}
		if result.Value == "__whimper__" {
			break
		}
		if result.Value == "__hiss__" {
			continue
		}
	}

	return res.success(nil)
}

func (i *Interpreter) visitBestiaryNode(node BestiaryNode, context *Context) *RTResult {
	res := NewRTResult()

	filenameVal := res.register(i.visit(node.Filename, context))
	if res.Error != nil {
		return res
	}

	filename, ok := filenameVal.(string)
	if !ok {
		return res.failure(fmt.Errorf("%bestiary expects a string filename"))
	}

	if importedFiles[filename] {
		return res.success(nil) // Already imported
	}
	importedFiles[filename] = true

	data, err := os.ReadFile(filename)
	if err != nil {
		return res.failure(fmt.Errorf("Failed to read bestiary file: %v", err))
	}

	code := string(data)

	_, err = CustomRun(code, filename, context)
	if err != nil {
		return res.failure(fmt.Errorf("Error running bestiary file: %v", err))
	}
	shelterRaw := context.Symbol_Table.get("__shelter__")
	if shelterRaw != nil {
		if shelterList, ok := shelterRaw.([]interface{}); ok {
			newSymbols := make(map[string]interface{})
			for _, name := range shelterList {
				if strName, ok := name.(string); ok {
					val := context.Symbol_Table.get(strName)
					if val != nil {
						newSymbols[strName] = val
					}
				}
			}
			context.Symbol_Table.symbols = newSymbols
		}
	}

	return res.success(nil)
}

func (i *Interpreter) visitTrySymbolicNode(node TrySymbolicNode, context *Context) *RTResult {
	res := NewRTResult()

	tryRes := i.visit(node.TryBody, context)
	if tryRes.Error == nil {
		return tryRes
	}

	// Create a new child context for the catch block
	catchCtx := &Context{
		DisplayName:  "catch-block",
		Parent:       context,
		Symbol_Table: NewSymbolTable(),
	}
	catchCtx.Symbol_Table.parent = context.Symbol_Table
	catchCtx.Symbol_Table.Set(node.ErrorName, tryRes.Error.Error())

	catchVal := i.visit(node.CatchBody, catchCtx)
	if catchVal.Error != nil {
		return catchVal
	}
	return res.success(catchVal.Value)
}

func (i *Interpreter) visitThrowSymbolicNode(node ThrowSymbolicNode, context *Context) *RTResult {
	res := NewRTResult()
	errVal := res.register(i.visit(node.ErrorValue, context))
	if res.Error != nil {
		return res
	}
	return res.failure(fmt.Errorf("%v", errVal)) // Throw the error
}
