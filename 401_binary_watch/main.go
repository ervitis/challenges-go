package main

import "fmt"

func convertTime(h, m int) string {
	return fmt.Sprintf("%d:%02d", h, m)
}

func backtrack(h, m, pos, turnedOn int, res *[]string) {
	if turnedOn == 0 {
		*res = append(*res, convertTime(h, m))
	}
	for i := pos; i < 10; i++ {
		if i < 4 {
			nh := h | (1 << i)
			if nh <= 11 {
				backtrack(nh, m, i+1, turnedOn-1, res)
			}
		} else {
			nm := m | (1 << (i - 4))
			if nm <= 59 {
				backtrack(h, nm, i+1, turnedOn-1, res)
			}
		}
	}
}

func readBinaryWatch(turnedOn int) []string {
	res := make([]string, 0)
	backtrack(0, 0, 0, turnedOn, &res)
	return res
}

func main() {
	fmt.Println(readBinaryWatch(1))
	fmt.Println(readBinaryWatch(9))
	fmt.Println(readBinaryWatch(2))
}
