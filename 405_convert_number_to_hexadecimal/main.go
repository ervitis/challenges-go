package main

import (
	"fmt"
)

var hxDigits = "0123456789abcdef"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	var hex string

	if num < 0 {
		hex = twosComplement(num)
	} else {
		for num > 0 {
			digit := num % 16
			hex = string(hxDigits[digit]) + hex
			num /= 16
		}
	}

	return hex
}

func twosComplement(num int) string {
	num = 0xFFFFFFFF + num + 1

	var hex string

	for num > 0 {
		digit := num % 16
		hex = string(hxDigits[digit]) + hex
		num /= 16
	}

	for len(hex) < 8 {
		hex = "0" + hex
	}

	return hex
}

func main() {
	fmt.Println(toHex(26))
	fmt.Println(toHex(-1))
}
