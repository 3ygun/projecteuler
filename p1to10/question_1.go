package p1to10

import "fmt"

/**

Multiples of 3 and 5


If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

**/
func Question1() {
	total := 0
	to := 1000

	for i := 3; i < to; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	fmt.Printf("\tQ1: %d", total)
}
