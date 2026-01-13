package main

import (
	"fmt"
)

func printOkYes(n int) {
	for i:= 1 ; i <= n ; i++ {
		if i%3 == 0 && i%4 == 0 {
			fmt.Print(" OK Yes ")
		}else if i%3 == 0 {
			fmt.Print(" OK ")
		}else if i%4 == 0 {
			fmt.Print(" Yes ")
		}else {
			fmt.Printf("%d ",i)
		}
	}
}

func countOk (n int) int {
	count := 0
	for i:= 1 ; i <= n ; i++ {
		if i%3 == 0 && i%4 != 0 {
			count++
		}
	}
	return count
}

func sumOk (n int) int {
	sum := 0
	for i:= 1 ; i <= n ; i++ {
		if i%3 == 0 && i%4 != 0 {
			sum += i
		}
	}
	return sum
}

func sumOkContinue (n int) int {
	sum := 0
	for i:= 1 ; i <= n ; i++ {
		if i%3 == 0 &&i%4 == 0 {
			break
		}else if i%3 == 0 {
			sum += i
		}else if i%4 == 0 {
			continue
		}else {
			continue
		}
	}
	return sum
}

func main()  {
	n := 15
	printOkYes(n)	
	fmt.Println("countOK : ",countOk(n))
	fmt.Println("sumOK : ",sumOk(n))

	fmt.Println("sumOkContinue : ",sumOkContinue(n))

	counter := 0

//for looping but old school
target :
	fmt.Println("Counter",counter)
	counter++
	if counter < 10 {
		goto target	
	}
}
