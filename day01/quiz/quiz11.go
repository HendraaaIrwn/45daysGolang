package main

import "fmt"

func isNumberPalindrome(n int64) {
	if n < 0 {
		fmt.Println(false)
		return
	}

	original := n
	var reversed int64 = 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	fmt.Println(reversed == original)
}

func main() {
	isNumberPalindrome(121)		// true
	isNumberPalindrome(2147447412) // true
	isNumberPalindrome(333) 	// true
	isNumberPalindrome(110)		// false
	isNumberPalindrome(11)		// true
}
