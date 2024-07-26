package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func addBinary(a string, b string) string {
	res := []byte{}
	i, j, flag := len(a)-1, len(b)-1, byte(0)

	for i >= 0 || j >= 0 || flag > 0 {
		if i >= 0 {
			flag += a[i] - '0'
			i--
		}
		if j >= 0 {
			flag += b[j] - '0'
			j--
		}
		res = append([]byte{flag%2 + '0'}, res...)
		flag /= 2
	}
	return string(res)
}

func main() {
	fmt.Println(addBinary("1011", "1010"))
}
