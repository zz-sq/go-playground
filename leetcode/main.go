package main

import (
	"fmt"
	"strings"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
// 141 2 21 138 92
func wordPattern(pattern string, s string) bool {
	word := strings.Split(s, " ")
	if len(pattern) != len(word) {
		return false
	}
	patternMaps := make(map[rune]string)
	wordMaps := make(map[string]rune)
	for i, p := range pattern {
		if c, ok := patternMaps[p]; ok {
			if c != word[i] {
				return false
			}
		} else {
			patternMaps[p] = word[i]
		}
		if c, ok := wordMaps[word[i]]; ok {
			if c != p {
				return false
			}
		} else {
			wordMaps[word[i]] = p
		}
	}
	return true
}

func main() {
	// arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	res := wordPattern("abba", "dosg cat cat dog")
	fmt.Println(res)
}
