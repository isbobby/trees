package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	root := shuffleArrayAndCreateBST([]int{0, 1, 2, 3, 4, 5, 6})
	assert.Equal(t, 7, root.Size())
}
