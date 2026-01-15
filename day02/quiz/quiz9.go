package main

import "fmt"

func displayMatrix(matrix [7][7]int) {
	for row := range matrix {
		for col := range matrix {
			fmt.Printf("%3d", matrix[row][col])
		}
		fmt.Println()
	}
}

func matrixFullNumber(matrix [7][7]int) [7][7]int {
	for row := range matrix {
		for col := range matrix {
			if row == 0 {
				matrix[row][col] = col
			} else if row == len(matrix)-1 {
				matrix[row][col] = row + col
			} else if col == 0 {
				matrix[row][col] = row
			} else if col == len(matrix)-1 {
				matrix[row][col] = row + col
			} else {
				matrix[row][col] = row + col
			}
		}
	}
	return matrix
}


func main() {
	var matrix [7][7]int
	displayMatrix(matrixFullNumber(matrix))
}
