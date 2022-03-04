package main

func subArrayRanges(nums []int) int64 {
	var ans int64
	n := len(nums)
	for i := range nums {
		min, max := nums[i], nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] < min {
				min = nums[j]
			}
			if nums[j] > max {
				max = nums[j]
			}

			ans += int64(max - min)
		}
	}

	return ans
}
