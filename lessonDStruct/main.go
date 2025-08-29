package main

import "fmt"

/*
Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present, return false instead.
*/

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}

	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

// Main function to test the solution
func main() {
	// A valid message
	validMessage := messageToSend{
		message: "Hello there!",
		sender: user{
			name:   "Obi-Wan Kenobi",
			number: 12345,
		},
		recipient: user{
			name:   "General Grievous",
			number: 67890,
		},
	}
	fmt.Printf("Is the valid message sendable? %v\n", canSendMessage(validMessage)) // Should return true

	// An invalid message with a missing sender number
	invalidSenderNumber := messageToSend{
		message: "Hello there!",
		sender: user{
			name:   "Obi-Wan Kenobi",
			number: 0, // This is a default zero value
		},
		recipient: user{
			name:   "General Grievous",
			number: 67890,
		},
	}
	fmt.Printf("Is the message with a missing sender number sendable? %v\n", canSendMessage(invalidSenderNumber)) // Should return false

	// An invalid message with a missing recipient name
	invalidRecipientName := messageToSend{
		message: "Hello there!",
		sender: user{
			name:   "Obi-Wan Kenobi",
			number: 12345,
		},
		recipient: user{
			name:   "", // This is a default zero value
			number: 67890,
		},
	}
	fmt.Printf("Is the message with a missing recipient name sendable? %v\n", canSendMessage(invalidRecipientName)) // Should return false
}
