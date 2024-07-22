package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题集：https://leetcode.cn/studyplan/top-interview-150/

// @lc code=start
type Trie struct {
	children [26]*Trie
	isWord   bool
	word     string
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, ch := range word {
		idx := ch - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &Trie{}
		}
		cur = cur.children[idx]
	}
	cur.isWord = true
	cur.word = word
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return res
	}
	m, n := len(board), len(board[0])
	root := &Trie{}
	for _, word := range words {
		root.Insert(word)
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(node *Trie, i int, j int, visited [][]bool)
	dfs = func(node *Trie, i int, j int, visited [][]bool) {
		if i < 0 || j < 0 || i >= m || j >= n || visited[i][j] {
			return
		}
		idx := board[i][j] - 'a'
		next := node.children[idx]
		if next == nil {
			return
		}
		if next.isWord {
			res = append(res, next.word)
			next.isWord = false
		}
		visited[i][j] = true
		dfs(next, i+1, j, visited)
		dfs(next, i-1, j, visited)
		dfs(next, i, j+1, visited)
		dfs(next, i, j-1, visited)
		visited[i][j] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(root, i, j, visited)
		}
	}
	return res
}

func main() {
	// arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// res := averageOfLevels("abba", "dosg cat cat dog")
	// fmt.Println(res)
	fmt.Println(findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}))
}
