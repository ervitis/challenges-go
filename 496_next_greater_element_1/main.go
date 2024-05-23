package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	next := make(map[int]int)

	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			next[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	res := make([]int, 0)
	for _, num := range nums1 {
		if v, ok := next[num]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}
