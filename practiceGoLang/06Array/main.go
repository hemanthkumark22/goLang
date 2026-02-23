package main

import "fmt"

func main() {

	//Array is a collection of similar data types, It similar to slice but it has fixed size and it is not dynamic in nature
	//Array Literal
	mark := [10]int{10, 20, 30}
	fmt.Println(len(mark), cap(mark), "length and capacity of mark array")

	name := [4]string{"Hem", "Tem", "Ram", "Abdul"}
	fmt.Println(name)
}
