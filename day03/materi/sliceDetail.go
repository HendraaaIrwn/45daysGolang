package main

import (
	"fmt"
)

func main() {
	//1. init slice string
	slice1 := []string{"A", "B", "C", "D", "E", "F", "G"}
	slice2 := slice1[2:6]

	slice1 = append(slice1, "Z", "X", "Y", "XXX")

	infoSlice(slice1)

	sliceOutCap := slice1[7:12]

	infoSlice(sliceOutCap)

	//slice2 = append(slice2, "XX")

	//2. no side-effect,low=2,high=length element,
	//max=4 (4-2) capacity yg akan ditambah
	slice3 := slice1[2:4:4]
	slice3 = append(slice3, "YY")

	//3. copy all slices
	slice4 := slice1[:]
	slice4[0] = "AB"

	//4. copy slice
	slice5 := make([]string, len(slice1))
	copy(slice5, slice1)
	slice5[0] = "XYZ"

	infoSlice(slice1)
	infoSlice(slice2)
	infoSlice(slice3)
	infoSlice(slice4)
	infoSlice(slice5)

	//5. comparing slice
	/* fmt.Println("slice1 == slice4", reflect.DeepEqual(slice1, slice4))
	fmt.Println("slice1 == slice5", reflect.DeepEqual(slice1, slice5)) */

	//6. delete slice
	slice5 = append(slice5[:1], slice5[2:]...)
	infoSlice(slice5)
}

func infoSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] [%p] [%s]\n", i, &slice[i], slice[i])
	}
}