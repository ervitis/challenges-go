package main

import "fmt"

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func main() {
	fmt.Println(countBits(3))
	fmt.Println(countBits(5))
}
