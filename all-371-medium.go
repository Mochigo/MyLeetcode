package main

import "fmt"

// 两个感悟
// 一个是如果一个变量会在每个循环中被重新赋值，例如这里的c，更好的处理方法声明一个作用域更大的变量，而不是在循环中短声明一个变量，因为每个循环中，这个变量都会被释放然后重新分配内存，会增加内存消耗
// 一个是a ^ b可以得到 a和b 无进位情况下的相加结果， (a & b) << 1可以得到 进位情况， 然后在循环中把进位情况相加，直到不再产生进位以后，就可以得到相加的结果了
func getSum(a int, b int) int {
	var c int
	for b != 0 {
		c = (a & b) << 1
		a ^= b
		b = c
	}

	return a
}

func main() {
	fmt.Println(getSum(2, 3))
}
