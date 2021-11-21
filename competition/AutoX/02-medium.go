package main

func wateringPlants(plants []int, capacity int) int {
	has := capacity
	steps := 0
	pos := -1 // 起始位置
	for i, need := range plants {
		// 如果水是足够的
		if has >= need {
			has -= need
			steps += i - pos // 表示到达这个位置需要的步数
			pos = i          // 表示走到这个格子，并浇水
		} else {
			has = capacity //重新装满水
			has -= need
			steps += 2*pos + 3 // pos+1表示回到河边，pos+2表示走到下一个浇水
			pos = i
		}
	}

	return steps
}
