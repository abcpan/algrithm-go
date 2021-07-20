package maxarea

import "math"

func maxArea(nums []int) int {
	var startIdx, endIdx = 0, len(nums) - 1
	var max float64 = 0
	// 直到撞针结束
  for startIdx != endIdx {
  	var left = nums[startIdx]
  	var right = nums[endIdx]
  	area := math.Min(float64(left), float64(right)) * float64(endIdx - startIdx)
		max = math.Max(max, area)
		if left < right {
			startIdx++
		}else {
			endIdx--
		}
	}
	return int(max)
}