/*2520 is the smallest number that can be divided by each of the numbers
 from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers 
from 1 to 20? */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Calculate(20))
}

func Calculate(n int) int {
	number := n
	result := 1
	for i := 1; i < number+1; i++ {
		result = LCM(result, i)
	}
	return result
}

func GCD(x int, y int) int {
	a := x
	b := y
	var remainder int = 0
	if a < b {
		for b != 0 {
			remainder = a % b
			a = b
			b = remainder
		}
		return a
	} else {
		for a != 0 {
			remainder = b % a
			b = a
			a = remainder
		}
	}
	return b
}

func LCM(x int, y int) int {
	m := x
	n := y
	result := (m * n) / GCD(m, n)
	return result
}
