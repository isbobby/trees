package bst

func New(val int) BinarySearchTree {
	root := &Node{val: val}
	nodeSize = map[*Node]int{}
	root.incrementSize()
	return root
}
