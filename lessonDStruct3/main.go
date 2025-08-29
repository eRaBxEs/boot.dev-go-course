package main

import (
	"fmt"
)

/*
embedded structs provide a kind of data-only inheritance that can be useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, but embedded structs are a way to elevate and share fields between struct definitions.
*/

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender name:", s.rateLimit)
	fmt.Println("==============================================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 127664544,
		},
	})

	test(sender{
		rateLimit: 15000,
		user: user{
			name:   "Segun",
			number: 34986464,
		},
	})

	test(sender{
		rateLimit: 20000,
		user: user{
			name:   "Henry",
			number: 73442552,
		},
	})
}
