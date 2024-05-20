package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var res string

	carry := 0
	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 || carry > 0 {
		d1, d2 := 0, 0
		if i >= 0 {
			d1 = int(num1[i] - '0') //convert u8 into digit
			i--
		}
		if j >= 0 {
			d2 = int(num2[j] - '0')
			j--
		}

		sum := d1 + d2 + carry
		res = strconv.Itoa(sum%10) + res
		carry = sum / 10
	}
	return res
}

func main() {
	fmt.Println(addStrings("11", "123"))
}
