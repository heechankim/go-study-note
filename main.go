package main

import "testing"

func TestLenForMap(t *testing.T) {
	v := map[string]int{"A": 1, "B":2}
	actual, expected := Len(v), 2
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func TestLenForString(t *testing.T) {
	v := "one"
	actual, expected := Len(v), 3
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}