package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n == 0 {
		return 0
	}

	totalDuration := 0
	startTime := timeSeries[0]
	endTime := startTime + duration

	for i := 1; i < n; i++ {
		if timeSeries[i] > endTime {
			totalDuration += endTime - startTime
			startTime = timeSeries[i]
		}
		endTime = timeSeries[i] + duration
	}
	totalDuration += endTime - startTime
	return totalDuration
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
}
