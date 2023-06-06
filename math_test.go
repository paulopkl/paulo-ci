package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(15, 15)

	if total != 30 {
		t.Errorf("Invalid sum result: %d, expected value: %d", total, 30)
	}
}

func TestMinus(t *testing.T) {
	total := Minus(15, 7)

	if total != 8 {
		t.Errorf("Invalid sum result: %d, expected value: %d", total, 8)
	}
}