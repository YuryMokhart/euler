package main

import(
"testing"
)

func TestDifference(t *testing.T){
	if Difference(10) != 2640{
		t.Error("Expected that the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640")
	}
}