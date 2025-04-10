package animal_test

import (
	"animal/animal/v2"
	"reflect"
	"testing"
)

func TestList_SniffAppend(t *testing.T) {
	code := `
a -> [1, 2]
a.sniff(3)
a`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := []interface{}{float64(1), float64(2), float64(3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestList_WagLength(t *testing.T) {
	code := `
a -> [1, 2, 3, 4]
a.wag()`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if result != float64(4) {
		t.Errorf("Expected 4, got %v", result)
	}
}

func TestList_SnarlReverse(t *testing.T) {
	code := `
a -> [1, 2, 3]
a.snarl()
a`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := []interface{}{float64(3), float64(2), float64(1)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestList_ProwlShuffle(t *testing.T) {
	code := `
a -> [1, 2, 3]
a.prowl()
a`
	_, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error (shuffle shouldn't fail): %v", err)
	}
}

func TestList_LickFlatten(t *testing.T) {
	code := `
a -> [[1, 2], [3], 4]
a.lick()`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := []interface{}{float64(1), float64(2), float64(3), float64(4)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestList_HowlAtFilter(t *testing.T) {
	code := `
a -> [1, 3, 5, 7]
a.howl_at(4)`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := []interface{}{float64(5), float64(7)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestList_NestChunk(t *testing.T) {
	code := `
a -> [1, 2, 3, 4, 5]
a.nest(2)`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := []interface{}{
		[]interface{}{float64(1), float64(2)},
		[]interface{}{float64(3), float64(4)},
		[]interface{}{float64(5)},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
