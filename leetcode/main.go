package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reformatNumber(number string) string {
	re := regexp.MustCompile(`[\s-]+`)
	s := re.ReplaceAllString(number, "")
	l := len(s)
	var ss []string
	for i := 0; i < l; i += 3 {
		if l-i <= 4 {
			if l-i == 4 {
				ss = append(ss, s[i:i+2])
				ss = append(ss, s[i+2:])
			} else {
				ss = append(ss, s[i:])
			}
			break
		}
		ss = append(ss, s[i:i+3])
	}
	return strings.Join(ss, "-")
}

func main() {
	fmt.Println(reformatNumber("1-23-45 6"))
	fmt.Println(reformatNumber("123 4-567"))
	fmt.Println(reformatNumber("123 4-5678"))
}
