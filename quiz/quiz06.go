package main

import "fmt"

func printDashNumberGrid(n int) {
	for row := 1; row <= n; row++ {
		for col := 1; col <= n; col++ {
			// Baris ganjil -> tampilkan angka di kolom genap
			// Baris genap  -> tampilkan angka di kolom ganjil
			if (row%2 == 1 && col%2 == 0) || (row%2 == 0 && col%2 == 1) {
				fmt.Printf("%2d ", col)
			} else {
				fmt.Printf("%2s ", "-")
			}
		}
		fmt.Println()
	}
}

func main() {
	printDashNumberGrid(9)
	printDashNumberGrid(5)
}