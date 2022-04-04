package main

type NumArray struct {
	trees []int
	nums  []int
	n     int
}

// 返回最低位的二进制数
func lowBit(x int) int {
	return x & (-x)
}

func Constructor(nums []int) NumArray {
	trees := make([]int, len(nums)+1)
	n := len(nums)
	na := NumArray{trees, nums, n}
	for i, num := range nums {
		na.Add(i+1, num)
	}

	return na
}

func (na *NumArray) Add(index int, val int) {
	for ; index <= na.n; index += lowBit(index) {
		na.trees[index] += val
	}
}

func (na *NumArray) Update(index int, val int) {
	v := na.nums[index]
	na.Add(index+1, val-v)
	na.nums[index] = val
}

func (na *NumArray) Prefix(index int) (sum int) {
	for index > 0 {
		sum += na.trees[index]
		index -= lowBit(index)
	}
	return
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.Prefix(right+1) - na.Prefix(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
