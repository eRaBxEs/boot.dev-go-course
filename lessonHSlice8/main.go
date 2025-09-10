package main

import (
	"fmt"
)

/*
Password Strength
As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

At least 5 characters long but no more than 12 characters.
Contains at least one uppercase letter.
Contains at least one digit.
A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in the previous lesson to iterate over the characters of a string.

Assignment
Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.

Assume passwords consist of ASCII characters only.

Tip
Remember that characters in Go strings are really just bytes under the hood. You can compare a character to another character like 'A' or '0' to check if it falls within a certain range.
*/

func isValidPassword(password string) bool {
	// ?
	upperChar := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	digitChar := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	if len(password) < 5 || len(password) > 12 {
		fmt.Println("password should be atleast 5 characters long but not more than 12 characters")
		return false
	}

	// check for upper case
	foundUppercase := false
	for _, char := range password {
		for _, upChar := range upperChar {
			if char == upChar {
				foundUppercase = true
			}
		}
	}

	if !foundUppercase {
		fmt.Println("password should contain at least one uppercase letter")
		return false
	}

	// check for digit
	foundDigit := false
	for _, char := range password {
		for _, digChar := range digitChar {
			if char == digChar {
				foundDigit = true
			}
		}
	}

	if !foundDigit {
		fmt.Println("password should contain at least one digit")
		return false
	}

	return true
}

func test(password string) {
	defer fmt.Println("==========================================================")
	fmt.Printf("The given password is %s\n", password)
	validity := isValidPassword(password)
	fmt.Printf("The validity of the password :: %v\n", validity)
}

func main() {
	test("congratulations")
	test("Congrats")
	test("congrats9")
	test("Congrats9")
}
