package main

import "fmt"

func printReversePyramid(n int) {
	for i:= 0; i < n ; i++ {
		for j:=n-i ; j > 0 ; j-- {
			fmt.Print("* ")
		}
		fmt.Println();
	}
}

func main() {
	printReversePyramid(5)
}