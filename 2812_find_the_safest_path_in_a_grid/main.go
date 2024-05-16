package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/find-the-safest-path-in-a-grid

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	dp[0][0] = math.MaxInt32

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if i > 0 {
					dp[i][j] = max(dp[i][j], min(dp[i-1][j], abs(i-1-i)+abs(j-j)))
				}
				if j > 0 {
					dp[i][j] = max(dp[i][j], min(dp[i][j-1], abs(i-i)+abs(j-1-j)))
				}
			}
		}
	}

	return dp[n-1][n-1]
}

func main() {
	fmt.Printf("result: %+v\n", maximumSafenessFactor([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Printf("result: %+v\n", maximumSafenessFactor([][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}))
	fmt.Printf("result: %+v\n", maximumSafenessFactor([][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}))
}
