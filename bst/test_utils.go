package bst

import "math/rand"

func shuffleArrayAndCreateBST(arr []int) BinarySearchTree {
	if len(arr) == 0 {
		return nil
	}

	// shuffle in place
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)

		arr[i], arr[j] = arr[j], arr[i]
	}

	root := New(arr[0])

	for i := 1; i < len(arr); i++ {
		root.Insert(arr[i])
	}

	return root
}
