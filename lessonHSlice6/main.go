package main

import (
	"fmt"
)

/*
Assignment
We support various visualization dashboards on Textio that display message analytics to our users. The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.

Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers. The value of each cell is i * j where i and j are the indexes of the row and column respectively. Basically, we're building a multiplication chart.

For example, a 5x10 matrix, produced from calling createMatrix(5, 10), would look like this:

[0  0  0  0  0  0  0  0  0  0]
[0  1  2  3  4  5  6  7  8  9]
[0  2  4  6  8 10 12 14 16 18]
[0  3  6  9 12 15 18 21 24 27]
[0  4  8 12 16 20 24 28 32 36]

*/

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix ....\n", rows, cols)
	matrix := createMatrix(rows, cols)

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	fmt.Println("===================== END OF REPORT ======================")

}

func main() {
	test(2, 2)
	test(3, 3)
	test(5, 5)
}
