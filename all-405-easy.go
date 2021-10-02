package main

import (
	"fmt"
	"strconv"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	if num < 0 {
		num += 1 << 32
	}
	ans := ""
	var digit int
	for num > 0 {
		digit = num % 16
		if 0 <= digit && digit <= 9 {
			ans = strconv.Itoa(digit) + ans
		} else {
			ans = string(digit-10+'a') + ans
		}
		num /= 16
	}

	return ans
}

func main() {
	fmt.Println(toHex(1000))
}
