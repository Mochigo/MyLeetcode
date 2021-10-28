package main

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := (r-l)/2 + l
		if nums[m] == target {
			return m
		}

		if nums[m] >= nums[0] {
			if target >= nums[0] && nums[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target <= nums[n-1] && nums[m] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}

	}

	return -1
}
