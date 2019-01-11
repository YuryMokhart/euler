/*By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?*/

package main

import "fmt"
import "math"

func main() {
	var i int64 = 3
	finish := 1
	for finish < 10002 {
		if isPrime(i) {
			finish++
		}
		if finish == 10001 {
			fmt.Println(i)
			break
		}
		i++
	}
}

func isPrime(x int64) bool {
	if x == 2 {
		return true
	}
	var i, z int64 = 2, int64(math.Sqrt(float64(x)))
	for ; i <= z; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
