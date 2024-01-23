package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Daan")
	want := "hello Daan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
