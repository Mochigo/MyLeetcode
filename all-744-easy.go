package main

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	n := sort.Search(len(letters), func(a int) bool {
		return letters[a] > target
	})

	if n == len(letters) {
		return letters[0]
	}
	return letters[n]
}
