package main

import (
	"fmt"
)

/*
Assignment
We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

Complete the sum function to return the sum of all inputs.

Take note of how the variadic inputs and the spread operator are used in the test suite.
*/

func sum(nums ...int) int {

	sumTotal := 0
	for i := 0; i < len(nums); i++ {
		sumTotal += nums[i]
	}

	return sumTotal
}

func test(nums ...int) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %v\n", total)
	fmt.Println("==========END REPORT==============")

}

func main() {
	test([]int{1, 2, 3, 4}...)
	test(2, 3, 4, 5)
}
