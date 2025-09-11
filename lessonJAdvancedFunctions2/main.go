package main

import "fmt"

/*
Assignment
Complete the reformat function. It takes a message string and a formatter function as input:

Apply the given formatter three times to the message
Add a prefix of TEXTIO: to the result
Return the final string
For example, if the message is "General Kenobi" and the formatter adds a period to the end of the string, the final result should be

TEXTIO: General Kenobi...
*/

func reformat(message string, formatter func(string) string) string {
	// Apply the formatter function three times
	formattedMessage := formatter(message)
	formattedMessage = formatter(formattedMessage)
	formattedMessage = formatter(formattedMessage)

	// Add the prefix and return the final string
	return fmt.Sprintf("TEXTIO: %s", formattedMessage)
}

// Example formatter function (as described in the prompt)
func addPeriod(s string) string {
	return s + "."
}

func main() {
	message := "General Kenobi"

	// Pass the message and the formatter function to reformat
	finalString := reformat(message, addPeriod)

	fmt.Println(finalString)
}
