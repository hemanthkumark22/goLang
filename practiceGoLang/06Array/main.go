package main

import "fmt"

func main() {

	//Array is a collection of similar data types, It similar to slice but it has fixed size and it is not dynamic in nature
	//Array Literal
	mark := [10]int{10, 20, 30}
	fmt.Println(len(mark), cap(mark), "length and capacity of mark array")

	name := [4]string{"Hem", "Tem", "Ram", "Abdul"}
	fmt.Println(name)

	myArray := [3]int{1, 2, 3}
	myArrayCopy := myArray
	myArray[2] = 4 // it will change the value for the initial array but not for the copy of the array because it is a value type
	fmt.Println(myArray, myArrayCopy)
}