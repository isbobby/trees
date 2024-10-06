package bst

func (root *Node) isLeaf() bool {
	if root == nil {
		return false
	}
	return root.left == nil && root.right == nil
}
