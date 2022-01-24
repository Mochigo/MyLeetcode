package main

func numberOfMatches1(n int) int {
	m := 0
	for n > 1 {
		m += n / 2
		n = n/2 + n%2
	}

	return m
}

func numberOfMatches2(n int) int {
	return n - 1
}
