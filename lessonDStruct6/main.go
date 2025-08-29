package main

import (
	"fmt"
)

/*
We need a way to differentiate between standard and premium users. When a new user is created, they need a membership type, and that type will determine the message character limit.

Assignment
Create a new struct called Membership, it should have:
	A Type string field
	A MessageCharLimit integer field
Update the User struct to embed a Membership.
Complete the newUser function. It should return a new User with all the fields set as you would expect based on the inputs. If the user is a "premium" member, the MessageCharLimit should be 1000, otherwise, it should only be 100.
*/

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	newUser := User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
		},
	}

	if membershipType == "standard" {
		newUser.MessageCharLimit = 100
	}
	if membershipType == "premium" {
		newUser.MessageCharLimit = 1000
	}

	return newUser
}

func main() {
	fmt.Println(newUser("Ehioje", "premium"))
	fmt.Println("=============================================")

	fmt.Println(newUser("Henry", "standard"))
	fmt.Println("=============================================")
}
