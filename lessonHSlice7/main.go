package main

import (
	"fmt"
	"strings"
)

/*
Message Filter
Textio is introducing a feature that allows users to filter their messages based on specific criteria. For this feature, messages are categorized into three types: TextMessage, MediaMessage, and LinkMessage. Users can filter their messages to view only the types they are interested in.

Assignment
Your task is to implement a function that filters a slice of messages based on the message type.

Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.
*/

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {

	finalMessage := []Message{}

	for _, message := range messages {
		switch messageType := message.(type) {
		case TextMessage:
			if strings.EqualFold(filterType, message.Type()) {
				finalMessage = append(finalMessage, messageType)
			}
		case MediaMessage:
			if strings.EqualFold(filterType, message.Type()) {
				finalMessage = append(finalMessage, messageType)
			}
		case LinkMessage:
			if strings.EqualFold(filterType, message.Type()) {
				finalMessage = append(finalMessage, messageType)
			}
		}
	}

	return finalMessage

}

func test(messages []Message, filterType string) {
	defer fmt.Println("============= END REPORT==================")
	fmt.Printf("We have a total of %v messages\n", len(messages))
	fmt.Printf("filtering the messages by %s\n", filterType)
	filteredMessages := filterMessages(messages, filterType)
	fmt.Printf("We now have a total of %v %s messages\n", len(filteredMessages), filterType)

}

func main() {
	test(
		[]Message{
			TextMessage{
				Sender:  "Segun",
				Content: "God is your helper in these times",
			},
			TextMessage{
				Sender:  "Ehioje",
				Content: "You are God sent, He has not forgotten you, keep praying",
			},
			TextMessage{
				Sender:  "September",
				Content: "The month of bringing forth your testimonies",
			},
			TextMessage{
				Sender:  "Pastor",
				Content: "It is joy time, rejoice evermopre",
			},
			MediaMessage{
				Sender:    "Olusola",
				MediaType: "images",
				Content:   "graduation.jpg",
			},
			MediaMessage{
				Sender:    "Victoria",
				MediaType: "images",
				Content:   "promotion.jpg",
			},
			MediaMessage{
				Sender:    "Crystal",
				MediaType: "images",
				Content:   "celebration.jpg",
			},
			LinkMessage{
				Sender:  "Google",
				URL:     "https://google.com",
				Content: "Contains the most powerful search engine ever",
			},
			LinkMessage{
				Sender:  "ChatGPT",
				URL:     "https://chatgpt.com",
				Content: "Contains the world leading AI",
			},
			LinkMessage{
				Sender:  "Github",
				URL:     "https://github.com",
				Content: "Contains the world's biggest online repository platform",
			},
			LinkMessage{
				Sender:  "LinkedIn",
				URL:     "https://linkedin.com",
				Content: "Contains the world's biggest job platform platform",
			},
			LinkMessage{
				Sender:  "ubuntu",
				URL:     "https://ubuntu.com",
				Content: "Contains the world's biggest linux system",
			},
		},
		"media",
	)
}
