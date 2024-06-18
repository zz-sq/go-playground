package main

import "fmt"

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func productExceptSelf(nums []int) []int {
	tmp, zero := 1, 0
	for _, num := range nums {
		if num == 0 {
			zero += 1
		} else {
			tmp *= num
		}
	}
	ans := make([]int, len(nums))
	for i, num := range nums {
		if zero == 1 && num == 0 {
			ans[i] = num
		} else if zero > 1 {
			ans[i] = 0
		} else {
			ans[i] = tmp / num
		}
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
}
