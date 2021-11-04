package main

func isPerfectSquare1(num int) bool {
	// [1， 17)
	l, r := 1, num+1
	for l < r {
		m := (r-l)/2 + l
		sqr := m * m
		if sqr == num {
			return true
		} else if sqr > num {
			r = m
		} else {
			l = m + 1
		}
	}

	return false
}

func isPerfectSquare2(num int) bool {
	i := 1
	for i*i < num {
		i++
	}
	return i*i == num
}
