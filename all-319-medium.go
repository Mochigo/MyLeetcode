package main

import "math"

func bulbSwitch(n int) int {
	// bulbs := make([]bool, n+1)

	// for round := 1; round <= n; round++ {
	//     for i := round; i <= n; i += round {
	//         bulbs[i] = !bulbs[i]
	//     }
	// }

	// ans := 0
	// for i := 1; i <= n; i++ {
	//     if bulbs[i] {
	//         ans++
	//     }
	// }

	// return ans

	return int(math.Sqrt(float64(n)))
}
