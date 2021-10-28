package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	// min := nums[0]
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}
