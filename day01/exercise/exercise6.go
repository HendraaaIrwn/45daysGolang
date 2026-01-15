package main

import "fmt"

func printLshapeNumberOnBoxStar (n int) {
	Number := 0
	for i := 0; i < n; i++ {
		for j := 1; j < n+1; j++ {
			if j==1 {
				Number += j
				fmt.Printf("%d ",Number)
				
			}else if i==n-1{
				Number = i+j
				fmt.Printf("%d ",Number)
			}else{
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}

func main() {
	printLshapeNumberOnBoxStar(5)
}