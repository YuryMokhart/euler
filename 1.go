/*If we list all the natural numbers below 10 that are multiples of 3 or 5,
 we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(3))
	fmt.Println(sum(5)) // not folowwing the task with value (5), because it includes 1000th element
	//works with operator "if"
	fmt.Println(sum(15))
	totalSum()
}

func sum(a1 int) int {
	var d int  // value that we add each time for calculation
	var n int  //number of variable
	an := 1000 //value of the last element()
	d = a1
	if a1 == 5 {
		n = (a1+an+1)/d - 1
	} else {
		n = (a1 + an) / d
	}
	an = a1 + (n-2)*d
	sum := ((a1 + an) * (n - 1)) / 2
	return sum
}

func totalSum() {
	fmt.Println(sum(3) + sum(5) - sum(15))
}
