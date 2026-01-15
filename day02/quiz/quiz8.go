package main

import "fmt"

func displayMatrix(matrix [7][7]string) {
	for row := range matrix {
		for col := range matrix {
			fmt.Printf("%3s", matrix[row][col])
		}
		fmt.Println()
	}
}

func matrixNumberHollow(matrix [7][7]string) [7][7]string {
	for row := range matrix {
		for col := range matrix {
			if row == 0 {
				matrix[row][col] = fmt.Sprintf("%d", col)
			} else if row == len(matrix)-1 {
				matrix[row][col] = fmt.Sprintf("%d", row+col)
			} else if col == 0 {
				matrix[row][col] = fmt.Sprintf("%d", row)
			} else if col == len(matrix)-1 {
				matrix[row][col] = fmt.Sprintf("%d", row+col)
			} else {
				matrix[row][col] = " "
			}
		}
	}
	return matrix
}

func main() {
	var matrix [7][7]string
	displayMatrix(matrixNumberHollow(matrix))
}
