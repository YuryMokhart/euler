package main

import(
"testing"
)

func TestDifference(t *testing.T){
	var test = []struct{
		input int
		output int
	}{
		{10,2640},
		{5,170},
	}
	for _, check := range test{
		checkTotal := Difference(check.input)
		if checkTotal != check.output{
			t.Errorf("Expected %d, got %d", check.output, checkTotal)
		}
	}
}