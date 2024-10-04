package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWithSoftDelete(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6, 8}
	testTreeRoot := shuffleArrayAndCreateBST(elements)
	deleteElement := 3
	testTreeRoot.SoftDelete(deleteElement)

	for i := 0; i < len(elements); i++ {
		got, gotErr := testTreeRoot.Find(elements[i])
		if gotErr != nil {
			assert.Equal(t, deleteElement, elements[i])
			assert.Equal(t, ErrNodeNotFound, gotErr)
		} else {
			assert.Equal(t, got.val, elements[i])
			assert.Nil(t, gotErr)
		}
	}
}
