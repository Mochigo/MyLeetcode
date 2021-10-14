package main

import (
	"fmt"
	"sort"
)

// 调用了sort包中的Search函数，他会返回满足传入func最小的下标值
// 这里利用了二分法实际上是看左右两侧的规律性的一个函数，只要有一定的规律，就可以使用二分法
// 当arr[i] > arr[i+1]时，说明此时在峰顶的右侧或者就是峰顶
// 当arr[i] < arr[i+1]时，说明此时在峰顶的左侧
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr)-1, func(x int) bool {
		return arr[x] > arr[x+1]
	})
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	fmt.Println(peakIndexInMountainArray(nums))
}
