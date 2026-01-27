package main

import "fmt"

func displaySlice(name string, slice []string) {
	fmt.Printf("%s: %T len: %d, cap %d, %v\n", name, slice, len(slice), cap(slice), slice)
	for i := range slice {
		fmt.Printf("[%d] addr : [%p]\t value:[%s]\n", i, &slice[i], slice[i])
	}
}

func main() {
	// var slice1 []string

	// slice2 := []string{}

	// slice3 := make([]string, 5)

	// slice4 := make([]string, 5, 8)

	// slice5 := []string{"a", "b", "c", "d"}

	// displaySlice("slice1", slice1)
	// displaySlice("slice2", slice2)
	// displaySlice("slice3", slice3)
	// displaySlice("slice4", slice4)
	// displaySlice("slice5", slice5)

	slice1 := []string{"A","B","C","D","E"}
	// slice2 := slice1[2:4]
	// slice2[0] = "UP"
	// slice2 = append(slice2, "DOWN")
	// slice3 := slice1[2:4]
	// slice3 = append(slice3, "XX") 
	
	slice3 := slice1[2:4:4]
	slice3 = append(slice3, "IO")

	displaySlice("slice1", slice1)
	// displaySlice("slice2", slice2)
	displaySlice("slice3", slice3)




}
