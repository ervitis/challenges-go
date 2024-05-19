package main

import "fmt"

func longestPalindrome(s string) int {
	set := make(map[byte]int)
	count := 0
	orig := len(s)
	for i := 0; i < len(s); i++ {
		set[s[i]]++
	}
	for k, v := range set {
		if v%2 == 0 {
			count += v
		} else {
			set[k]--
			count += set[k]
		}
	}
	if count < orig {
		return count + 1
	}

	return count
}

func main() {
	fmt.Println(longestPalindrome("aab"))
	fmt.Println(longestPalindrome("aaba"))
	fmt.Println(longestPalindrome("abccccdd"))
}
