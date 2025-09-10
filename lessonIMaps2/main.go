package main

import (
	"errors"
	"fmt"
)

/*
Assignment
It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function. The user struct has a scheduledForDeletion field that determines if they are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original — we don't have a copy.
*/

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {

	elem, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}

	if !elem.scheduledForDeletion {
		return false, nil
	}

	delete(users, name)
	return true, nil

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func test(users map[string]user, name string) {
	defer fmt.Println("===================================================")
	fmt.Printf("checking if name %s exist\n", name)
	isDeleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Printf("Is name scheduled for deletion ?: %v\n", isDeleted)
}

func main() {
	users := map[string]user{
		"James": user{
			name:                 "James",
			number:               12222,
			scheduledForDeletion: false,
		},
		"Jack": user{
			name:                 "Jack",
			number:               222233,
			scheduledForDeletion: true,
		},
		"John": user{
			name:                 "John",
			number:               378888,
			scheduledForDeletion: false,
		},
		"Jared": user{
			name:                 "Jared",
			number:               100777,
			scheduledForDeletion: true,
		},
		"Jesus": user{
			name:                 "Jesus",
			number:               777777,
			scheduledForDeletion: false,
		},
	}

	test(users, "Junior")
	test(users, "James")
	test(users, "Jared")
}
