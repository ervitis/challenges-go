package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	width := int(math.Sqrt(float64(area)))

	for area%width != 0 {
		width--
	}
	height := area / width

	return []int{height, width}
}

func main() {
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(122122))
}
