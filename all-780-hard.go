package main

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	// 第一个问题：从起点开始推，使用dfs会造成大量分支，
	// 解决办法： 自终点至起点分析，却只有一条路径，相对较大的值减去相对较小的值
	// 对于(x,y)，只可能来自(x, y-x), (x-y, y)

	// 第二个问题，形如 1999996 和 3，即使能确定相减运算的方向，也需要执行多次，
	// 解决办法： 使用取模运算，可以快速得到多次减运算的最终结果，这适用于需要多次减运算直到两者大小关系发生变化，可以在辗转相除法中看见这个方法的影子
	nx, ny := tx, ty
	// 保证最小化
	for nx > sx && ny > sy && nx != ny {
		if nx > ny {
			nx = nx % ny // nx = nx - ny 只是单次操作, 使用取模，可以把这个这种单次操作起到变成减除直到不能再减的做哦那个
		} else {
			ny = ny % nx
		}
	}

	// 从上述循环中出来，只有可能，nx == sx || ny == sy
	switch {
	case nx == sx && ny == sy:
		return true
	case nx == sx:
		return ny > sy && (ny-sy)%sx == 0 // 这里注意的是，(ny-sy) % sx == 0 并不等价于 ny % sx = sy， 如 2，1，1，1（nx，ny，sx，sy）
	case ny == sy:
		return nx > sx && (nx-sx)%sy == 0
	default:
		return false
	}
}
