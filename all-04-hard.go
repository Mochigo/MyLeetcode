package main

import (
	"fmt"
	"sort"
)

// 两种方法类似，即简单粗暴地构建了合并后的数组，根据奇偶性返回数据
// 前者自己构建，后者使用了切片合并，再调用官方sort包
/*
执行用时：16 ms, 在所有 Go 提交中击败了76%的用户
内存消耗：5.6 MB, 在所有 Go 提交中击败了40%的用户
*/
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	sumLength := m + n
	num := make([]int, sumLength)
	index := 0
	i, j := 0, 0
	for i < m && j < n {
		// 按大小
		if nums1[i] <= nums2[j] {
			num[index] = nums1[i]
			index++
			i++
		} else {
			num[index] = nums2[j]
			index++
			j++
		}
	}

	for i < m {
		num[index] = nums1[i]
		i++
		index++
	}

	for j < n {
		num[index] = nums2[j]
		j++
		index++
	}

	// fmt.Println(num, len(num))
	if sumLength%2 == 1 {
		return float64(num[sumLength/2])
	} else {
		return float64(num[sumLength/2-1]+num[sumLength/2]) / 2.0
	}

	return 0
}

/*
执行用时：20 ms, 在所有 Go 提交中击败了49.39%的用户
内存消耗：6.1 MB, 在所有 Go 提交中击败了27.64%的用户
*/
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	num := append(nums1, nums2...)
	sort.Ints(num)
	sumLength := len(nums1) + len(nums2)
	if sumLength%2 == 1 {
		return float64(num[sumLength/2])
	} else {
		return float64(num[sumLength/2-1]+num[sumLength/2]) / 2.0
	}

	return 0
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays2(nums1, nums2))
}
