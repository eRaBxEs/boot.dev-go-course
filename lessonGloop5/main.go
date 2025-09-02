package main

import (
	"fmt"
)

/*
As an Easter egg, we decided to reward our users with a free text message if they send a prime number of text messages this year.

Complete the printPrimes function. It should print all of the prime numbers up to and including max. It should skip any numbers that are not prime.

Breakdown
We skip even numbers because they can't be prime
We only check up to the square root of n. A factor higher than the square root of n must multiply with a factor lower than the square root of n, meaning we only need to check up to the square root of n for potential factors.
In your code, you can set the stop condition as i * i <= n
We start checking at 2 because 1 is not prime
*/

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if !isPrime {
			continue
		}

		fmt.Println(n)

	}
}

func printPrimes2BruteForce(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		factorCount := 0
		for i := 1; i <= n; i++ {
			if factorCount >= 3 {
				break
			}
			if n%i == 0 {
				factorCount++
			}
		}
		// prime numbers can only be divided by 1 and itself; meaning the number of factors should only be 2.
		if factorCount == 2 {
			fmt.Println(n)
		}

	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}
func test2(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes2BruteForce(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)

	test2(10)
	test2(20)
	test2(30)

}
