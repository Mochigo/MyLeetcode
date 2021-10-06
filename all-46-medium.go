package main

import "fmt"

func permute(nums []int) [][]int {
	solutions := [][]int{}

	n := len(nums)
	used := make([]bool, n)
	path := []int{}

	var dfs func(int)
	dfs = func(index int) {
		if index == n {
			solutions = append(solutions, append([]int{}, path...))
			return
		}

		for i, num := range nums {
			if !used[i] {
				path = append(path, num)
				used[i] = true
				dfs(index + 1)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return solutions
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
