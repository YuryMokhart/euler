/*If we list all the natural numbers below 10 that are multiples of 3 or 5,
 we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("The sum of all the multiples of 3 below 1000 should be 166833")
	fmt.Println(sum(3, 1000))
	fmt.Println("The sum of all the multiples of 5 below 1000 should be 99500")
	fmt.Println(sum(5, 1000))
	fmt.Println("The sum of all the multiples of 15 below 1000 should be 33165")
	fmt.Println(sum(15, 1000))
	fmt.Println("The sum of all the multiples of 3 or 5 below 1000 should be 233168")
	totalSum()
}

func sum(a1 int, an int) int {
	var d int // value that we add each time for calculation(we can avoid this var and use a1)
	var n int //number of variable
	//an is a value of the last element()
	an = an - 1 //do it deliberately, because of the conditions. So that last element isn't included
	d = a1
	n = ((a1 + an) / d) - 1 //substruct 1, because otherwise it is a number of the element out of the limit
	an = a1 + (n-1)*d
	sum := ((a1 + an) * n) / 2
	return sum
}

func totalSum() {
	fmt.Println(sum(3, 1000) + sum(5, 1000) - sum(15, 1000))
}