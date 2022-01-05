package main

func modifyString(s string) string {
	res := []byte(s)
	n := len(s)
	for i, c := range res {
		if c == '?' {
			for b := byte('a'); b <= 'c'; b++ {
				// 不满足，如果存在左右两侧，那么左右两侧与相同的情况
				if !(i > 0 && res[i-1] == b || i < n-1 && res[i+1] == b) {
					res[i] = b
					break
				}
			}
		}
	}
	return string(res)
}
