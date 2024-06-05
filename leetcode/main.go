package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	var count = make(map[rune]int)
	for _, v := range magazine {
		count[v] += 1
	}
	for _, v := range ransomNote {
		count[v] -= 1
		if count[v] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aac", "aab"))

}
