package main

func trap1(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxLeft[0] = height[0]
	maxRight[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	ans := 0
	for i := 1; i < n; i++ {
		ans += min(maxRight[i], maxLeft[i]) - height[i]
	}

	return ans
}

func trap2(height []int) int {
	stack := []int{}
	ans := 0
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			ans += (i - left - 1) * (min(height[left], h) - height[top])
		}
		stack = append(stack, i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
