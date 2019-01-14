package main

import(
	"testing"
)

func TestSum(t *testing.T){
	if Sum(90) != 44{
		t.Error("Expected result is 44")
	}
}