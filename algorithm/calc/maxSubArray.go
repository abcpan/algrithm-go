package calc

import "math"

func maxSubArray(nums []int) int {
	var ret = nums[0]
	var sum = 0
	for _, num := range nums {
		if sum > 0 {
			sum +=num
		}else {
			sum = num
		}
		ret = int(math.Max(float64(sum), float64(ret)))
	}
	return ret
}
