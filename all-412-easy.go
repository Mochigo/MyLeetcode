package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ans[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			ans[i-1] = "Buzz"
		} else if i%3 == 0 {
			ans[i-1] = "Fizz"
		} else {
			ans[i-1] = strconv.Itoa(i)
		}
	}

	return ans
}

func main() {
	fmt.Println(fizzBuzz(15))
}
