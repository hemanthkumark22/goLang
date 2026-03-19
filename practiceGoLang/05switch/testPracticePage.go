package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = float64(a)

	fmt.Println(b)

	var c float64 = 9.7
	var d int = int(c)

	fmt.Println(d)
}
