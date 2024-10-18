package main

import "testing"

func TestAdd(t *testing.T) {
    sum := Add(5, 5)
    if sum != 10 {
        t.Errorf("Sum was incorrect, got: %d, expected: %d.", sum, 10)
    }
}