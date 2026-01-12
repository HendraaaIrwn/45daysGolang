package main

import "fmt"

func printPyramid (n int ) {
	for i := 0 ; i <= n ; i++ {

	//print spasi
	for j := 0 ; j <= n-i ; j ++ {
		fmt.Print(" ")
	}

	//print bintang 
	for k := 0 ; k <= i ; k++ {
		fmt.Print("*")
	}
		
	fmt.Println()
	}
}

func printReversePyramid (n int ) {
	for i := 0 ; i <= n ; i++ {

	//print spasi
	for k := 0 ; k <= i ; k++ {
		fmt.Print(" ")
	}

	//print bintang 
	for j := 0 ; j <= n-i ; j ++ {
		fmt.Print("*")
	}
	fmt.Println()
	}
}


func main() {
	printPyramid(4)
	printReversePyramid(4)
}