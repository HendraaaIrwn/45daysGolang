package main

import (
	"fmt"
)

func displayMatrix(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%3d", matrix[row][col])
		}
		fmt.Println()
	}
}

func displayMatrix1(matrix [5][5]int) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%3d", value)
		}
		fmt.Println()
	}
}

func displayMatrix2(matrix [5][5]int) {
	for row := range matrix {
		for col := range matrix[row] {
			fmt.Printf("%3d", matrix[row][col])
		}
		fmt.Println()
	}
}

func displayMatrixString(matrix [5][5]string) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%3s", value)
		}
		fmt.Println()
	}
}

func initMatrix() [5][5]int {
	var matrix [5][5]int
	counter := 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = counter
			counter++
		}
	}
	return matrix
}

func initMatrix2() [5][5]int {
	var matrix [5][5]int
	counter := 1
	for row := range matrix {
		for col := range matrix {
			matrix[row][col] = counter
			counter++
		}
	}
	return matrix
}

func main() {
	matrix := initMatrix2()
	displayMatrix(matrix)
	fmt.Println("-----")
	displayMatrix1(matrix)
	fmt.Println("-----")
	displayMatrix2(matrix)
}
