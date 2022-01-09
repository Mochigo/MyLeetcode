package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	n := len(releaseTimes)
	max, maxIndex := releaseTimes[0], 0
	for i := 1; i < n; i++ {
		d := releaseTimes[i] - releaseTimes[i-1]
		if d > max || d == max && keysPressed[i] > keysPressed[maxIndex] {
			maxIndex = i
			max = d
		}
	}

	return byte(keysPressed[maxIndex])
}
