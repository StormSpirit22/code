package main

import "testing"

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Fatal("xxx")
	}
}

func add(a, b int) int {
	return a + b
}
