package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	m := -1
	var c int
	for _, n := range nums {
		if n == 1 {
			c++
		} else {
			if c > m {
				m = c
			}
			c = 0
		}
	}

	// last
	if c > m {
		m = c
	}

	return m
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
