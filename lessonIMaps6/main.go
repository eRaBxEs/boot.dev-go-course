package main

import (
	"fmt"
	"strings"
)

/*

Distinct Words
Complete the countDistinctWords function using a map. It should take a slice of strings and return the total count of distinct words across all the strings. Assume words are separated by spaces. Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).

For example:

messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)

count should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.

Tips
Go's strings package can be very helpful here. Specifically the Fields method and the ToLower method.

Since all that matters is counting distinct words, we don't care about the value of each key in the map. You can use struct{}{} as the value of your map key. This empty struct uses no memory. For example, both map[string]bool and map[string]struct{} can track unique words, but they use different amounts of memory:

distinctWordsBool := make(map[string]bool)
distinctWordsStruct := make(map[string]struct{})

distinctWordsBool["hello"] = true         // Uses 1 byte for the bool value
distinctWordsStruct["hello"] = struct{}{} // Uses 0 bytes for the empty struct
*/

func countDistinctWords(messages []string) int {

	distinctCounts := map[string]struct{}{}
	// distinctCounts := make(map[string]struct{})
	for _, sentence := range messages {
		lowerSentence := strings.ToLower(sentence)
		// assume words are separated by spaces
		words := strings.Fields(lowerSentence)
		for _, w := range words {
			distinctCounts[w] = struct{}{}
		}
	}

	return len(distinctCounts)
}

func test(messages []string) {
	defer fmt.Println("===============================================")
	countDist := countDistinctWords(messages)
	fmt.Printf("The count for distinct words is %v\n", countDist)
}

func main() {
	messages := []string{"Hello world", "hello there", "General Kenobi"}
	test(messages)
	messages = []string{"Hello", "hello there", "General", "great general", "skywalker", "General Kenobi", "Hello world"}
	test(messages)
}
