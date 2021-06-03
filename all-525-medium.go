package main

import (
	"fmt"
)

func findMaxLength(nums []int) int {
    // key | value 
    // k     index  表示此时sum中 1 多k个
    // -k    index  表示此时sum中 0 多k个
    // 0     index  表示此时sum中 0和1数量一样多
    mp := map[int]int{0: -1}
    sum := 0 // 存放前缀和
    res := 0
    for i, num := range nums {
        sum += num
        n := i+1 // 每一个位都是1的最大情况，n-sum表示0的数量
        numOfOne := sum
        numOfZero := n - sum
        if preIndex, ok := mp[numOfOne - numOfZero]; !ok {
            mp[numOfOne - numOfZero] = i
        } else {
            res = max(res, i - preIndex)
        }
    }
    return res
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}

func main() {
	a := []int{1, 0, 0, 1, 1}
	fmt.Println(findMaxLength(a))
}