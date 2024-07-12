package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题集：https://leetcode.cn/studyplan/top-interview-150/

func minWindow(s string, t string) string {
	needs := map[byte]int{}
	current := map[byte]int{}
	for _, c := range t {
		needs[byte(c)]++
	}
	valid, l, r, start, end := 0, 0, 0, -1, len(s)+1
	for r < len(s) {
		if needs[s[r]] > 0 {
			current[s[r]]++
			if current[s[r]] == needs[s[r]] {
				valid++
			}
		}
		for valid == len(needs) {
			if end-start > r-l {
				start = l
				end = r + 1
			}
			if needs[s[l]] > 0 {
				if needs[s[l]] == current[s[l]] {
					valid--
				}
				current[s[l]]--
			}
			l++
		}
		r++
	}
	if start == -1 {
		return ""
	} else {
		return s[start:end]
	}
}

func main() {
	// arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// res := averageOfLevels("abba", "dosg cat cat dog")
	// fmt.Println(res)
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
