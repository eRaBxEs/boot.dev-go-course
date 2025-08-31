package main

import (
	"errors"
	"fmt"
)

/*
User Input
In Textio, users can set their profile status to communicate their current activity to those that choose to read it... However, there are some restrictions on what these statuses can contain. Your task is to implement a function that validates a user's status update. The status update cannot be empty and must not exceed 140 characters.

Assignment
Complete the validateStatus function. It should return an error when the status update violates any of the rules:

If the status is empty, return an error that reads status cannot be empty.
If the status exceeds 140 characters, return an error that says status exceeds 140 characters.
*/

func validateStatus(status string) error {

	if len(status) == 0 {
		return fmt.Errorf("status cannot be empty.")
	}
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters.")
	}
	return nil
}

func test(status string) {
	defer fmt.Println("=================================================")
	err := validateStatus(status)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	test("")
	test("This is not empty like the first test")
	longerStatus := `he quick brown fox jumped over the lazy dog, but the dog didn't care. It was busy contemplating the true nature of the universe and whether a belly rub was a form of cosmic alignment. This philosophical musing was interrupted by a squirrel, who chattered incessantly about acorns and the fleeting nature of time, a concept the fox, in its pursuit of fleeting brownness, never truly grasped. The squirrel's monologue was a testament to the chaotic and unpredictable journey of existence, filled with more twists and turns than a labyrinth designed by a bored deity. The dog, ever the stoic observer, simply wagged its tail and dreamt of biscuits.`
	test(longerStatus)
}
