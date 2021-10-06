package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	solutions := [][]int{}
	used := make([]bool, n)
	path := []int{}

	var dfs func(int)
	dfs = func(index int) {
		if index == n {
			solutions = append(solutions, append([]int{}, path...))
			return
		}

		for i, num := range nums {
			if used[i] || (i-1 >= 0 && !used[i-1] && nums[i] == nums[i-1]) {
				continue
			}

			path = append(path, num)
			used[i] = true
			dfs(index + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return solutions
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
