package main

import "strconv"

/**

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

**/
func problem1() string {
	total := 0
	to := 1000

	for i := 3; i < to; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return "Problem 1: " + strconv.Itoa(total)
}
