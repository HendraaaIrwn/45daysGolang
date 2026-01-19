package main

func removeDuplicateLetters(s string) string {
	seen := make(map[rune]bool)
	result := []rune{}

	for _, char := range s {
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}

	return string(result)
}

func main() {
	
}