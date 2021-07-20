package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	var result = make([][]int, 0)
	var size = len(nums)
	var left int
	var right int
	var sum  int
	// 将数组排序
	sort.Ints(nums)
	for i := 0; i < size; i++ {
		// 如果和前一位数字相同, 则代表结果重复
		if i != 0 && nums[i] == nums[i - 1] {
			continue
		}
		// 如果当前数 大于0, 则代表 其余两数 必须小于之和必须小于0, 但是经过排序, 这是不存在的
		if nums[i] > 0 {
			return result
		}

		left = i + 1
		right = size - 1
		for left < right {
			sum = nums[left] + nums[right] + nums[i]
			if sum == 0 {
				result = append(result, []int {nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left - 1] {
					left++
				}
				for left < right && nums[right] == nums[right + 1] {
					right--
				}
			}
			if sum < 0 {
				left ++
			}else {
				right--
			}
		}
	}
	return result
}