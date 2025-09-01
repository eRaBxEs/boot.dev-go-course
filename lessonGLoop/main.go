package main

import "fmt"

/*
Assignment
At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send. Complete the bulkSend() function.

It should return the total cost (as a float64) to send a batch of numMessages messages. Each message costs 1.0, plus an additional fee. The fee structure is:

1st message: 1.0 + 0.00
2nd message: 1.0 + 0.01
3rd message: 1.0 + 0.02
4th message: 1.0 + 0.03
...
Use a loop to calculate the total cost and return it.
*/

func bulkSend(numMessages int) float64 {
	// ?
	const fee float64 = 1.0
	total := 0.00

	for i := 0; numMessages > i; i++ {
		total += fee + (0.01 * float64(i))

	}

	return total
}

func test(numMessages int) {
	fmt.Printf("sending %v bulk messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("bulk send complete! Cost: %.2f\n", cost)
	fmt.Println("=====================================================================")
}

func main() {
	test(1)
	test(2)
	test(3)
	test(4)
	test(10)
}
