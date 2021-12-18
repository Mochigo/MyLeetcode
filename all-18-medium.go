package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	n := len(nums)
	if n < 4 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		// 最小的四个数相加超过 target,直接退出,没有必要再看了
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		if i > 0 && nums[i-1] == nums[i] || nums[i]+nums[n-1] < target-nums[n-2]-nums[n-3] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// 一个数固定 最小的三个数相加超过 target,直接退出,没有必要再看了
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if j > i+1 && nums[j-1] == nums[j] || nums[i]+nums[j] < target-nums[n-1]-nums[n-2] {
				continue
			}
			for l, r := j+1, n-1; l < r; {
				if nums[l]+nums[r] < target-nums[i]-nums[j] {
					l++
				} else if nums[l]+nums[r] > target-nums[i]-nums[j] {
					r--
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				}
				// 去掉使用重复的数字
				for l < r && l > j+1 && nums[l] == nums[l-1] {
					l++
				}
				// 去掉使用重复的数字
				for l < r && r < n-1 && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}

	return ans
}
