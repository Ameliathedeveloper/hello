package hello

import "testing"

// test for greetings variable
func TestGreeting(t *testing.T) {
	want := "Hello, world."
	if got := Greeting; got != want {
		t.Errorf("Greeting = %q, want %q", got, want)
	}
}
