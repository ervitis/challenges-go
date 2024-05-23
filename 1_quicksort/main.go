package main

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quicksort(a []int, low, high int) {
	if low < high {
		pivot := partition(a, low, high)
		quicksort(a, low, pivot-1)
		quicksort(a, pivot+1, high)
	}
}

func main() {
	arr := []int{0}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arr = []int{1, 5, 6, 2, 3, 9, 8, 1, 2, 7, 0, 9, 0, 7}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
