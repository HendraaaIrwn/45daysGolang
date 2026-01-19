package main

import "fmt"

func main() {
	//1. declare slice
	var slice1 []string
	slice1 = append(slice1, "1", "2", "3")
	fmt.Printf("slice1:[value:%#v\taddress:%p]\n", slice1, &slice1)

	//2. short variable
	slice2 := []string{}
	fmt.Printf("slice2:[value:%#v\taddress:%p]\n", slice2, &slice2)

	//3. short variable
	slice3 := make([]string, 5, 8)
	fmt.Printf("slice3:[value:%#v\taddress:%p]\n", slice3, &slice3)

	//5. with values
	slice5 := []string{"A", "B", "C", "D", "E"}
	fmt.Printf("slice5:[value:%#v\taddress:%p]\n", slice5, &slice5)

	//6. short variable
	slice6 := make([]string, 5, 8)
	slice6 = append(slice6, "A", "B", "C", "D", "E")
	fmt.Printf("slice6:[value:%#v\taddress:%p]\n", slice6, &slice6)

	for i := range slice6 {
		fmt.Printf("[value:%#v\taddress:%p]\n", slice6[i], &slice6[i])
	}

	//jika ingin pake function
	display(slice6)
}

func display(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[value:%#v\taddress:%p]\n", slice[i], &slice[i])
	}
}