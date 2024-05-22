package main

import (
	"fmt"
)

func licenseKeyFormatting(s string, k int) string {
	c := 0
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		if c == k {
			res = "-" + res
			c = 0
		}
		if s[i] != '-' {
			ch := s[i]
			if ch >= 'a' && ch <= 'z' {
				ch -= 32
			}
			res = string(ch) + res
			c++
		}
	}
	if len(res) > 0 && res[0] == '-' {
		res = res[1:]
	}
	return res
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
	fmt.Println(licenseKeyFormatting("--a-a-a-a--", 2))
	fmt.Println(licenseKeyFormatting("---", 3))
}
