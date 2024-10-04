package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6, 8}
	testTreeRoot := shuffleArrayAndCreateBST(elements)

	for i := 0; i < len(elements); i++ {
		got, gotErr := testTreeRoot.Find(elements[i])
		assert.Equal(t, got.val, elements[i])
		assert.Nil(t, gotErr)
	}
}
