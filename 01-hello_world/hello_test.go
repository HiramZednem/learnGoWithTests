package abs

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hiram")
	want := "Hello, Hiram"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}