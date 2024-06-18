package main

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for j, v := range nums {
		for i := j + 1; i < j+v; i++ {
			if dp[i] == 0 {
				dp[i] = dp[j] + 1
				if i == n-1 {
					return dp[i]
				}
			}
		}
	}
	return dp[n-1]
}

func main() {

}
