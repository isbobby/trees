package bst

import "fmt"

func (root *Node) String() string {
	if root == nil {
		return "_"
	}
	if root.deleted {
		return "X"
	}
	return fmt.Sprint(root.val)
}
