package main

func maxProduct1(words []string) int {
	n := len(words)
	// 如果只判断字母是否出现，不需要在意位置和个数的话，除了使用[26]bool，还可以使用二进制数表示
	mask := make([]int, len(words))
	for i, word := range words {
		for _, c := range word {
			mask[i] |= 1 << (c - 'a') // 这里使用 ^ 和 | 没有区别
		}

	}

	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if mask[i]&mask[j] == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}

	return ans
}

func maxProduct(words []string) int {
	// 如果只判断字母是否出现，不需要在意位置和个数的话，除了使用[26]bool，还可以使用二进制数表示
	dict := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, c := range word {
			mask |= 1 << (c - 'a') // 这里使用 ^ 和 | 没有区别
		}

		if dict[mask] < len(word) {
			dict[mask] = len(word)
		}
	}

	ans := 0
	for x, l1 := range dict {
		for y, l2 := range dict {
			if x&y == 0 && l1*l2 > ans {
				ans = l1 * l2
			}
		}
	}

	return ans
}
