/*2520 is the smallest number that can be divided by each of the numbers
 from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers 
from 1 to 20? */

package main

func main() {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
    result := 1
    var divisor int
    for i := range primes {
        divisor = primes[i]
        for divisor <= 20 {
            divisor *= primes[i]
        }
        divisor /= primes[i]
        result *= divisor
    }
    println(result)
}