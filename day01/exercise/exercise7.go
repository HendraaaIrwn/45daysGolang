package main

import "fmt"

func countDigit(number int) {
	countNumber := 0
	switch {
	case number/10 != 0:
		for number > 0 {
			number = number / 10
			countNumber++
		}
	case number/10 == 0:
		countNumber = 1
	}
	fmt.Printf("%d", countNumber)
}

func countDigit1(number int) int {
	counter := 0
	for {
		number = number / 10
		counter++
		if number == 0 {
			break
		}
	}
	return counter
}

func main() {
	// countDigit1(34567890)
	fmt.Printf("%d", countDigit1(3456789000))
}
