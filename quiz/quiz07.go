package main

import "fmt"

func perfectNumber(limit int) {
	for n := 2; n <= limit; n++ {
		sum := 0

		for d := 1; d <= n/2; d++ {
			if n%d == 0 {
				sum += d
			}
		}

		if sum == n {
			fmt.Printf("%d ", n)
		}
	}
	fmt.Println()
}

func main() {
	perfectNumber(28)
	perfectNumber(1000)
	perfectNumber(10000)
}
