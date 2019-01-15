package main

import(
	"testing"
)

func TestSum(t *testing.T){
	var test = []struct {
		x int
		y int
	}{
		{100,44},
		{200,188},
		{300,188},
	}

	for _, check := range test{
		checkTotal := Sum(check.x)
		if checkTotal != check.y{
			t.Errorf("Expected result is %d, got %d", check.y, checkTotal)
		}
	}
}