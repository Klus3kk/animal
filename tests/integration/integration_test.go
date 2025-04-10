package animal_test

import (
	"animal/animal/v2"
	"testing"
)

func TestIntegration_FizzBuzz(t *testing.T) {
	code := `
leap i from 1 to 6 {
    growl i squeak 4 == 0 {
        roar "FizzBuzz"
    } sniff i squeak 3 == 0 {
        roar "Fizz"
    } sniff i squeak 5 == 0 {
        roar "Buzz"
    } wag {
        roar i
    }
}`
	_, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("FizzBuzz error: %v", err)
	}
}

func TestIntegration_Factorial(t *testing.T) {
	code := `
howl fact(n) {
    growl n <= 1 {
        1 sniffback
    } wag {
        n moo fact(n woof 1) sniffback
    }
}
fact(5)`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Factorial error: %v", err)
	}
	expected := float64(120)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIntegration_ListOps(t *testing.T) {
	code := `
a -> [1, 2]
a.sniff(3)
a.snarl()
a`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("ListOps error: %v", err)
	}
	if result == nil || len(result.([]interface{})) != 3 {
		t.Errorf("Expected reversed list, got %v", result)
	}
}

func TestIntegration_NestAndFunction(t *testing.T) {
	code := `
nest Dog {
    name
    howl greet() {
        "hello " purr this.name sniffback
    }
}

howl greetDog(d) {
    d.greet() sniffback
}

d -> Dog()
d.name -> "Barky"
greetDog(d)`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Nest+Func error: %v", err)
	}
	expected := "hello Barky"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
