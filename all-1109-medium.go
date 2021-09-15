package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	// 差分数组，dif 通过对dif的加减，可以快速的对一段区间内的值进行同时的加减
	// 例如想要对 bookings 中 [i, j] 同时加上 a, 就可以对dif[i] += a, dif[j+1] -= a

	dif := make([]int, n+1) // 由于预定作为的原因，此时每个航班的预先安排都为 0， 这里不用特别地进行处理
	// 如果dif长度设置为 n 的话，则需要用一个新的变量base 来记录生成数组的第一个元素的。
	for _, book := range bookings {
		first, last, seats := book[0]-1, book[1], book[2]
		dif[first] += seats
		dif[last] -= seats
	}

	ans := make([]int, n)
	ans[0] = dif[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + dif[i]
	}

	return ans
}

func main() {
	bookings := [][]int{[]int{1, 2, 10}, []int{2, 3, 20}, []int{2, 5, 25}}
	n := 5
	fmt.Println(corpFlightBookings(bookings, n))
}
