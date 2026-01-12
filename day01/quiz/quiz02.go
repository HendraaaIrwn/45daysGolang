package main

import "fmt"

func extractDigit(number int) {
	for number > 0 {
		digit := number % 10
		fmt.Printf("%d ", digit)
		number = number / 10
	}
	fmt.Println()
}

func main() {
	extractDigit(12234)
	extractDigit(5432) 
	extractDigit(1278)
}