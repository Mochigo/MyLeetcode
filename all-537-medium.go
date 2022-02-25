package main

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	idx1, idx2 := strings.IndexByte(num1, '+'), strings.IndexByte(num2, '+')
	r1, _ := strconv.Atoi(num1[:idx1])
	r2, _ := strconv.Atoi(num2[:idx2])
	i1, _ := strconv.Atoi(num1[idx1+1 : len(num1)-1])
	i2, _ := strconv.Atoi(num2[idx2+1 : len(num2)-1])

	a := r1 * r2
	b := r1*i2 + r2*i1
	c := -1 * i1 * i2

	return fmt.Sprintf("%d+%di", a+c, b)
}
