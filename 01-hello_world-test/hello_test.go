package helloworldtest_test

import (
	"testing"

	"go-with-test/01-hello_world"
)

func TestHello(t *testing.T) {
	want := "Hello, world";
	got := helloworld.Hello("world", "")

	if want != got {
		t.Fatalf("Error: want %q, got %q", want, got)
	}

}