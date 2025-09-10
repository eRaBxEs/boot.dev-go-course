package main

import (
	"fmt"
	"strings"
)

/*
Message Tagger
Textio needs a way to tag messages based on specific criteria.

Assignment
Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes a sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.
It should loop through each message and set the tags to the result of the passed in function.
Be sure to modify the messages of the original slice using bracket notation messages[i].
Ensure the tags field contains an initialized slice. No nil slices.
Complete the tagger function. It should take an sms message and return a slice of strings.
For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.
Regardless of casing just means that even "uRGent" or "SALE" should trigger the tag.

Example usage:

messages := []sms{
	{id: "001", content: "Urgent! Last chance to see!"},
	{id: "002", content: "Big sale on all items!"},
	// Additional messages...
}
taggedMessages := tagMessages(messages, tagger)
// `taggedMessages` will now have tags based on the content.
// 001 = [Urgent]
// 002 = [Promo]

Tip
The go strings package, specifically the Contains and ToLower functions, can be very helpful here!
*/

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, msg := range messages {
		messages[i].tags = tagger(msg)
	}
	return messages
}

func tagger(msg sms) []string {

	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "[Urgent]")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "[Promo]")
	}

	return tags
}

func test(message []sms, tagger func(sms) []string) {
	defer fmt.Println("================================================")
	fmt.Println("A brief view of some of the messages:")
	fmt.Println("=======================   BEFORE ADJUSTMENT  =========================")

	for i, _ := range message {
		fmt.Printf("The [%v] message before adjustment %#v\n", i, message[i])
	}
	newMessages := tagMessages(message, tagger)
	fmt.Println("=======================   AFTER ADJUSTMENT  =========================")

	for i, _ := range newMessages {
		fmt.Printf("The [%v] message before adjustment %#v\n", i, newMessages[i])
	}

}

func main() {
	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see!", tags: []string{}},
		{id: "002", content: "Big sale on all items!", tags: []string{}},
		{id: "003", content: "General urGenT is a big sale", tags: []string{}},
		{id: "004", content: "Great SaLE is what we consider", tags: []string{}},
		{id: "005", content: "Return SAle if URgenT", tags: []string{}},
	}

	test(messages, tagger)

}
