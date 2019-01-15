package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	testTables := []struct {
		input  int64
		output int64
	}{
		{13195, 29},
		{100, 5},
		{25, 5},
	}

	for _, check := range testTables {
		checkTotal := Calculate(check.input)
		if checkTotal != check.output {
			t.Errorf("The largest prime of %d is expected to be %d, got %d", check.input, check.output, checkTotal)
		}
	}
}
