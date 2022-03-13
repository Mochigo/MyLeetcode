package main

// 先求前导1，0对应1字节，2对应2字节，3对应3字节，4对应4字节，超过非法，1也非法
func validUtf8(data []int) bool {
	rule := map[int]int{
		0: 1,
		2: 2,
		3: 3,
		4: 4,
	}

	n := len(data)
	i := 0
	cur, skip := 0, 0
	for i < n {
		cur = data[i]
		skip = bits(cur)
		if bits(cur) > 4 || bits(cur) == 1 {
			return false
		}
		j, end := i+1, i+rule[skip]
		for j < end && j < n {
			if bits(data[j]) != 1 {
				return false
			}
			j++
		}

		i += rule[skip]
	}

	return i == n
}

// 返回前导1的个数
func bits(num int) int {
	i := 7
	for i >= 0 && ((num&(1<<i))>>i) == 1 {
		i--
	}

	return 7 - i
}
