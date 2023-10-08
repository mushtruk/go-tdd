package greet_buffer

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ihor")
	
	got := buffer.String()
	want := "Hello, Ihor"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}