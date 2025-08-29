package main

import "fmt"

/*
Assignment
Implement the getExpenseReport function.

If the expense is an email, return the email's toAddress and the cost of the email.
If the expense is an sms, return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, return an empty string and 0.0 for the cost.
*/

func getExpenseReport(e expense) (string, float64) {
	mail, ok := e.(email)
	if ok {
		return mail.toAddress, mail.cost()
	}

	mess, ok := e.(sms)
	if ok {
		return mess.toPhoneNumber, mess.cost()
	}

	return "", 0.0

}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email is going to %s; will cost %.2f\n", address, cost)
		fmt.Println("========================================================================")
	case sms:
		fmt.Printf("Report: The sms is going to %s; will cost %.2f\n", address, cost)
		fmt.Println("========================================================================")
	default:
		fmt.Println("Report: Invalid Expense")
		fmt.Println("========================================================================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "Hello there",
		toAddress:    "johndoe@gmail.com",
	})

	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "john@doe.com",
	})

	test(sms{
		isSubscribed:  false,
		body:          "We are expecting you home",
		toPhoneNumber: "+13666377",
	})

	test(sms{
		isSubscribed:  true,
		body:          "welcome to the factory",
		toPhoneNumber: "+444667748",
	})
}
