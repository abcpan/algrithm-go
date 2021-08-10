package findRestaurant

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	sumMemo := make(map[int] []string)
	min := math.MaxInt32
	memo := make(map[string]int)
	for i := 0; i < len(list1);i++ {
		memo[list1[i]] = i
	}

	for i:= 0; i < len(list2);i++ {
		cur := list2[i]
		if idx, ok := memo[cur];ok {
			 sum := i + idx
			 sumMemo[sum] = append(sumMemo[sum], cur)
			 if min > sum {
			 	 min = sum
			 }
		}
	}
	return sumMemo[min]
}