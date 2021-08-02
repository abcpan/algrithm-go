package algorithm

import "math"

type Solution struct {
	MaxCount float64
}

func (solution *Solution) Depth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	var L = solution.Depth(root.Left)
	var R = solution.Depth(root.Right)
	// 最大节点数量
	solution.MaxCount = math.Max(solution.MaxCount, L+R+1)
	// 求出深度
	return math.Max(L, R) + 1
}

func DiameterOfBinaryTree(root *TreeNode) int {
	solution := &Solution{
		MaxCount: 0,
	}
	solution.Depth(root)
	return int(solution.MaxCount - 1)
}
