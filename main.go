package main

import (
	"fmt"
	"my-trees/bst"
)

func main() {
	myTree := bst.NewFromArray([]int{4, 7, 1, 6, 8, 2, 3, -2})
	myTree.PrintSubtreeLevelOrder()

	found, _ := myTree.Find(7)
	found.PrintSubtreeLevelOrder()
	found.PrintNodeInfo()

	myTree.SoftDelete(-2)
	myTree.PrintSubtreeLevelOrder()

	myTree.Insert(9)
	_, err := myTree.Insert(9)
	fmt.Println(err)
	myTree.PrintSubtreeLevelOrder()

	worseTree := bst.NewFromArray([]int{1, 2, 3, 4, 5})
	worseTree.PrintSubtreeLevelOrder()

}
