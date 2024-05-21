package main

import "fmt"

func countSegments(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}

	var word bool

	for _, c := range s {
		if c != ' ' {
			if !word {
				res++
				word = true
			}
		} else {
			word = false
		}
	}
	return res
}

func main() {
	fmt.Println(countSegments("a"))
	fmt.Println(countSegments("Hello, my name is John"))
	fmt.Println(countSegments("                "))
}
