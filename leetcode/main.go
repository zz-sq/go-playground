package main

import "fmt"

// 题集：https://leetcode.cn/studyplan/top-interview-150/
// https://leetcode.cn/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(maxProfit(nums1))
}
