package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repetead := Repeat("a")
	expected := "aaaaa"

	if repetead != expected {
		t.Errorf("expected: %q but got: %q", expected, repetead);
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}