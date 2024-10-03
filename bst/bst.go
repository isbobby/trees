package bst

import "errors"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var (
	ErrNodeAlreadyExists = errors.New("failed to insert node, node already in sub-tree")
	ErrNodeNotFound      = errors.New("failed to find node in sub-tree")
)

type BinarySearchTree interface {
	Size() int
	Insert(val int) (*Node, error)
	Find(Val int) (*Node, error)
	Delete(val int) error
}

func New() BinarySearchTree {
	return &Node{}
}
