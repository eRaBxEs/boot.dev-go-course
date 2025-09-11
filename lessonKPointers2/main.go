package main

import (
	"fmt"
	"strings"
)

/*
Assignment
Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

Word	Replacement
fubb	****
shiz	****
witch	*****
It should update the value in the pointer and return nothing.

Do not alter the function signature.
*/

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal // passing the value back into the pointer
}

func test(messages []string) {
	for i := range messages { // accessing via index enables one to avoid error by copying value
		removeProfanity(&messages[i])
		fmt.Println(messages[i])
	}

	fmt.Println(messages)
}

func main() {
	test([]string{"You are not a fubb", "This is all shiz", "She is an old witch"})
}
