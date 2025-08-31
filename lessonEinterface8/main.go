package main

import (
	"fmt"
)

/*
Assignment
We offer a product that allows businesses to send pairs of messages to couples. It is mostly used by flower shops and movie theaters.

Complete the sendSMSToCouple function. It should send 2 messages, first to the customer and then to the customer's spouse.

	Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0 and the error.
	Do the same for the msgToSpouse
	If both messages are sent successfully, return the total cost of the messages added together.
*/

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	// ?
	msgCustomerCost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}

	msgSpouseCost, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}

	return msgCustomerCost + msgSpouseCost, nil

}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func test(msgCustomer, msgToSpouse string) {
	defer fmt.Println("=====================================")
	fmt.Println("message for customer:", msgCustomer)
	fmt.Println("message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total Cost:%v", totalCost)
}

func main() {
	test("Thanks for joining us", "Have a good day")
	test("Thank you", "Enjoy!")
	test("We loved having you in", "We hope the rest of your evening is absolutely fantastic!")
}
