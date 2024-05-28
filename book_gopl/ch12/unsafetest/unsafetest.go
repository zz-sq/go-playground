package main

import (
	"fmt"
	"unsafe"
)

type X struct {
	a bool
	b int16
	c []int
	d string
}

func main() {
	var x = X{}
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(x.a))
	fmt.Println(unsafe.Sizeof(x.b))
	fmt.Println(unsafe.Sizeof(x.c))
	fmt.Println(unsafe.Sizeof(x.d))
}
