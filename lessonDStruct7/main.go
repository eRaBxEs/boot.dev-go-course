package main

import (
	"fmt"
)

/*
Send Message
Assignment
Create a SendMessage method for the User struct.

It should take a message string and messageLength int as inputs.

If the messageLength is within the user's message character limit, return the original message and true (indicating the message can be sent), otherwise, return an empty string and false.
*/

// ?
func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	}
	return "", false
}

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}

func main() {
	user1 := newUser("Ehioje", "premium")
	fmt.Println(user1)
	fmt.Println(user1.SendMessage("Jesus is alive and He is coming very soon", 30))
	fmt.Println("=============================================")

	user2 := newUser("Henry", "standard")
	fmt.Println(user2)
	fmt.Println(user2.SendMessage("He said I am the way, the truth and the life. He is the Son of God, whether you agree or not, it changes nothing", 120))
	fmt.Println("=============================================")
}
