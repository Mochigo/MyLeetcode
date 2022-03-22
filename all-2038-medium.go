package main

// 统计三个相连在一起的谁多
func winnerOfGame(colors string) bool {
	var cnt [2]int
	n := len(colors)

	for i := 1; i < n-1; i++ {
		if colors[i] == colors[i-1] && colors[i] == colors[i+1] {
			cnt[colors[i]-'A']++
		}
	}

	return cnt[0] > cnt[1]
}
