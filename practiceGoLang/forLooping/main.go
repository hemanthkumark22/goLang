package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, "Hello World")
	}

	//Print all even numbers from 1 to 20
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is an even number")
		}
	}
	////Find the sum of numbers from 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum, "total sum")

}
