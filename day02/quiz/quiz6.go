package main

import "fmt"

func displayMatrix(matrix [5][5]int) {
	for row := range matrix {
		for col := range matrix {
			fmt.Printf("%3d", matrix[row][col])
		}
		fmt.Println()
	}
}

func printHorizontalNumber(matrix [5][5]int) [5][5]int {
	for row := range matrix {
		for col := range matrix {
			if col == row {
				matrix[row][col] = col+1
			} else if col < row {
				matrix[row][col] = 20
			} else {
				matrix[row][col] = 10
			}
		}
	}
	return matrix
}

func main() {
	var matrix [5][5]int
	displayMatrix(printHorizontalNumber(matrix))
}
