package main

import "fmt"

type d struct {
	i int
	o int
}

func firstUniqChar(s string) int {
	m := make(map[string]*d)
	for i := 0; i < len(s); i++ {
		if c, ok := m[string(s[i])]; !ok {
			m[string(s[i])] = &d{i: i, o: 1}
		} else {
			m[string(s[i])].o = c.o + 1
		}
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if v, ok := m[string(c)]; ok && v.o == 1 {
			return v.i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
