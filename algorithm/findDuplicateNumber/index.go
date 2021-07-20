package findDuplicateNumber

func findDuplicateNumber(nums []int) int {
	var slow, fast = 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if fast == slow {
			break
		}
	}

	// 相遇后进行处理
	var finder = 0
	for {
		finder = nums[finder]
		slow = nums[slow]
		if slow == finder {
			break
		}
	}
	return slow
}
