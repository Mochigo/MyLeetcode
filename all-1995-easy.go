package main

func countQuadruplets1(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				for m := k + 1; m < n; m++ {
					if nums[i]+nums[j]+nums[k] == nums[m] {
						ans++
					}
				}
			}
		}
	}

	return ans
}

func countQuadruplets2(nums []int) (ans int) {
	cnt := map[int]int{}
	n := len(nums)
	for c := n - 2; c >= 2; c-- {
		cnt[nums[c+1]]++
		// 枚举a
		for a, x := range nums[:c] {
			for _, y := range nums[a+1 : c] {
				if sum := x + y + nums[c]; cnt[sum] > 0 {
					ans += cnt[sum]
				}
			}
		}
	}
	return ans
}
