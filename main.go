package main

import (
	"my-trees/bst"
)

func main() {
	myTree := bst.NewFromArray([]int{4, 7, 1, 6, 8, 2, 3})
	myTree.PrintSubtreeLevelOrder()
	myTree.SoftDelete(6)
	myTree.PrintSubtreeLevelOrder()

	myTree.Insert(11)
	myTree.Insert(-2)
	myTree.Insert(9)
	myTree.PrintSubtreeLevelOrder()
}
