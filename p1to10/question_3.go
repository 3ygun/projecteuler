package p1to10

import "fmt"

/**

Largest prime factor


The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

**/
func Question3() {
	fmt.Printf("\tQ3: %d", calculate())
}

func calculate() int64 {
	var NUM int64 = 600851475143
	var test int64
	var largest int64

	// = unneeded because NUM not even
	for test = 2; test*test <= NUM; test++ {
		if NUM%test == 0 {
			if isPrime(test) {
				largest = test
			}
		}
	}

	return largest
}

func isPrime(num int64) bool {
	var i int64
	for i = 3; i <= num/2; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
