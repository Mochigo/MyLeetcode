package main

// 10 由质因数2和5构成
// 尾随0的个数由2和5的较小值决定

func trailingZeroes(n int) int {
	// 一般数学法，计算每个可能提供质因数5的数有的质因数5的个数
	ans := 0
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			ans++
		}
	}

	return ans
}

// 优化计算
func trailingZeroes1(n int) int {
	// 举例方便理解，第一次 /= 5， 找出所有有一个质因数5的数（本质是找出所有（5^1）的倍数）
	// 第二次 /= 5, 找出所有有因数25的数，本来应该累计+2，但是由于第一个循环已经将这部分的加过了，所以只用加一遍就行了
	// ……

	ans := 0
	for n > 0 {
		n /= 5
		ans += n
	}

	return ans
}
