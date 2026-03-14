package main

import "fmt"

func main() {

	marks := 10

	switch marks {
	case 10:
		fmt.Println("you are failed")
	case 50:
		fmt.Println("you are passed")

	default:
		fmt.Println("marks: ", marks)
	}
}
