package main

import "fmt"

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0, 1, 0, 3, 12, 0, 0, 0, 0, 5, 6, 3, 4, 5, 3}
	moveZeroes(nums)
	fmt.Println(nums)
}
