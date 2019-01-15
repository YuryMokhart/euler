/*By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?*/

package main

import "fmt"
import "math"

func main() {
	fmt.Println(Calculate(10001))
}

func Calculate(n int)int{
	var i int = 3
	finish := 1
	var result int
	for finish < (n+1) {
		if isPrime(i) {
			finish++
		}
		if finish == n {
			result = i
			break
		}
		i++
	}
	return result
}

func isPrime(x int) bool {
	if x == 2 {
		return true
	}
	var i, z int = 2, int(math.Sqrt(float64(x)))
	for ; i <= z; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
