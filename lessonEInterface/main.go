package main

import (
	"fmt"
	"time"
)

/*
Assignment
The birthdayMessage and sendingReport structs already have implementations of the getMessage method. The getMessage method returns a string, and any type that implements the method can be considered a message (meaning it implements the message interface).

	Add the getMessage() method signature as a requirement on the message interface.
	Complete the sendMessage function. It should return:
		The content of the message.
		The cost of the message, which is the length of the message multiplied by 3.
Notice that your code doesn't care at all about whether a specific message is a birthdayMessage or a sendingReport!
*/

func sendMessage(msg message) (string, int) {
	return msg.getMessage(), len(msg.getMessage()) * 3
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test(m message) {
	message, cost := sendMessage(m)
	fmt.Println("message:", message, "cost:", cost)
	fmt.Println("====================================================================")
}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})

	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
}
