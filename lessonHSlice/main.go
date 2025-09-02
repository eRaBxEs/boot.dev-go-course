package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
Assignment
Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input as well as an array of 3 messages. You've been provided with constants representing the plan types at the top of the file.

If the plan is a pro plan, return all the strings from the messages input in a slice.
If the plan is a free plan, return the first 2 strings from the messages input in a slice.
If the plan isn't either of those, return a nil slice and an error that says unsupported plan.
*/

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if strings.EqualFold(plan, planPro) {
		return messages[:], nil
	}
	if strings.EqualFold(plan, planFree) {
		return messages[:2], nil
	}

	return []string{}, errors.New("unsupported plan")
}

func test(plan string, messages [3]string) {
	defer fmt.Println("==============================================")
	fmt.Printf("sending for plan : %v\n", plan)
	retV, err := getMessageWithRetriesForPlan(plan, messages)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	for i := 0; i < len(retV); i++ {
		fmt.Printf("sending :%v", messages[i])
		fmt.Println()
	}

}

func main() {
	test(planFree, [3]string{"Jesus is Lord", "Do you agree", "You had better"})
	test(planPro, [3]string{"Jesus is Lord", "Do you agree", "You had better"})
}
