/*The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?*/

package main

import "fmt"
import "math"

func isPrime(x int64) bool {
	var i int64 = 2
	for ; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// WORKS WITH NUMBERS THAT HAVE a SQUARE ROOT
func Calculate(n int64) int64{
	var i int64 = int64(math.Sqrt(float64(n)))
	var result int64
	for ; i > 1; i-- {
		if n%i == 0 && isPrime(i) {
			result = i
			break
		}
	}
	return result
}

func main() {
	fmt.Println(Calculate(600851475143))
}
