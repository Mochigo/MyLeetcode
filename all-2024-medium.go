package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(helper(answerKey, 'T', k), helper(answerKey, 'F', k))
}

func helper(key string, char byte, k int) int {
	ans := 0
	l := 0
	cnt := 0  // 记录此时的长度
	left := 0 // 记录其他字符的数量，当其他字符数量超过k时，需要弹出第一个其他字符
	for r := range key {
		if key[r] != char {
			left++
		}
		cnt++

		for left > k {
			if key[l] != char {
				left--
			}
			cnt--
			l++
		}

		ans = max(ans, cnt)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
