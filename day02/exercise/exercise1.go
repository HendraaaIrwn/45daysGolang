package main

import "fmt"

func displayMatrix(matrix [5][5]string) {
	for row := range matrix {
		for col := range matrix[row] {
			fmt.Printf("%3s", matrix[row][col])
		}
		fmt.Println()
	}
}

func matrixHollow(matrix [5][5]string) [5][5]string {
	for row := range matrix {
		for col := range matrix {
			if row == 0 || row == len(matrix)-1 || col == 0 || col == len(matrix)-1 {
				matrix[row][col] = "*"

			}
		}
	}
	return matrix
}

func matrixTriangleLeftAligned(matrix [5][5]string) [5][5]string {
	for row := range matrix {
		for col := range matrix {
			if col <= row {
				matrix[row][col] = "*"
			}
		}
	}
	return matrix
}


func main() {
	var matrix [5][5]string
	displayMatrix(matrixHollow(matrix))
	fmt.Println("-----")
	displayMatrix(matrixTriangleLeftAligned(matrix))
}
