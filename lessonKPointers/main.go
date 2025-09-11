package main

import "fmt"

/*
Assignment
Fix the bug in the getMessageText function. It's supposed to return a nicely formatted message to the console containing an SMS's recipient and message body. However, it's not working as expected. Run the code and see what happens, then fix the bug.
*/

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}

func test(m Message) {
	defer fmt.Println("=====================================")
	fmt.Printf("the recipient: %s\nthe message: %s\n", m.Recipient, m.Text)
	formatted := getMessageText(m)

	fmt.Printf("\nIntended format: %s\n", formatted)
}

func main() {
	test(Message{
		Recipient: "Holy Spirit",
		Text:      "Please my Lord, lead me throgh this month of September 2025",
	})

	test(Message{
		Recipient: "Jesus",
		Text:      "Please my Lord, smile upon me this month of September 2025",
	})

	test(Message{
		Recipient: "God the Father",
		Text:      "Please my Lord, change my story this month of September 2025",
	})
}
