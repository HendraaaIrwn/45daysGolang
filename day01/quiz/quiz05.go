package main

import "fmt"

func printZigzagGrid(n int) {
	cols := n + 1 

	for i := 1; i <= n; i++ {
		left := i
		right := n - i + 1

		for col := 1; col <= cols; col++ {
			if col%2 == 1 {
				fmt.Printf("%d ", left)
			} else {
				fmt.Printf("%d ", right)
			}
		}
		fmt.Println()
	}
}


func main() {
	printZigzagGrid(9)
	printZigzagGrid(5)
}