package main

import (
	"errors"
	"fmt"
)

func addToCart(stock *int, cartQuantity int) (int, error) {
	//check if the addCartItem negatif or no
	if cartQuantity < 0 {
		return *stock, errors.New("Item quantity can't be less than 0")
	}

	//check if addCartItem larger than stock
	if cartQuantity > *stock {
		return *stock, errors.New("Insufficient stock, reduce shopping items")
	}

	*stock -= cartQuantity
	return *stock, nil
}

func checkError(err error, cartQuantity int, remainingStock int) {
	if err != nil {
		fmt.Print("Error : ", err)
		return
	} else {
		fmt.Printf("add success ! added to cart:%d, remaining stock:%d\n", cartQuantity, remainingStock)
	}

}
func main() {

	stock := 100
	cartQuantity := -10

	remainingStock, err := addToCart(&stock, cartQuantity)
	checkError(err, cartQuantity, remainingStock)

}
