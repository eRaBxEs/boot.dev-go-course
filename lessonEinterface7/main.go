package main

import (
	"fmt"
)

/*
Process Notifications
Textio now has a system to process different types of notifications: direct messages, group messages, and system alerts. Each notification type has a unique way of calculating its importance score based on user interaction and content.

Assignment
a. Implement the importance methods for each message type. They should return the importance score for each message type.
	1. For a directMessage the importance score is based on if the message isUrgent or not. If it is urgent, the importance score is 50 otherwise the importance score is equal to the DM's priorityLevel.
	2. For a groupMessage the importance score is based on the message's priorityLevel
	3. All systemAlert messages should return a 100 importance score.
b. Complete the processNotification function. It should identify the type and return different values for each type
	1. For a directMessage, return the sender's username and importance score.
	2. For a groupMessage, return the group's name and the importance score.
	3. For a systemAlert, return the alert code and the importance score.
	4. If the notification does not match any known type, return an empty string and a score of 0.
*/

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (s systemAlert) importance() int {
	return 100
}

type invalid struct {
}

func (i invalid) importance() int {
	return 0
}

// ?

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, v.importance()
	case groupMessage:
		return v.groupName, v.importance()
	case systemAlert:
		return v.alertCode, v.importance()
	default:
		return "", v.importance()
	}
}

func test(n notification) {
	switch v := n.(type) {
	case directMessage:
		fmt.Printf("DirectMessage sendername:%s and importance score:%v\n", v.senderUsername, v.importance())
		fmt.Println("==================================================================================================")
	case groupMessage:
		fmt.Printf("GroupMessage groupName:%s and importance score:%v\n", v.groupName, v.importance())
		fmt.Println("==================================================================================================")
	case systemAlert:
		fmt.Printf("SystemAlert alertCode:%s and importance score:%v\n", v.alertCode, v.importance())
		fmt.Println("==================================================================================================")
	default:
		fmt.Printf("Invalid:%s and importance score:%v", `""`, v.importance())
		fmt.Println("==================================================================================================")
	}
}

func main() {
	test(directMessage{
		senderUsername: "Jonny",
		messageContent: "hello there",
		priorityLevel:  200,
		isUrgent:       true,
	})

	test(directMessage{
		senderUsername: "Jack",
		messageContent: "whats going on with the application in production",
		priorityLevel:  50,
		isUrgent:       true,
	})

	test(directMessage{
		senderUsername: "Billy",
		messageContent: "hope the tests have all been carried out",
		priorityLevel:  10,
		isUrgent:       false,
	})

	test(groupMessage{
		groupName:      "Admin",
		messageContent: "we are shutting down the application to reboot in 30 mins",
		priorityLevel:  5,
	})

	test(groupMessage{
		groupName:      "User",
		messageContent: "we are testing various versions of the application",
		priorityLevel:  2,
	})

	test(systemAlert{
		alertCode:      "red",
		messageContent: "reboot required",
	})

	test(systemAlert{
		alertCode:      "blue",
		messageContent: "refresh the application",
	})
}
