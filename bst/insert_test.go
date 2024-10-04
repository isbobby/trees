package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsnsert(t *testing.T) {
	testRoot := New(3)

	leftVal := 1
	rightVal := 5
	testRoot.Insert(leftVal)
	testRoot.Insert(rightVal)

	left, err := testRoot.Left()
	assert.Nil(t, err)
	gotLeftVal, _ := left.Val()
	assert.Equal(t, leftVal, gotLeftVal)

	right, err := testRoot.Right()
	assert.Nil(t, err)
	gotRightVal, _ := right.Val()
	assert.Equal(t, rightVal, gotRightVal)
}

func TestInsertWithDuplicates(t *testing.T) {
	elements := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}

	root := New(2)
	for _, element := range elements {
		root.Insert(element)
	}

	left, _ := root.Left()
	right, _ := root.Right()

	leftVal, _ := left.Val()
	rightVal, _ := right.Val()

	assert.Equal(t, 1, leftVal)
	assert.Equal(t, 3, rightVal)
}
