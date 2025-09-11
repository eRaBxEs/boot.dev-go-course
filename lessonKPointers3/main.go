package main

import (
	"fmt"
)

/*
Assignment
Write a analyzeMessage function. It should accept a pointer to an Analytics struct and a Message struct (not a pointer).

It should look at the Success field of the Message struct and, based on that, increment the MessagesTotal, MessagesSucceeded, or MessagesFailed fields of the Analytics struct as appropriate.
*/

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

// ?
func analyzeMessage(analytics *Analytics, message Message) {
	if message.Success {
		(*analytics).MessagesSucceeded++
		(*analytics).MessagesTotal++
	} else {
		(*analytics).MessagesFailed++
		(*analytics).MessagesTotal++
	}

	fmt.Println("message success::", analytics.MessagesSucceeded)
	fmt.Println("message failed::", analytics.MessagesFailed)
	fmt.Println("message total::", analytics.MessagesTotal)
}

func main() {
	analytics := &Analytics{
		MessagesTotal:     2,
		MessagesSucceeded: 1,
		MessagesFailed:    1,
	}

	message := Message{
		Recipient: "Segun",
		Success:   true,
	}

	analyzeMessage(analytics, message)
}
