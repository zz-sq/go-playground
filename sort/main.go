package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type personSlice []person

func (s personSlice) Len() int           { return len(s) }
func (s personSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
	aa := []int{3, 5, 4, -1, 9, 11, -14}
	sort.Ints(aa)
	fmt.Println(aa)
	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	sort.Strings(ss)
	fmt.Println(ss)
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("After reverse: %v\n", ss)

	a := personSlice{
		{
			Name: "AAA",
			Age:  55,
		},
		{
			Name: "BBB",
			Age:  22,
		},
		{
			Name: "CCC",
			Age:  0,
		},
		{
			Name: "DDD",
			Age:  22,
		},
		{
			Name: "EEE",
			Age:  11,
		},
	}
	sort.Sort(a)
	fmt.Println("Sort:", a)

	sort.Stable(a)
	fmt.Println("Stable:", a)

	b := []person{
		{
			Name: "AAA",
			Age:  55,
		},
		{
			Name: "BBB",
			Age:  22,
		},
	}
	// does an in-place sort
	sort.Slice(b, func(i, j int) bool {
		return b[i].Age < b[j].Age
	})
	fmt.Println(b)

}
