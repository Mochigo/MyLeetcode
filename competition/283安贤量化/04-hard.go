package main

func replaceNonCoprimes(nums []int) []int {
	s := []int{nums[0]}
	for _, num := range nums[1:] {
		s = append(s, num)
		for len(s) > 1 {
			x, y := s[len(s)-2], s[len(s)-1]
			g := gcd(x, y)
			if g == 1 {
				break
			}
			s = s[:len(s)-1]
			s[len(s)-1] = s[len(s)-1] * y / g
		}
	}

	return s
}

func gcd(x, y int) int {
	for x != 0 {
		x, y = y%x, x
	}

	return y
}
