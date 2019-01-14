package main

import(
"testing"
)

func TestCalculate(t *testing.T){
	if Calculate(13195) != 29{
		t.Error("The largest prime of 13195 is expected to be 29")
	}
}