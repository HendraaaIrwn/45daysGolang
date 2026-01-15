package main

import "fmt"

func printBoxNumber(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

		}
		println()
	}

}

func printFizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(" fizz buzz ")
		} else if i%3 == 0 {
			fmt.Print(" fizz ")
		} else if i%5 == 0 {
			fmt.Print(" buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}

func countFizzBuzz(n int) (int, int, int) {
	countFizz, countBuzz, countFizzBuzz := 0, 0, 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			countFizz++
		}

		if i%5 == 0 {
			countBuzz++
		}

		if i%3 == 0 && i%5 == 0 {
			countFizzBuzz++
		}
	}
	return countFizz, countBuzz, countFizzBuzz
}

func main() {

}
