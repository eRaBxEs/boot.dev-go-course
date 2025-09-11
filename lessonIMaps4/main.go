package main

import (
	"fmt"
)

/*
Assignment
Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

Complete the getNameCounts function. It takes a slice of strings names and returns a nested map. The parent map's keys are all the unique first characters (see runes) of the names, the nested maps keys are all the names themselves, and the value is the count of each name.

For example:

billy
billy
bob
joe

Creates the following nested map:

b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}

Tips
The test suite is not printing the map you're returning directly, but instead checking a few specific keys.
You can convert a string into a slice of runes as follows:
runes := []rune(word)
*/

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}

		firstChar := name[0]
		_, ok := counts[rune(firstChar)]
		if !ok {
			counts[rune(firstChar)] = make(map[string]int)
		}
		counts[rune(firstChar)][name]++
	}
	return counts

}

func test(names []string) {
	fmt.Printf("Generating counts for %v names ... \n", len(names))

	nameCounts := getNameCounts(names)
	fmt.Printf("data structute for returned value: %#v\n", nameCounts)
	fmt.Println("To further interpret the rune which is showing as numbers, will loop through to show a better data structure:")
	for v, k := range nameCounts {
		fmt.Printf("%c: %#v\n", v, k)
	}
	fmt.Println("=======================================================================")

}

func main() {
	test([]string{"billy", "billy", "bob", "joe"})
	test([]string{"jake", "jane", "bob", "billy", "jesus", "peter", "paul"})
	test([]string{"jake", "jane", "bob", "billy", "jesus", "peter", "paul", "jesus", "jane", "jesus"})
}
