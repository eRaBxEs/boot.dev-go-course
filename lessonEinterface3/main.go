package main

import (
	"fmt"
)

/*

Assignment
Complete the required methods so that the email type implements both the expense and formatter interfaces.

Complete the cost() method:

	If the email is not "subscribed", then the cost is 5 cents for each character in the body.
	If it is, then the cost is 2 cents per character.
	Return the total cost of the entire email in cents.
Complete the format() method.

	It should return a string in this format:
	'CONTENT' | Subscribed

	If the email is not subscribed, change the second part to "Not Subscribed":
	'CONTENT' | Not Subscribed

The single quotes are included in the string, and CONTENT is the email's body. For example:

	'Hello, World!' | Subscribed
*/

func (e email) cost() int {
	if !e.isSubscribed {
		return 5 * len(e.body)
	}
	return 2 * len(e.body)
}

func (e email) format() string {
	if !e.isSubscribed {
		return fmt.Sprintf("'%s' | Not Subscribed", e.body)
	}
	return fmt.Sprintf("'%s' | Subscribed", e.body)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

func test(e expense, f formatter) {
	fmt.Println("cost in cents:", e.cost())
	fmt.Println("format is ", f.format())
	fmt.Println("============================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "Hell there",
	}
	test(e, e)

	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
}
