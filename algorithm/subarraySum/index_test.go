package subarraySum


func subarraySum(nums []int, k int) int {
	count := 0
	memo := map[int]int {
		0: 1,
	}
	preSum := 0
	for i := 0; i < len(nums);i++ {
		preSum += nums[i]
		// preSum[j] - preSum[i] == k
		key := preSum - k
		if memo[key] > 0 {
			count += memo[preSum - k]
		}
		memo[preSum]++
	}

	return count
}