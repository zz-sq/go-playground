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

func calculate(s string) int {
	stack := []int{1}
	ans, sign := 0, 1
	for i := 0; i < len(s); {
		c := s[i]
		switch c {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}

func main() {
	// arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// res := averageOfLevels("abba", "dosg cat cat dog")
	// fmt.Println(res)
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
