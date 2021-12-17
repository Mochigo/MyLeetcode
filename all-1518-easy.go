package main

func numWaterBottles(numBottles int, numExchange int) int {
	empty := numBottles
	ans := numBottles
	for empty >= numExchange {
		exchanged := empty / numExchange
		ans += exchanged
		empty = empty%numExchange + exchanged
	}

	return ans
}
