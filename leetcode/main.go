package main

import (
	"fmt"
	"strings"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func fullJustify(words []string, maxWidth int) []string {
	ret := make([][]string, len(words))
	wc := make([]int, 0)
	i, curLen := 0, 0
	for _, word := range words {
		if len(ret[i])+curLen+len(word) <= maxWidth {
			ret[i] = append(ret[i], word)
			curLen += len(word)
		} else {
			wc = append(wc, curLen)
			curLen = len(word)
			i++
			ret[i] = append(ret[i], word)
		}
	}
	wc = append(wc, curLen)
	res := make([]string, 0)
	for m := 0; m <= i; m++ {
		line := ret[m]
		leftSpace := maxWidth - wc[m]
		var perWord, extraSpace int
		if len(line) == 1 {
			perWord = leftSpace
			extraSpace = 0
		} else if m == i {
			perWord = 1
			extraSpace = 0
		} else {
			perWord = leftSpace / (len(line) - 1)
			extraSpace = leftSpace % (len(line) - 1)
		}
		var builder strings.Builder
		for j, l := range line {
			builder.WriteString(l)
			if j != 0 && j == len(line)-1 {
				continue
			}
			for k := 0; k < perWord; k++ {
				builder.WriteRune(' ')
			}
			if j < extraSpace {
				builder.WriteRune(' ')
			}
		}
		if m == i {
			len := builder.Len()
			for k := 0; k < maxWidth-len; k++ {
				builder.WriteRune(' ')
			}
		}
		res = append(res, builder.String())
	}
	return res
}

func main() {
	fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
}
