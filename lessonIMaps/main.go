package main

import (
	"errors"
	"fmt"
)

/*
Assignment
We can speed up our contact-info lookups by using a map!

Key-based map lookup: O(1)
Slice brute-force search: O(n)
Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and an error. A user struct just contains a user's name and phone number. The first element in the names slice pairs with the first phone number, and so on.

If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".
*/

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	mapUsers := make(map[string]user)
	for i := 0; i < len(names); i++ {
		mapUsers[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}

	return mapUsers, nil
}

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	defer fmt.Println("=====================================================")
	fmt.Printf("The first name on the list: %v\n", names[0])
	fmt.Printf("The first address on the list: %v\n", names[0])
	createdMap, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range createdMap {
		fmt.Printf("The name %s has the user detail %#v\n", k, v)
	}
}

func main() {
	test([]string{"James", "John", "Jack"}, []int{12222, 344445, 3452353, 777777})
	test([]string{"James", "John", "Jack", "Jesus"}, []int{12222, 344445, 3452353, 777777})
}
