package main

import "fmt"

/*
We need to be able to quickly detect bad words in the messages our system sends.

Complete the indexOfFirstBadWord function. If it finds any bad words in the message it should return the index of the first bad word in the msg slice. This will help us filter out naughty words from our messaging system. If no bad words are found, return -1 instead.

Use the range keyword.
*/

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, v := range msg {
		for _, val := range badWords {
			if v == val {
				return i
			}
		}

	}
	return -1
}

func test(msg []string, badWords []string) {
	fmt.Println("=========================================================")
	fmt.Printf("The list of words that we intend to check for bad words includes: %#v\n", msg)
	fmt.Printf("The list of bad words we are checking for includes: %#v\n", badWords)
	index := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("The index of the first bad word is: %v\n", index)

}

func main() {
	test([]string{"Jesus", "Great", "f**ck", "God", "wonder"}, []string{"f**ck", "shit", "poop", "pile"})
	test([]string{"Peace", "Joy", "Amazing", "Gentleness", "sh**t"}, []string{"f**ck", "shit", "poop", "pile"})
	test([]string{"Peace", "Joy", "Complexity", "Gentleness", "shit"}, []string{"f**ck", "shit", "poop", "pile"})
	test([]string{"Happy", "Cry", "Deck", "Glory", "Fullness", "sex", "male"}, []string{"f**ck", "shit", "sex", "poop", "pile"})
}
