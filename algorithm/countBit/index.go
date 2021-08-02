package algorithm



func countOne(n int) int {
	var count = 0
	for i := 0; i < 32; i++ {
		x := n >> i
		count += x & 1
		if x == 0 {
			break
		}
	}
	return count
}


func countBit(n int) []int {
	var ret = make([]int, n + 1)
	for idx := range ret {
		ret[idx] = countOne(idx)
	}
	return ret
}

// 不断利用前面已经计算完成的结果
func countBitByDynamic(n int) [] int {
	var memo = make([]int, n + 1)
	memo[0] = 0
	for i := 0; i <=n; i++ {
		// 如果时奇数, 1 的个数等于前面偶数个数 + 1
		if n & 1 == 1 {
			memo[i] = memo[i - 1] + 1
		} else {
			// 如果等于偶数, 1 的个数
			memo[i] =  memo[i >> 1]
		}
	}
	return memo
}