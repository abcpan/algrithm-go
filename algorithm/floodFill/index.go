package floodFill

// Directions 方向
var Directions = [][]int {
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	queue := make([][]int, 0)
	queue = append(queue, []int{sr, sc})
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	image[sr][sc] = newColor
	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		for _, direction := range Directions {
			r := point[0] + direction[0]
			c := point[1] + direction[1]
			if inArea(image, r, c) && image[r][c] == oldColor {
				// 着色
				image[r][c] = newColor
				queue = append(queue, []int{r, c})
			}
		}
	}
	return image
}


func inArea(image [][]int, row, col int) bool {
	rows := len(image)
	cols := len(image[0])
	return row >= 0 && row < rows && col >= 0 && col < cols
}
