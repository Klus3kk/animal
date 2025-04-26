package tests

import (
	"animal/core"
	"testing"
)

func parseInput(t *testing.T, input string) *core.ParseResult {
	t.Helper()
	lexer := core.NewLexer("<stdin>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		t.Fatalf("Lexer error: %v", err)
	}

	parser := core.NewParser(tokens)
	res := parser.Parse()
	if res.Error != "" {
		t.Fatalf("Parser error: %v", res.Error)
	}
	return res
}

func TestParser_ParseRoar(t *testing.T) {
	res := parseInput(t, "roar 5")

	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode, got %T", res.Node)
	}
}

func TestParser_ParseGrowlSniffWag(t *testing.T) {
	code := `
growl true {
    roar "yes"
}
sniff false {
    roar "maybe"
}
wag {
    roar "no"
}
`
	res := parseInput(t, code)

	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode for growl-sniff-wag")
	}
}

func TestParser_ParsePounce(t *testing.T) {
	code := `
pounce 5 > 0 {
    roar "loop"
}
`
	res := parseInput(t, code)

	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode for pounce")
	}
}

func TestParser_ParseLeap(t *testing.T) {
	code := `
leap i from 0 to 5 {
    roar i
}
`
	res := parseInput(t, code)

	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode for leap")
	}
}

func TestParser_ParseFunctionDefinition(t *testing.T) {
	code := `
howl greet(name) {
    roar "hello", name
}
`
	res := parseInput(t, code)

	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode for function")
	}
}

func TestParser_ParseSniffback(t *testing.T) {
	code := `
howl double(x) {
    x meow x sniffback
}
`
	res := parseInput(t, code)

	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode with sniffback inside")
	}
}

func TestParser_ParseNestDefinition(t *testing.T) {
	code := `
nest Dog {
    name
    howl bark() {
        roar "woof"
    }
}
`
	res := parseInput(t, code)

	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode for nest")
	}
}

func TestParser_ParseMimic(t *testing.T) {
	code := `
mimic 5 {
    1 -> roar "one"
    2 -> roar "two"
    _ -> roar "other"
}
`
	res := parseInput(t, code)

	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode for mimic")
	}
}

func TestParser_ParseFileOperations(t *testing.T) {
	codes := []string{
		`fetch("file.txt")`,
		`drop("file.txt", "data")`,
		`drop_append("file.txt", "append")`,
		`sniff_file("file.txt")`,
		`fetch_json("file.json")`,
		`fetch_csv("file.csv")`,
	}

	for _, code := range codes {
		res := parseInput(t, code)
		if _, ok := res.Node.(core.StatementsNode); !ok {
			t.Errorf("Expected StatementsNode for %s", code)
		}
	}
}

func TestParser_ParseTryCatchThrow(t *testing.T) {
	code := `
*[ 
    throw *{ "error" }*
]* *(
    roar "caught"
)*
`
	res := parseInput(t, code)
	if _, ok := res.Node.(core.StatementsNode); !ok {
		t.Errorf("Expected StatementsNode for try-catch-throw")
	}
}
