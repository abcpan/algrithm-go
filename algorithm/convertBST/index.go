package convertBST

import "abcpan.net.cn/algorithm/tree"

type BSTSolution struct {
	sum int
}

func (this *BSTSolution) solve(root *tree.Node) *tree.Node {
	solution := &BSTSolution{0}
	solution.dfs(root)
	return root
}

func (this *BSTSolution) dfs(root *tree.Node) {
	if root == nil {
		return
	}
	if root.Right != nil {
		this.dfs(root.Right)
	}
	root.Val = this.sum + root.Val
	if root.Left != nil {
		this.dfs(root.Left)
	}
}


