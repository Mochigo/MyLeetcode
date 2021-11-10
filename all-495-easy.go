package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	expired := 0
	ans := 0
	for _, t := range timeSeries {
		if expired < t {
			ans += duration
		} else {
			ans += t + duration - expired
		}
		expired = t + duration
	}

	return ans
}
