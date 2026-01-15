package main

import "fmt"

func printHorizontalNumberOnBoxStar (n int) {
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if j==i {
				fmt.Printf("%d ",j)
				
			}else{
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}


func printHorizontalOddNumberOnBoxStar (n int) {
	oddNumber := 0
	for i := 0; i < n; i++ {
		for j := 1; j < n+1; j++ {
			if j==i+1 {
				oddNumber = i+j
				fmt.Printf("%d ",oddNumber)
				
			}else{
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// printHorizontalNumberOnBoxStar(5)
	printHorizontalOddNumberOnBoxStar(5)
}