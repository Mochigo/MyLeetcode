package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	num1 := int(math.Sqrt(float64(area)))
	for area%num1 > 0 {
		num1--
	}

	return []int{area / num1, num1}
}

func main() {
	fmt.Println(constructRectangle(99))
}
