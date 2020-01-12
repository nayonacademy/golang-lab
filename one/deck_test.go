package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 9 {
		t.Errorf("Expected len of dec is 20 but got %v", len(d))
	}
	if d[0] != "green one" {
		t.Errorf("Expected first color is green but got %v", d[0])
	}
}
