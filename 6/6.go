/*The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers
 and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers
 and the square of the sum.*/

package main

import "fmt"

func main() {
	fmt.Println(Difference(100))
}

func Difference(n int) int{
	first := 1
	var last int
	last = n
	difference := 0
	var sumOfSqrs int
	var sqrOfSum int

	sumOfSqrs = (n * (n + 1) * (2*n + 1)) / 6
	sqrOfSum = (((first + last) * n) / 2) * (((first + last) * n) / 2)
	difference = sqrOfSum - sumOfSqrs

	return difference
}
