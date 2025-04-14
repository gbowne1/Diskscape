package main

import "testing"

func TestSomething(t *testing.T) {
	if 2+2 != 4 {
		t.Errorf("Math is broken")
	}
}
