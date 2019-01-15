package main 

import(
"testing"
)

func Test7(t *testing.T){
	test := []struct{
		input int 
		output int
	}{
		{6,13},
	}
	for _, check := range test{
		checkTotal := Calculate(check.input)
		if checkTotal != check.output{
			t.Errorf("Expected %d, got %d", check.output, checkTotal)
		}
	}
}