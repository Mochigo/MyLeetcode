package main

func maxDepth(s string) int {
	max := 0
	stack := []byte{}
	d := 0
	for _, c := range s {
		if c == '(' {
			d++
			stack = append(stack, '(')
			if d > max {
				max = d
			}
		} else if c == ')' {
			d--
			stack = stack[:len(stack)-1]
		}
	}

	return max
}
