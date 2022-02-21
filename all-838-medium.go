package main

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	nds := []byte(dominoes)

	lastIdx := -1
	lastO := byte('L')
	for i, d := range nds {
		if d == '.' {
			continue
		}
		if d == 'L' {
			if lastO == 'L' {
				for j := i - 1; j > lastIdx; j-- {
					nds[j] = 'L'
				}
			} else if lastO == 'R' {
				for j, k := i-1, lastIdx+1; k < j; j, k = j-1, k+1 {
					nds[j] = 'L'
					nds[k] = 'R'
				}
			}
		} else if d == 'R' {
			if lastO == 'R' {
				for j := i - 1; j > lastIdx; j-- {
					nds[j] = 'R'
				}
			}
		}

		lastO = d
		lastIdx = i
	}

	if lastO == 'R' {
		for i := lastIdx + 1; i < n; i++ {
			nds[i] = 'R'
		}
	}

	return string(nds)
}
