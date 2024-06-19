package main

import "fmt"

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumGas, sumCost, j := 0, 0, 0
		for j < n {
			k := (i + j) % n
			sumGas += gas[k]
			sumCost += cost[k]
			if sumCost > sumGas {
				break
			}
			j++
		}
		if j == n {
			return i
		} else {
			i += j + 1
		}
	}
	return -1
}

// 1 2 3 4 5
// 3 4 5 1 2
// -2 -2 -2 4 3
func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
