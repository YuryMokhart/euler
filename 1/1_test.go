package main 

import(
"testing"
)

func Test1(t *testing.T) {
	if sum(2,10) != 20{
		t.Error("The sum of all the multiples of 2 below 10 EXPECTED to be 20")
	}
 }