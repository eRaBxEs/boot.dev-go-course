package main

import (
	"fmt"
)

/*
Message Formatter
As Textio evolves, the team has decided to introduce a new feature for custom message formats. Depending on the user's preferences, messages can be sent in different formats, such as plain text, markdown, code, or even encrypted text. To efficiently manage this, you'll implement a system using interfaces.

Assignment
Implement the formatter interface with a method format that returns a formatted string.
Define structs that satisfy the formatter interface: plainText, bold, code.
The structs must all have a message field of type string
plainText should return the message as is.
bold should wrap the message in two asterisks (**) to simulate bold text (e.g., **message**).
code should wrap the message in a single backtick (`) to simulate code block (e.g., `message`)
*/

type formatter interface {
	format() string
}

type plainText struct {
	contentText string
}

func (p plainText) format() string {
	return p.contentText
}

type bold struct {
	contentBold string
}

func (b bold) format() string {
	return fmt.Sprintf("**%s**", b.contentBold)
}

type code struct {
	contentCode string
}

func (c code) format() string {
	return fmt.Sprintf("`%s`", c.contentCode)
}

type invalid struct {
}

func (i invalid) format() string {
	return ""
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}

func test(f formatter) {
	switch v := f.(type) {
	case plainText:
		fmt.Printf("display as plaintext:%s\n", v.format())
		fmt.Println("===============================================================")
	case bold:
		fmt.Printf("display as bold :%s\n", v.format())
		fmt.Println("===============================================================")
	case code:
		fmt.Printf("display as code: %s\n", v.format())
		fmt.Println("===============================================================")
	default:
		fmt.Println("invalid")
		fmt.Println("===============================================================")
	}
}

func main() {
	test(plainText{
		contentText: "welcome to the world of interface type",
	})

	test(bold{
		contentBold: "The earth is the Lord's and the fullness thereof",
	})

	test(code{
		contentCode: `{fmt.Println("Hello World in code")}`,
	})

	test(invalid{})
}
