package main

func isRectangleCover1(rectangles [][]int) bool {
	// 通过一次遍历，来确定左下角和右上角
	// 有两个问题：怎么判断是否发生了重叠，以及，直到了左右顶角，怎么判断，它是正好包裹住了整个大矩形
	// 现有的想法是，先防止出现了重叠问题，再用面积来判断是否刚好构成了矩形

	// n := len(rectangles)
	type point struct{ x, y int }
	cnt := make(map[point]int)

	sum := 0
	minX, minY, maxX, maxY := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	for _, rectangle := range rectangles {
		x, y, a, b := rectangle[0], rectangle[1], rectangle[2], rectangle[3]
		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, a)
		maxY = max(maxY, b)

		sum += (b - y) * (a - x)

		cnt[point{x, y}]++
		cnt[point{a, y}]++
		cnt[point{x, b}]++
		cnt[point{a, b}]++
	}
	area := (maxY - minY) * (maxX - minX)

	if area != sum || cnt[point{minX, minY}] != 1 || cnt[point{maxX, maxY}] != 1 || cnt[point{maxX, minY}] != 1 || cnt[point{minX, maxY}] != 1 {
		return false
	}

	delete(cnt, point{minX, minY})
	delete(cnt, point{maxX, maxY})
	delete(cnt, point{minX, maxY})
	delete(cnt, point{maxX, minY})

	for _, c := range cnt {
		if c != 2 && c != 4 {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
