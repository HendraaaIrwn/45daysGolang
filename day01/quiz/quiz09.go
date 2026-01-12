package main

import "fmt"

func reverse(s string) {
	bytes := []byte(s)

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	fmt.Println(string(bytes))
}

func main() {
	reverse("ABCD") // DCBA
	reverse("tamaT") // TamaT
	reverse("XYnb") // bnYX
}
