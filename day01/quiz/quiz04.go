package main

import "fmt"

func printNumberPyramid(n int) {
	for i := n; i >= 1; i-- {
		// bagian turun: i ... 1
		for d := i; d >= 1; d-- {
			fmt.Printf("%d ", d)
		}
		// bagian naik: 2 ... i
		for d := 2; d <= i; d++ {
			fmt.Printf("%d ", d)
		}

		fmt.Println()
	}
}


func main() {
	var n int
	fmt.Print("Input jumlah baris piramid: ")
	fmt.Scan(&n)

	printNumberPyramid(n)
}