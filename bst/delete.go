package bst

import "fmt"

func (root *Node) HardDelete(val int) (*Node, error) {
	toDelete, err := root.binarySearchFromRoot(val)
	if err != nil {
		return nil, err
	}
	if toDelete == nil {
		return nil, ErrRootIsNil
	}
	fmt.Println("deleting", toDelete)
	var newSubtreeNode *Node
	if toDelete.right != nil {
		newSubtreeNode = toDelete.promoteRight()
	} else {
		newSubtreeNode = toDelete.promoteLeft()
	}

	if toDelete.parent != nil {
		if toDelete.val > toDelete.parent.val {
			toDelete.parent.right = newSubtreeNode
		} else {
			toDelete.parent.left = newSubtreeNode
		}
		toDelete.parent = nil
	}
	delete(nodeSize, toDelete)
	fmt.Println("GOT newSubTreeNode", newSubtreeNode)
	newSubtreeNode.PrintNodeInfo()
	fmt.Println(nodeSize)
	return newSubtreeNode.getRoot(), nil
}

func (root *Node) promoteRight() *Node {
	fmt.Println("Promote right child", root.right)
	currLeft := root.left
	newRoot := root.right

	newRootLeft := root.right.left
	newRootLeftMostLeaf := newRootLeft
	leftPath := []*Node{newRootLeftMostLeaf}
	for newRootLeftMostLeaf.left != nil {
		newRootLeftMostLeaf = newRootLeft.left
		leftPath = append(leftPath, newRootLeftMostLeaf)
	}

	newRootLeftMostLeaf.left = currLeft
	currLeft.parent = newRootLeftMostLeaf
	newRoot.parent = nil
	fmt.Println(leftPath)
	for len(leftPath) > 0 {
		node := leftPath[len(leftPath)-1]
		leftPath = leftPath[:len(leftPath)-1]
		fmt.Println("increment size for", node)
		node.incrementSizeBy(currLeft.Size())
	}
	newRoot.incrementSizeBy(currLeft.Size())
	fmt.Println("Return promoted", newRoot)

	return newRoot
}

func (root *Node) promoteLeft() *Node {
	fmt.Println("Promote left child", root.left)
	return nil
}

func (root *Node) SoftDelete(val int) error {
	node, err := root.Find(val)
	if err != nil {
		return err
	}
	node.deleted = true
	return nil
}
