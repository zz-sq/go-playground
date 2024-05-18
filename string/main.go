package main

import (
	"fmt"
	"unicode"
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

	// 双引号中的字符串可以包含转义字符，而反引号中的字符串则不行
	str1 := "hello, world\n"
	fmt.Printf("%s", str1)
	str2 := `hello, world\n`
	fmt.Printf("%s\n", str2)

	str2 += "this is a concat"
	fmt.Printf("%s\n", str2)

	numStr := "123"
	// %c 表示输出一个字符
	var c byte = 'A'
	fmt.Printf("%c\n", c)
	fmt.Printf("%d\n", c)

	for _, v := range numStr {
		fmt.Printf("%c is a digit: %t\n", v, unicode.IsDigit(v))
	}

}
