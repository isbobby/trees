package bst

import "errors"

type Node struct {
	val int

	deleted bool
	parent  *Node
	left    *Node
	right   *Node
}

type NodeInfo struct {
	Val     int
	Deleted bool
	Parent  *Node
	Left    *Node
	Right   *Node
}

var (
	ErrNodeAlreadyExists = errors.New("failed to insert node, node already in sub-tree")
	ErrNodeNotFound      = errors.New("failed to find node in sub-tree")
	ErrRootIsNil         = errors.New("cannot perform operation on nil root")
)

type BinarySearchTree interface {
	Size() int
	Insert(val int) (*Node, error)
	Find(Val int) (*Node, error)
	SoftDelete(val int) error
	HardDelete(val int) (*Node, error)

	Left() (*Node, error)
	Right() (*Node, error)
	Val() (int, error)
	PrintSubtreeLevelOrder()
	PrintNodeInfo()
}

func (root *Node) Val() (int, error) {
	if root == nil {
		return 0, ErrRootIsNil
	}
	return root.val, nil
}

func (root *Node) Left() (*Node, error) {
	if root == nil {
		return nil, ErrRootIsNil
	}
	return root.left, nil
}

func (root *Node) Right() (*Node, error) {
	if root == nil {
		return nil, ErrRootIsNil
	}
	return root.right, nil
}
