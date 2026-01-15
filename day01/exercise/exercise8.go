package main

import "fmt"

func isPrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrimeNumber(11))
	fmt.Println(isPrimeNumber(8))

}
