package main

import (
	"fmt"
	"strings"
)

func repeatedPattern(s string) bool {
	l := len(s)

	for i := 1; i <= l/2; i++ {
		if l%i == 0 {
			c := l / i
			substr := s[:i]

			if strings.Repeat(substr, c) == s {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(repeatedPattern("abab"))
	fmt.Println(repeatedPattern("aba"))
}
