package hello

import "testing"

func TestGreeting(t *testing.T) {
    want := "Hello, world."
    if got := Greeting; got != want {
        t.Errorf("Greeting = %q, want %q", got, want)
    }
}