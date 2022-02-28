package main

import "math/bits"

func maximumRequests(n int, requests [][]int) int {
	// 请求和楼的数量都比较少，那么可以遍历回溯了
	// 对于某个请求，保留进不进行
	m := len(requests)
	all := (1 << m) - 1
	ans := 0

next:
	for mask := all; mask >= 0; mask-- {
		cnt := bits.OnesCount(uint(mask))
		// 即这种情况不需要考虑，因为就算能够成功也不如已有的方案
		if cnt <= ans {
			continue
		}

		building := make([]int, n)
		for i, req := range requests {
			if (mask>>i)&1 == 1 {
				from, to := req[0], req[1]
				building[from]--
				building[to]++
			}
		}

		for _, v := range building {
			if v != 0 {
				continue next
			}
		}

		ans = cnt
	}

	return ans
}
