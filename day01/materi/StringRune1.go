package main

import (
	"fmt"
)

func countVowels(s string) (int, int) {
	vowelCount := 0
	consonantCount := 0

	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount++
		default:
			consonantCount++
		}
	}
	return vowelCount, consonantCount
}

func countVowelsConsonants(s string) (int, int, string) {
	vowelCount := 0
	consonantCount := 0
	consonantString := " "

	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowelCount++
		} else if char != ' ' && char != '\n' && char != '\t' && char != '\r' {
			consonantCount++
			consonantString += string(char)
		}
	}
	return vowelCount, consonantCount, consonantString
}	

func countVowelsConsonantsAlternative(s string) (int, int) {
	vowelCount := 0
	consonantCount := 0

	for i:=0 ; i < len(s); i++ {
		char := s[i]
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowelCount++
		} else if char != ' ' && char != '\n' && char != '\t' && char != '\r' {
			consonantCount++
		}
	}
	return vowelCount, consonantCount
}	


func main() {
	str1 := `hello 🌍
	saya sedang belajar go`

	str2 := "Folder in \"C:\\User\\Admin\""
	fmt.Println(str2)

	vowels, consonants, consonantString := countVowelsConsonants(str1)
	println("Vowels:", vowels)
	println("Consonants:", consonants)
	println("Consonant String:", consonantString)

	vowels2, consonants2, consonantString2 := countVowelsConsonants(str2)
	println("Vowels:", vowels2)
	println("Consonants:", consonants2)
	println("Consonant String:", consonantString2)

	username := "gopher"

	for _ , char := range username {
		fmt.Print(char, " \n")
	}

	for i , char := range username {
		// fmt.Printf("index: %d, char: %v\n", i, string(char))
		fmt.Printf("index: %d, char: %c\n", i, char)
	}

	runeUsername := []rune(username)
	for i := 0; i < len(runeUsername); i++ {
		fmt.Printf("index: %d, char: %c\n", i, runeUsername[i])
	}

	var (
		tab = '\t'
		newLine = '\n'
		backslash = '\\'
		singleQuote = '\''
		doubleQuote = '"'
		hagul = '한'
		arabic = 'م'
		hexadecimal = '\x48'
		unicodeChar = '\u0048'
	)

	println("tab :", tab)
	println("newLine :", newLine)
	println("backslash :", backslash)
	println("singleQuote :", singleQuote)
	println("doubleQuote :", doubleQuote)
	println("hagul :", hagul)
	println("arabic :", arabic)
	println("hexadecimal :", hexadecimal)
	println("unicodeChar :", unicodeChar)
}