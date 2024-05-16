package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	req := make(map[int]int)
	for _, v := range nums1 {
		req[v]++
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		if c, ok := req[v]; ok && c > 0 {
			res = append(res, v)
			req[v]--
		}
	}
	return res
}

func main() {
	fmt.Println(intersect([]int{1, 2, 3}, []int{1, 3, 2}))
	fmt.Println(intersect([]int{1}, []int{2}))
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 5, 9}, []int{9, 4, 9, 8, 4}))
}
