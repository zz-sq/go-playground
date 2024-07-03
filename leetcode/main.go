package main

import (
	"fmt"
)

// 0 1 0
// 0 0 1
// 1 1 1
// 0 0 0
// Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
// 题集：https://leetcode.cn/studyplan/top-interview-150/
func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			count := lives(board, i, j)
			if board[i][j] == 1 && (count < 2 || count > 3) {
				board[i][j] = 3
			} else if board[i][j] == 0 && count == 3 {
				board[i][j] = 4
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 3 {
				board[i][j] = 0
			} else if board[i][j] == 4 {
				board[i][j] = 1
			}
		}
	}
}

func lives(board [][]int, i int, j int) int {
	res := 0
	for m := i - 1; m <= i+1; m++ {
		for n := j - 1; n <= j+1; n++ {
			if m == i && n == j {
				continue
			}
			if n >= 0 && m >= 0 && m < len(board) && n < len(board[0]) {
				if board[m][n] == 1 || board[m][n] == 3 {
					res++
				}
			}
		}
	}
	return res
}

func main() {
	arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(arr1)
	fmt.Println(arr1)
}
