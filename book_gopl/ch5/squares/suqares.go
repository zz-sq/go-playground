package main

import "fmt"

func squaresFunc() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squaresFunc()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
