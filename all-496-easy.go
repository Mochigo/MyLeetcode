package main

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	index := make(map[int]int)
	for i, num2 := range nums2 {
		index[num2] = i + 1
	}

	for i, num1 := range nums1 {
		for j := index[num1]; j < n; j++ {
			if nums2[j] > num1 {
				ans[i] = nums2[j]
				break
			}
		}
	}

	return ans
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	index := make(map[int]int)
	for i, num2 := range nums2 {
		index[num2] = i + 1
	}

	for i, num1 := range nums1 {
		for j := index[num1]; j < n; j++ {
			if nums2[j] > num1 {
				ans[i] = nums2[j]
				break
			}
		}
	}

	return ans
}
