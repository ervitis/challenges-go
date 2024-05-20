package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	m, sec, th := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v == th || v == m || v == sec {
			continue
		}
		switch {
		case v > m:
			m, sec, th = v, m, sec
		case v > sec:
			sec, th = v, sec
		case v > th:
			th = v
		}
	}
	if th == math.MinInt64 {
		return m
	}
	return th
}

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
	fmt.Println(thirdMax([]int{5, 2, 2}))
}
