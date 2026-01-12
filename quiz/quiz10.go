package main

import "fmt"

func checkBraces(s string) {
	BracesBalance := 0

	for _, ch := range s {
		if ch == '(' {
			BracesBalance++
		} else if ch == ')' {
			BracesBalance--
			// jika tutup duluan → tidak valid
			if BracesBalance < 0 {
				fmt.Println(false)
				return
			}
		}
	}

	fmt.Println(BracesBalance == 0)
}

func main() {
	checkBraces("()()") // true
	checkBraces("()()") // true
	checkBraces("((()") // false
	checkBraces("(()))((())") // false
}
