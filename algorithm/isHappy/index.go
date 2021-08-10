package isHappy

import "math"

func isHappy(n int) bool {
	memo := make(map[int]bool)
	for n != 1 {
		if _, isIn := memo[n];isIn {
			return false
		}else {
			memo[n] = true
		}
		n = calNumSum(n)
	}
	return true
}


// calNumSum 计算一个数的各个数字平方和
func calNumSum(n int) int {
	var sum float64 = 0
	for n != 0 {
		sum += math.Pow(float64(n % 10), 2)
		n /= 10
	}
	return int(sum)
}