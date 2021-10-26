package main

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := []int{} // 从形象上来看，是为了维护一个从栈底到栈顶递减的栈
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	for i := 0; i < 2*n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			// 压出元素，知道出现栈为空，或者 栈顶元素要大于当前元素
			// 因为能被压出，说明栈顶元素的下标所指的值的下一个更大的值肯定是当前元素下标
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return ans
}
