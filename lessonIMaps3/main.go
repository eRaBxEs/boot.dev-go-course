package main

import (
	"fmt"
)

/*
Assignment
Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.

Implement the updateCounts function. It takes as input:

messagedUsers: a slice of strings.
validUsers: a map of string -> int.
It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.

So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.
*/

func updateCounts(messagedUsers []string, validUsers map[string]int) map[string]int {

	for _, messUser := range messagedUsers {
		if _, ok := validUsers[messUser]; ok {
			validUsers[messUser]++
		}
	}

	return validUsers
}

func test(messagedUsers []string, validUsers map[string]int) {
	defer fmt.Println("=================================================")
	fmt.Printf("The name %s in the slice ", messagedUsers[0])
	updatedUsers := updateCounts(messagedUsers, validUsers)
	fmt.Printf("has %v counts\n", updatedUsers[messagedUsers[0]])
	fmt.Println("Looping through the keys and their values")
	for k, v := range updatedUsers {
		fmt.Printf("The name %s in the slice has a count of %v \n", k, v)
	}

	fmt.Printf("Having a view of the new map data structure: %#v\n", updatedUsers)

}

func main() {
	messagedUsers := []string{"James", "John", "Jack", "Jake", "Jesus", "Jake", "Jesus", "James", "John", "Jesus", "James", "Jesus", "James", "Jake", "Jesus", "James"}
	validUsers := map[string]int{
		"Jesus": 3,
		"James": 1,
		"Jake":  0,
		"Jack":  2,
		"John":  1,
	}
	test(messagedUsers, validUsers)

}
