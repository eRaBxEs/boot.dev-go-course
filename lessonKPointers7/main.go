package main

import (
	"errors"
	"fmt"
)

/*
Assignment
Implement the updateBalance function. It should take a customer pointer and a transaction, and return an error. Depending on the transactionType, it should either add or subtract the amount from the customer's balance. If the customer does not have enough money, it should return the error insufficient funds. If the transactionType isn't a withdrawal or deposit, it should return the error unknown transaction type. Otherwise, it should process the transaction and return nil.

alice := customer{id: 1, balance: 100.0}
deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

updateBalance(&alice, deposit)
// id: 1 balance: 150
*/

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?

func updateBalance(customer *customer, transaction transaction) error {

	if transaction.transactionType == transactionDeposit {
		customer.balance += transaction.amount

	} else if transaction.transactionType == transactionWithdrawal {

		if transaction.amount > customer.balance {
			return errors.New("insufficient funds.")
		}
		customer.balance -= transaction.amount

	} else {
		return errors.New("unknown transaction type.")
	}

	return nil
}

func test(customer *customer, transaction transaction) {
	defer fmt.Println("============ END OF REPORT ==============")
	fmt.Printf("Getting the balance for customer id: %d who had a %s transaction of amount:: %.2f\n", customer.id, transaction.transactionType, transaction.amount)
	err := updateBalance(customer, transaction)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Customer has a balance of %.2f\n", customer.balance)
}

func main() {
	// Testing my solution
	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

	updateBalance(&alice, deposit)

	fmt.Printf("Alice has a balance of %.2f\n", alice.balance)
	fmt.Println("===== using my test cases =====")
	test(&customer{id: 2, balance: 200.0}, transaction{customerID: 2, amount: 70, transactionType: transactionDeposit})
	test(&customer{id: 4, balance: 300.0}, transaction{customerID: 4, amount: 180, transactionType: transactionWithdrawal})
	test(&customer{id: 7, balance: 100.0}, transaction{customerID: 7, amount: 100, transactionType: transactionWithdrawal})
	test(&customer{id: 8, balance: 400.0}, transaction{customerID: 8, amount: 450, transactionType: transactionWithdrawal})

}
