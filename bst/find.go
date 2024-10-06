package bst

func (root *Node) Find(val int) (*Node, error) {
	found, err := root.binarySearchFromRoot(val)
	if err != nil {
		return found, err
	}

	if found.deleted {
		return nil, ErrNodeNotFound
	}

	return found, err
}

func (root *Node) binarySearchFromRoot(val int) (*Node, error) {
	treeRoot := root.getRoot()
	return treeRoot.binarySearch(val)
}

func (root *Node) getRoot() *Node {
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func (root *Node) binarySearch(val int) (*Node, error) {
	if root == nil {
		return nil, ErrNodeNotFound
	}

	if val == root.val {
		return root, nil
	}

	if val > root.val {
		return root.right.binarySearch(val)
	} else {
		return root.left.binarySearch(val)
	}
}
