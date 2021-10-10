package main

import "fmt"

func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}

	ans := 2
	for (1+ans)*ans/2 <= n {
		ans++
	}

	return ans - 1
}

/*
func arrangeCoins(n int) int {
    return int((math.Sqrt(float64(8*n+1)) - 1) / 2)
}
*/

/*
func arrangeCoins(n int) int {
    return sort.Search(n, func(a int) bool {
        a++
        return a*(a+1) > 2 * n
    })
}
*/

func main() {
	fmt.Println(arrangeCoins(100))
}
