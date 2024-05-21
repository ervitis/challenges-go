package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	set := make(map[int]struct{})
	m := -1
	for _, num := range nums {
		if m < num {
			m = num
		}
		set[num] = struct{}{}
	}

	if m < len(nums) {
		m = len(nums)
	}

	var result []int
	for i := 1; i <= m; i++ {
		if _, ok := set[i]; !ok {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
