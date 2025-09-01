package main

import (
	"fmt"
)

/*
Assignment
Complete the maxMessages function. Given a cost threshold, it should calculate the maximum number of messages that can be sent.

Each message costs 100 pennies, plus an additional fee. The fee structure is:

1st message: 100 + 0
2nd message: 100 + 1
3rd message: 100 + 2
4th message: 100 + 3
*/

func maxMessages(thresh int) int {
	// ?
	var totalCost int
	const fee int = 100
	for i := 0; ; i++ {
		totalCost += fee + i
		if totalCost > thresh {
			return i
		}
	}
}

func test(thresh int) {
	fmt.Printf("For threshold %v \n", thresh)
	num := maxMessages(thresh)
	fmt.Printf("Maximum number of messages that can be sent: %d\n", num)
	fmt.Println("=====================================================================")
}

func main() {
	test(100)
	test(250)
	test(600)
	test(1000)
}
