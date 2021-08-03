package numSquares


func numSquares(n int) int {
	queue := []int{0}
	visited := map[int]int {0: 0}
	level := 0

	for len(queue) != 0 {
		size := len(queue)
		level++
		// 循环一层
		for i := 0; i < size;i++ {
			digit := queue[0]
			queue = queue[1:]
			for j := 1; j <= n; j++ {
				value := digit + j * j
				if value == n {
					return level
				}
				if value > n {
					break
				}
				if _, isIn := visited[value]; !isIn {
					queue = append(queue, value)
					visited[value] = 0
				}
			}
		}
	}
	return level
}