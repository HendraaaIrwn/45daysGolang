package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(text string) {
	// normalisasi: lowercase & ambil huruf/angka saja
	clean := ""

	for _, ch := range strings.ToLower(text) {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			clean += string(ch)
		}
	}

	// cek palindrome
	left := 0
	right := len(clean) - 1

	for left < right {
		if clean[left] != clean[right] {
			fmt.Println(false)
			return
		}
		left++
		right--
	}

	fmt.Println(true)
}

func main() {
	isPalindrome("Kasur ini rusak") // true
	isPalindrome("tamaT")   // true
	isPalindrome("Aku Usa") // false
}
