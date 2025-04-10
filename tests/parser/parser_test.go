package animal_test

import (
	"animal/animal/v2"
	"testing"
)

func parseTestHelper(t *testing.T, input string) {
	lexer := animal.NewLexer("<test>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}
	parser := animal.NewParser(tokens)
	result := parser.Parse()
	if result.Error != "" {
		t.Fatalf("Parser error: %s", result.Error)
	}
}

func TestParser_VariableAssign(t *testing.T) {
	parseTestHelper(t, `x -> 5`)
}

func TestParser_ArithmeticExpr(t *testing.T) {
	parseTestHelper(t, `2 meow 3 moo 4`)
}

func TestParser_FunctionDefinition(t *testing.T) {
	parseTestHelper(t, `
howl add(a, b) {
    a meow b sniffback
}
`)
}

func TestParser_ListLiteral(t *testing.T) {
	parseTestHelper(t, `mylist -> [1, 2, 3]`)
}

func TestParser_ConditionalFull(t *testing.T) {
	parseTestHelper(t, `
x -> 10
growl x > 5 {
    roar "Big"
} sniff x == 5 {
    roar "Equal"
} wag {
    roar "Small"
}`)
}

func TestParser_LeapLoop(t *testing.T) {
	parseTestHelper(t, `
leap i from 0 to 3 {
    roar i
}`)
}

func TestParser_PounceLoop(t *testing.T) {
	parseTestHelper(t, `
x -> 0
pounce x < 2 {
    x -> x meow 1
}`)
}

func TestParser_NestDefinition(t *testing.T) {
	parseTestHelper(t, `
nest Dog {
    name
    howl speak() {
        roar this.name
    }
}`)
}

func TestParser_DotCallProperty(t *testing.T) {
	parseTestHelper(t, `
dog -> Dog()
dog.name -> "Rex"
roar dog.name`)
}

func TestParser_SniffbackReturn(t *testing.T) {
	parseTestHelper(t, `
howl one() {
    1 sniffback
}`)
}
