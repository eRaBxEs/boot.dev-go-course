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
	if message == nil { // The correct check is if the pointer itself is nil
		fmt.Println("message is nil  :: no bad words to remove")
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
	// This will now work
	test([]string{"", "You are not a fubb", "This is all shiz", "She is an old witch"})

	// This will trigger the nil pointer check
	var message *string
	removeProfanity(message)

	// A single word test
	word := "a little fubb"
	removeProfanity(&word)
	fmt.Println(word)
}
