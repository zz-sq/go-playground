package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "其实就是rune"
	fmt.Println(len(s))                    // "16"
	fmt.Println(utf8.RuneCountInString(s)) // "8"

	s2 := "Go语言四十二章经"
	for k, v := range s2 {
		fmt.Printf("k: %d,v: %c == %d\n", k, v, v)
	}

	var sliceInt = []int{1, 2, 4}
	for k1, v1 := range sliceInt {
		fmt.Printf("k1: %d, v1:  %d\n", k1, v1)
	}

	var scoreMap = make(map[string]int)
	scoreMap["alice"] = 100
	scoreMap["bob"] = 97

	for k1, v1 := range scoreMap {
		fmt.Printf("k1: %s, v1:  %d\n", k1, v1)
	}

}
