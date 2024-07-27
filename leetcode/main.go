package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if coin > i || dp[i-coin] == -1 {
				break
			}
			if dp[i] == -1 {
				dp[i] = dp[i-coin] + 1
			} else {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{2, 5, 10, 1}, 27))
}
