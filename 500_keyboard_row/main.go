package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	mapper := make([]map[byte]bool, len(rows))
	for i, row := range rows {
		mapper[i] = make(map[byte]bool)
		for _, c := range row {
			mapper[i][byte(c)] = true
		}
	}

	var results []string

	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		lower := strings.ToLower(word)
		first := lower[0]
		row := -1

		for i, set := range mapper {
			if set[first] {
				row = i
				break
			}
		}
		if row == -1 {
			continue
		}
		sameRow := true
		for _, c := range lower {
			if !mapper[row][byte(c)] {
				sameRow = false
				break
			}
		}
		if sameRow {
			results = append(results, word)
		}
	}

	return results
}

func main() {
	fmt.Println(findWords([]string{"hello", "Alaska", "Dad", "Peace"}))
}
