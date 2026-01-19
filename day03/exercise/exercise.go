package main

import (
	"fmt"
)

func display(name string, slice []int) {
	fmt.Printf("%s:[%p],Length[%d] Capacity[%d]\n", name, &slice, len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] addr: [%p]\tvalue:[%d]\n", i, &slice[i], slice[i])
	}
}

func main() {
	numbers := make([]int, 2, 10)
	numbers = append(numbers[:0], 5, 10)

	display("numbers", numbers)

	slices1 := numbers[0:1:1]
	display("slices1", slices1)

	slices2 := numbers[1:2:2]
	display("slices2", slices2)

	numbers = append(numbers[0:1:1], 6, 7, 8, 9, 10)
	display("numbers", numbers)
}
