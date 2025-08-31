package main

import (
	"fmt"
)

/*
Our users frequently try to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It's become such a problem that we need to make a new type of error for division by zero.

Update the code so that the divideError type implements the error interface. Its Error() method should just return a string formatted in the following way:

can not divide DIVIDEND by zero

Where DIVIDEND is the actual dividend of the divideError. Use the %v verb to format the dividend as a float.
*/

type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero\n", d.dividend)
}

// ?

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	defer fmt.Println("=======================================")
	fmt.Printf("Dividing %.2f by %.2f\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Printf("Error: %v\n", divideError{dividend: dividend})
		return
	}
	fmt.Printf("QUOTIENT:%.2f\n", quotient)
}

func main() {
	test(10.0, 2.0)
	test(6.0, 3.0)
	test(8.0, 2.0)
	test(5.0, 0.0)
	test(2.0, 0.0)
	test(1.0, 0.0)
	test(0.0, 0.0)
}
