package bst

func New(val int) BinarySearchTree {
	root := &Node{val: val}
	nodeSize = map[*Node]int{}
	root.incrementSize()
	return root
}

func NewFromArray(vals []int) BinarySearchTree {
	if len(vals) == 0 {
		return nil
	}

	root := New(vals[0])

	for i := 1; i < len(vals); i++ {
		root.Insert(vals[i])
	}

	return root
}
