package tree

import "abcpan.net.cn/algorithm"

func inorderTravel(root *algorithm.TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	left := inorderTravel(root.Left)
	if left != nil {
		ret = append(ret, left...)
	}
	ret = append(ret, root.Val)
	right := inorderTravel(root.Right)
	if right != nil {
		ret = append(ret, right...)
	}
	return ret
}