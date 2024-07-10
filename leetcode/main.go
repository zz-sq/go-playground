package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题集：https://leetcode.cn/studyplan/top-interview-150/
// 199 637 102 103
func checkSymmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if q != nil && p != nil {
		return q.Val == p.Val && checkSymmetric(p.Left, q.Right) &&
			checkSymmetric(p.Right, q.Left)
	} else {
		return false
	}

}

func main() {
	// arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// res := averageOfLevels("abba", "dosg cat cat dog")
	// fmt.Println(res)
}
