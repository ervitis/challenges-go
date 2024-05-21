package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	nums := strconv.FormatInt(int64(num), 2)
	for i := 0; i < len(nums); i++ {
		if nums[i] == '1' {
			nums = nums[:i] + "0" + nums[i+1:]
		} else {
			nums = nums[:i] + "1" + nums[i+1:]
		}
	}
	t, _ := strconv.ParseInt(nums, 2, 64)
	return int(t)
}

func main() {
	fmt.Println(findComplement(5))
}
