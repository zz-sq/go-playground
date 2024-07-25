package main

import (
	"container/heap"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	res := make([][]int, 0, min(k, n*m))
	h := make(hp, min(k, n))
	for i := range h {
		h[i] = tuple{nums1[i] + nums2[0], i, 0}
	}
	for len(res) < k {
		p := heap.Pop(&h).(tuple)
		i, j := p.i, p.j
		res = append(res, []int{nums1[i], nums2[j]})
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return res
}

type tuple struct{ s, i, j int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].s < h[j].s }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func maxHeapify(a []int, i, n int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < n && a[l] > a[largest] {
		largest = l
	}
	if r < n && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, n)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println(findKthLargest(arr, 3))
}
