package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arrAge = [5]int{18, 20, 18, 22, 20}
	fmt.Println(arrAge)

	var arrSlice = arrAge[1:3]
	fmt.Println(arrSlice)
	fmt.Printf("type of arrAge[1:3]: %s\n", reflect.TypeOf(arrSlice))

	fmt.Printf("type of make([]string, 10): %s\n", reflect.TypeOf(make([]string, 10)))

	var arrRoom [20]int
	var arrBed = new([20]int)
	fmt.Printf("type of [20]int: %s\n", reflect.TypeOf(arrRoom))
	fmt.Printf("type of new([20]int): %s\n", reflect.TypeOf(arrBed))
	fmt.Printf("type of []int{1, 2, 3, 4, 5}: %s\n", reflect.TypeOf(arrSlice))

}
