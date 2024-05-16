package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for i := 1; i < n; {
		i *= 3
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(3))
	fmt.Println(isPowerOfThree(4))
}
