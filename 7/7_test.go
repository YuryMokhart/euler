package main 

import(
"testing"
)

func Test7(t *testing.T){
	if isPrime(9) == true{
		t.Error("9 is NOT a prime number")
	}
}