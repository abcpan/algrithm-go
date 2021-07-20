package algorithm

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	var want = 1
	root := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	ret := DiameterOfBinaryTree(root)
	if ret != want {
		t.Fatalf("error")
	}
}