package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repetead := Repeat("a", 5)
	expected := "aaaaa"

	if repetead != expected {
		t.Errorf("expected: %q but got: %q", expected, repetead);
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}