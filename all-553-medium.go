package main

import (
	"strconv"
	"strings"
)

// 最大的可能除法构造，就是要使除数尽可能的少，尽可能得小，这里分母连除的构造，可以使分母只有一个nums[1]
// 而其他数均为分子
func optimalDivision(nums []int) string {
	ret := &strings.Builder{}
	n := len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	} else if n == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	ret.WriteString(strconv.Itoa(nums[0]))
	ret.WriteString("/(")
	for i := 1; i < n; i++ {
		ret.WriteString(strconv.Itoa(nums[i]))
		if i < n-1 {
			ret.WriteByte('/')
		}
	}
	ret.WriteByte(')')

	return ret.String()
}
