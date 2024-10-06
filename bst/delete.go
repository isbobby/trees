package bst

func (root *Node) HardDelete(val int) (*Node, error) {
	delPath, err := root.pathFromRoot(val)
	if err != nil {
		return nil, err
	}
	var decrementSizeOnPath = func(nodes []*Node) {
		for _, node := range nodes {
			node.decrementSize()
		}
	}
	decrementSizeOnPath(delPath)

	toDelete := delPath[len(delPath)-1]
	if toDelete.isLeaf() {
		if toDelete.parent == nil {
			return nil, nil
		} else {
			if toDelete.val > toDelete.parent.val {
				toDelete.parent.right = nil
			} else {
				toDelete.parent.left = nil
			}
			toDelete.parent = nil
		}
	}

	newTreeRoot := toDelete.getRoot()
	if toDelete.right != nil && toDelete.left != nil {
		successorPath := toDelete.popSuccessor()
		decrementSizeOnPath(successorPath)
		successor := successorPath[len(successorPath)-1]

		if successor.val > successor.parent.val {
			successor.parent.right = successor.right
		} else {
			successor.parent.left = successor.right
		}
		toDelete.val = successor.val
		successor.parent = nil
		return toDelete.getRoot(), nil
	} else if toDelete.right != nil {
		if toDelete.parent != nil {
			if toDelete.val > toDelete.parent.val {
				toDelete.parent.right = toDelete.right
			} else {
				toDelete.parent.left = toDelete.right
			}
		} else {
			newTreeRoot = toDelete.right
		}
	} else {
		if toDelete.parent != nil {

			if toDelete.val > toDelete.parent.val {
				toDelete.parent.right = toDelete.left
			} else {
				toDelete.parent.left = toDelete.left
			}
		} else {
			newTreeRoot = toDelete.right
		}
	}
	toDelete.parent = nil
	delete(nodeSize, toDelete)
	return newTreeRoot, nil
}

func (root *Node) popSuccessor() []*Node {
	path := []*Node{}
	root = root.right
	path = append(path, root)
	for root.left != nil {
		root = root.left
		path = append(path, root)
	}

	return path
}

func (root *Node) SoftDelete(val int) error {
	node, err := root.Find(val)
	if err != nil {
		return err
	}
	node.deleted = true
	return nil
}
