package bst

func (root *Node) Insert(val int) (*Node, error) {
	if root == nil {
		return nil, ErrRootIsNil
	}

	if val == root.val {
		if root.deleted {
			root.deleted = false
			root.incrementSize()
			return root, nil
		} else {
			return nil, ErrNodeAlreadyExists
		}
	}

	newNode := Node{val: val}
	if val > root.val {
		if root.right == nil {
			root.right = &newNode
			root.right.incrementSize()
			root.incrementSize()
			return root.right, nil
		} else {
			newNode, err := root.right.Insert(val)
			if err == nil {
				root.incrementSize()
			}
			return newNode, err
		}
	} else {
		if root.left == nil {
			root.left = &newNode
			root.left.incrementSize()
			root.incrementSize()
			return root.left, nil
		} else {
			newNode, err := root.left.Insert(val)
			if err == nil {
				root.incrementSize()
			}
			return newNode, err
		}
	}
}
