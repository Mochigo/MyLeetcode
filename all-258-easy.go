package main

func addDigits(num int) int {
	for num/10 != 0 {
		tmp := 0
		for num != 0 {
			tmp += num % 10
			num /= 10
		}
		num = tmp
	}

	return num
}
