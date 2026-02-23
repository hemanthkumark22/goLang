package main

import "fmt"

func main() {
	percentage := 100
	if percentage < 35 && percentage < 100 && percentage > 0 {
		fmt.Println("You are Failed")
	} else if percentage >= 35 && percentage < 80 && percentage < 100 {
		fmt.Println("you are passed")
	} else if percentage >= 80 && percentage <= 100 {
		fmt.Println("you are passed with distinction")
	} else if percentage > 100 || percentage <= 0 {
		fmt.Println("Invalid percentage")
	}
	priceItem1 := -100
	priceItem2 := 200
	priceItem3 := 101
	totalPrice := priceItem1 + priceItem2 + priceItem3
	fmt.Println("Total price of 3 items:", totalPrice)
	if totalPrice >= 400 {
		fmt.Println("Declined limit exceeded")
	} else {
		fmt.Println("Approved Transaction")
	}
}
