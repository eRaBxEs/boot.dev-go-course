package main

import (
	"fmt"
	"strings"
)

/*
Assignment
Let's make our profanity checker safe. Update the removeProfanity function. If message is nil, return early to avoid a panic. After all, there are no bad words to remove.
*/

func removeProfanity(message *string) {
	if message == nil {
		fmt.Println("message is nil  :: no bad words to remove")
		return
	}
	if *message == "" {
		fmt.Println("message is empty  :: no bad words to remove")
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal // passing the value back into the pointer
}

func test(messages []string) {
	defer fmt.Println("======================================================")
	for i := range messages {
		removeProfanity(&messages[i])
		fmt.Println(messages[i])
	}

}

func main() {
	// test([]string{})
	test([]string{""})
	test([]string{"You are not a fubb", "This is all shiz", "She is an old witch"})

	// this will trigger the empty string check
	word := ""
	removeProfanity(&word) // In this case the string is empty and not nil

	// creating a nil pointer to a string
	var nilPointer *string
	removeProfanity(nilPointer)
}
