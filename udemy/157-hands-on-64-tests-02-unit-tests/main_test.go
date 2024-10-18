package main

import "testing"

func TestAdd(t *testing.T) {
	sum := add(1, 2)
	if sum != 3 {
		t.Errorf("Wrong Output")
	}
}


func TestSubtract(t *testing.T) {
	diff := subtract(2, 1)
	if diff != 1 {
		t.Errorf("Wrong Output")
	}
}