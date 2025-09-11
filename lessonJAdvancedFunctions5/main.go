package main

import "fmt"

/*
Assignment
Complete the printReports function. It takes as input a sequence of messages, intro, body, outro. It should call printCostReport once for each message.

For each call of printCostReport, give it an anonymous function that returns the cost of a message as an integer. Here are the costs:

Intro: 2x the message length
Body: 3x the message length
Outro: 4x the message length
Use the built-in len() function to get the length of a string:

helloLen := len("hello")
// helloLen = 5
*/

func printReports(intro, body, outro string) {
	printCostReport(func(intro string) int {
		return 2 * len(intro)
	}, intro)

	printCostReport(func(body string) int {
		return 3 * len(body)
	}, body)

	printCostReport(func(outro string) int {
		return 4 * len(outro)
	}, outro)
}

// don't touch below this line

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
