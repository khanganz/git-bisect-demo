package main

import (
	"testing"
)

func TestThatsBuggy(t *testing.T) {
	got := 1 + 1
	want := 3

	if got != want {
		t.Errorf("Expected %d, but got %d", got, want)
	}
}
