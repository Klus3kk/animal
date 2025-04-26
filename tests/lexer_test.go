package tests

import (
	"animal/core"
	"testing"
)

func TestLexer_TokenizeNumbers(t *testing.T) {
	lexer := core.NewLexer("<stdin>", "123 45.67")
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	if tokens[0].Type != core.TT_INT {
		t.Errorf("Expected INT token, got %s", tokens[0].Type)
	}
	if tokens[1].Type != core.TT_FLOAT {
		t.Errorf("Expected FLOAT token, got %s", tokens[1].Type)
	}
}

func TestLexer_TokenizeBooleans(t *testing.T) {
	lexer := core.NewLexer("<stdin>", "true false")
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	if tokens[0].Type != core.TT_BOOL || tokens[0].Value != "true" {
		t.Errorf("Expected BOOL 'true', got %v", tokens[0])
	}
	if tokens[1].Type != core.TT_BOOL || tokens[1].Value != "false" {
		t.Errorf("Expected BOOL 'false', got %v", tokens[1])
	}
}

func TestLexer_TokenizeStrings(t *testing.T) {
	lexer := core.NewLexer("<stdin>", `"hello" 'world'`)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	if tokens[0].Type != core.TT_STRING || tokens[0].Value != "hello" {
		t.Errorf("Expected STRING 'hello', got %v", tokens[0])
	}
	if tokens[1].Type != core.TT_STRING || tokens[1].Value != "world" {
		t.Errorf("Expected STRING 'world', got %v", tokens[1])
	}
}

func TestLexer_TokenizeArithmeticOperators(t *testing.T) {
	input := "5 meow 2 woof 1 moo 3 drone 4 squeak 5 soar 2 purr 3"
	lexer := core.NewLexer("<stdin>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	expectedTypes := []string{
		core.TT_INT, core.TT_PLUS, core.TT_INT, core.TT_MINUS, core.TT_INT,
		core.TT_MUL, core.TT_INT, core.TT_DIV, core.TT_INT, core.TT_MOD, core.TT_INT,
		core.TT_EXP, core.TT_INT, core.TT_CONC, core.TT_INT,
	}

	for i, expected := range expectedTypes {
		if tokens[i].Type != expected {
			t.Errorf("At token %d: expected %s, got %s", i, expected, tokens[i].Type)
		}
	}
}

func TestLexer_TokenizeComparisonOperators(t *testing.T) {
	input := "> < >= <= == !="
	lexer := core.NewLexer("<stdin>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	expectedTypes := []string{
		core.TT_GT, core.TT_LT, core.TT_GTE, core.TT_LTE, core.TT_EQEQ, core.TT_NEQ,
	}

	for i, expected := range expectedTypes {
		if tokens[i].Type != expected {
			t.Errorf("At token %d: expected %s, got %s", i, expected, tokens[i].Type)
		}
	}
}

func TestLexer_TokenizeLogicalOperators(t *testing.T) {
	input := "and or"
	lexer := core.NewLexer("<stdin>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	if tokens[0].Type != core.TT_AND {
		t.Errorf("Expected TT_AND token, got %s", tokens[0].Type)
	}
	if tokens[1].Type != core.TT_OR {
		t.Errorf("Expected TT_OR token, got %s", tokens[1].Type)
	}
}

func TestLexer_TokenizeKeywords(t *testing.T) {
	input := "growl sniff wag pounce leap roar howl nest listen sniffback fetch drop drop_append sniff_file fetch_json fetch_csv whimper hiss mimic _"
	lexer := core.NewLexer("<stdin>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	expectedKeywords := []string{
		"growl", "sniff", "wag", "pounce", "leap", "roar", "howl", "nest", "listen", "sniffback",
		"fetch", "drop", "drop_append", "sniff_file", "fetch_json", "fetch_csv",
		"whimper", "hiss", "mimic", "_",
	}

	for i, expected := range expectedKeywords {
		if tokens[i].Type != core.TT_KEY || tokens[i].Value != expected {
			t.Errorf("At token %d: expected keyword %s, got %v", i, expected, tokens[i])
		}
	}
}

func TestLexer_TokenizeSpecialSymbols(t *testing.T) {
	input := "(){}[],.:"
	lexer := core.NewLexer("<stdin>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	expectedTypes := []string{
		core.TT_LROUNDBR, core.TT_RROUNDBR,
		core.TT_LCURLBR, core.TT_RCURLBR,
		core.TT_LSQRBR, core.TT_RSQRBR,
		core.TT_COMMA, core.TT_DOT, core.TT_COLON,
	}

	for i, expected := range expectedTypes {
		if tokens[i].Type != expected {
			t.Errorf("At token %d: expected %s, got %s", i, expected, tokens[i].Type)
		}
	}
}

func TestLexer_TokenizeAssignment(t *testing.T) {
	lexer := core.NewLexer("<stdin>", "x -> 5")
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	if tokens[1].Type != core.TT_EQ {
		t.Errorf("Expected EQ token '->', got %s", tokens[1].Type)
	}
}

func TestLexer_TokenizeTryCatchThrow(t *testing.T) {
	input := "*[ ]* *( )* *{ }*"
	lexer := core.NewLexer("<stdin>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	expectedTypes := []string{
		"TT_TRY_START", "TT_TRY_END",
		"TT_CATCH_START", "TT_CATCH_END",
		"TT_THROW_START", "TT_THROW_END",
	}

	for i, expected := range expectedTypes {
		if tokens[i].Type != expected {
			t.Errorf("At token %d: expected %s, got %s", i, expected, tokens[i].Type)
		}
	}
}

func TestLexer_ErrorInvalidCharacter(t *testing.T) {
	lexer := core.NewLexer("<stdin>", "@")
	_, err := lexer.MakeTokens()
	if err == nil {
		t.Fatalf("Expected error for invalid character, got none")
	}
}
