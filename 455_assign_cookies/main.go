package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g) // 1 2 3
	sort.Ints(s) // 1 1

	var i, j int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{2}))
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(findContentChildren([]int{10, 9, 8, 7, 10, 9, 8, 7}, []int{10, 9, 8, 7}))
}
