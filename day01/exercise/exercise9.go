package main

import "fmt"

func findFibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
		fmt.Println(b)
	}
	return b
}

func main() {
	fmt.Printf("%d ", findFibonacci(10))
}
