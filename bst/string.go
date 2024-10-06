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

func (n NodeInfo) String() string {
	v := []byte(fmt.Sprint(n.Val))
	p, l, r, del := []byte("_"), []byte("_"), []byte("_"), []byte("X")
	if n.Parent != nil {
		if n.Parent.deleted {
			p = []byte(del)
		} else {
			p = []byte(fmt.Sprint(n.Parent.val))
		}
	}
	if n.Left != nil {
		if n.Left.deleted {
			l = []byte(del)
		} else {
			l = []byte(fmt.Sprint(n.Left.val))
		}
	}
	if n.Right != nil {
		if n.Right.deleted {
			r = []byte(del)
		} else {
			r = []byte(fmt.Sprint(n.Right.val))
		}
	}

	return fmt.Sprintf("  %v  \n  |\n  %v\n / \\\n%v   %v\n",
		string(p), string(v), string(l), string(r))
}
