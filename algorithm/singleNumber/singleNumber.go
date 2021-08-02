package algorithm

func singleNumber(nums []int) int {
	var memo = make(map[int] int)
	for _, num := range nums {
		memo[num]++
	}
	for k, v := range memo {
		if v == 1 {
			return k
		}
	}
	return 0
}


func singleNumberByXor(nums []int) int {
	var single = 0
	for _, num := range nums {
		single ^= num
	}
	return single
}