package main

import (
	"strconv"
	"unicode"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	if s[0] != '[' {
		ni := &NestedInteger{}
		num, _ := strconv.Atoi(s)
		ni.SetInteger(num)
		return ni
	}

	stack, num, negative := []*NestedInteger{}, 0, false
	for i, c := range s {
		if c == '[' {
			stack = append(stack, &NestedInteger{})
		} else if c == '-' {
			negative = true
		} else if unicode.IsDigit(c) {
			num = num*10 + int(c-'0')
		} else if c == ',' || c == ']' {
			if unicode.IsDigit(rune(s[i-1])) {
				if negative {
					num = -num
				}
				ni := NestedInteger{}
				ni.SetInteger(num)
				stack[len(stack)-1].Add(ni)
			}
			num, negative = 0, false
			if c == ']' && len(stack) > 1 {
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1].Add(*tmp)
			}
		}
	}

	return stack[len(stack)-1]
}
