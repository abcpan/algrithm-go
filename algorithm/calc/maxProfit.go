package calc

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var maxProfit float64 = 0
	var minPrice = prices[0]
	for i := 1; i < len(prices); i++ {
		// 更新最大利润
		var cur = prices[i]
		maxProfit = math.Max(float64(cur - minPrice), maxProfit)
		// 更新最高价格
		minPrice = int(math.Min(float64(cur), float64(minPrice)))
	}

	return int(maxProfit)
}