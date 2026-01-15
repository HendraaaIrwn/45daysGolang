package main

import "fmt"

func printPyramidLeftAligned(n int) {
	for i:=0 ; i < n ; i++ {
		for j:=0 ; j <= i ; j ++ {
			fmt.Print("* ")
		}
		fmt.Println();
	}

}

func main() {
	printPyramidLeftAligned(5)
	printReversePyramid(5)
}