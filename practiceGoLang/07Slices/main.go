package main

import "fmt"

func main() {
	names := make([]string, 3, 5)
	//names := []string{"test1", "test2", "test3"}
	//fmt.Println(names, names[len(names)-1])

	//change the value of the slice
	names[0] = "test0"
	//fmt.Println(names, names[len(names)-3])
	//append new value to the slice
	names = append(names, "test4", "test5")
	//fmt.Println(names, len(names), cap(names))
	// names = append(names, "test11", "test12", "test13","test14")
	fmt.Println(names, len(names), cap(names))
	rollNo := []string{"1", "2", "3", "4", "5"}
	fmt.Println(rollNo)
	//append two string variable "rollNo to names"
	names = append(names, rollNo...)
	fmt.Println(names)
}
