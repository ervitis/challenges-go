package main

import "fmt"

func arrangeCoins(n int) int {
	rowsCompleted := 0
	rowCoins := 1
	total := 0

	for total+rowCoins <= n {
		rowsCompleted++
		total += rowCoins
		rowCoins++
	}
	return rowsCompleted
}

func main() {
	fmt.Println(arrangeCoins(5))
}
