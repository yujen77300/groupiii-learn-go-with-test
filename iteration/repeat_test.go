package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 3)
	fmt.Println(repeated)
	// Output: ccc
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// +++++++++++++++++++++++++++++
// More in strings package
// +++++++++++++++++++++++++++++

func TestContains(t *testing.T) {
	result := strings.Contains("hello world", "world")
	if !result {
		t.Errorf("Expected Contains('hello world', 'world') to be true, but got false")
	}

	result = strings.Contains("hello", "")
	if !result {
		t.Errorf("Expected Contains('hello', '') to be true, but got false")
	}
}

func TestCutPrefix(t *testing.T) {
	after, found := strings.CutPrefix("Golang", "Go")
	if !found {
		t.Errorf("There is no prefix")
	}
	if after != "lang" {
		t.Errorf("There is no prefix")
	}
}
