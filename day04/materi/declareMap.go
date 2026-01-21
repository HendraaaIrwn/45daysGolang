package main

import "fmt"

func main() {
	products := make(map[string]float64, 10)
	products["tv"] = 3_500_000.00
	products["book"] = 100_000.00
	products["maousepad"] = 150_000.00

	products["book"] = 89_000.00

	delete(products, "book")

	for key, v := range products {
		fmt.Printf("key:[%s] value:[%.2f]\n", key, v)
	}
}
