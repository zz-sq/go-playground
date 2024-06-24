package main

import (
	"fmt"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([][]rune, numRows)
	i, flag := 0, -1
	for _, v := range s {
		res[i] = append(res[i], v)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	res2 := make([]rune, 0)
	for i := 0; i < numRows; i++ {
		res2 = append(res2, res[i]...)
	}
	return string(res2)
}

// PAHNAPLSIIGYIR
// PAHNAPLSIIGYIR
// 0 -> 2n -1
func main() {
	fmt.Println(convert("AB", 1))
}
