package leastInterval

import (
	"math"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	var buckets = make([]int, 26)
	// 统计每个字符的个数
	for _, task := range tasks {
		idx := task - byte('A')
		buckets[idx]++
	}

	// 进行排序
	sort.Ints(buckets)
	// 最多数量的字符
	maxTimes := buckets[25]
	maxCount := 1
	for i := 25; i >= 0;i-- {
		if buckets[i] == buckets[i - 1] {
			maxCount ++
		}else {
			break
		}
	}
	ret := (maxTimes - 1) * (n + 1) + maxCount
	return int(math.Max(float64(ret), float64(len(tasks))))
}