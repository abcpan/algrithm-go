package tree

func check(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && check(root1.Left, root2.Right) && check(root1.Right, root2.Left)
}


func isSymmetric(root *Node) bool {
 return check(root, root)
}