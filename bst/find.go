package bst

func (root *Node) Find(val int) (*Node, error) {
	if root == nil {
		return nil, ErrNodeNotFound
	}

	if val == root.val {
		if root.deleted {
			return nil, ErrNodeNotFound
		} else {
			return root, nil
		}
	}

	if val > root.val {
		return root.right.Find(val)
	} else {
		return root.left.Find(val)
	}
}
