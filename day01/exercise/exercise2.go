package main

import "fmt"


func printHollowBox (n int) {
	// for i := 0 ; i < n ; i++ {
	// 	for j := 0 ; j < n ; j++{
	// 		if i == 1 || i == n {
	// 			fmt.Print("* ")
	// 		} else if i != 1 && i != n && j != 1 && j != n{
	// 			fmt.Print("* ")
	// 		}
	// 	}
	// }

	for i := 0 ; i < n ; i++ {
		for j := 0 ; j < n ; j++{
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func main() {
	printHollowBox(5)
}