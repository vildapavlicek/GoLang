package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T)  {
	buffer := bytes.Buffer{}
	Greet("Chris", &buffer)

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("Got '%s'; want '%s'", got, want)
	}

}
