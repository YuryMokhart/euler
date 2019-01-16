package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	test := []struct {
		input  int
		output int
	}{
		{10, 2520},
		{4, 12},
		{5, 60},
	}
	for _, check := range test {
		checkTotal := Calculate(check.input)
		if checkTotal != check.output {
			t.Errorf("Incorrect result for No. %d. Expected %d, got %d", check.input, check.output, checkTotal)
		}
	}
}
