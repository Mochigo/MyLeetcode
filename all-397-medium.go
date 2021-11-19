package main

func integerReplacement(n int) int {
	queue := []int{n}
	ans := 0
	for len(queue) > 0 {
		tmp := []int{}
		for _, num := range queue {
			if num == 1 {
				return ans
			}
			if num%2 == 1 {
				tmp = append(tmp, num+1)
				tmp = append(tmp, num-1)
			} else {
				tmp = append(tmp, num/2)
			}
		}
		queue = tmp
		ans++
	}
	return -1
}
