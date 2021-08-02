package algorithm


func majorityElement(nums []int) int {
	var target int
	var count = 0
	for _, num := range nums {
		if count == 0 {
			target = num
		}
		if target == num {
			count++
		}else {
			count--
		}
	}
	return target
}