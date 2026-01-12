package main

import "fmt"

func findDivisor (n int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Printf("%d ",i)
		}
	}
fmt.Println()
}

func main() {
	findDivisor(6)
	findDivisor(24)
	findDivisor(7)
}