//unit tests focus on the smallest piece of testable code (e.g., a single function or method) in isolation
package hello

import "testing"

// test for greetings variable
func TestGreeting(t *testing.T) {
	want := "Hello, world."
	if got := Greeting; got != want {
		t.Errorf("Greeting = %q, want %q", got, want)
	}
}
