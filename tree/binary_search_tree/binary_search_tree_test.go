package binarysearchtree_test

import (
	"testing"

	binarySearchTree "github.com/madjiebimaa/dsa/tree/binary_search_tree"
	"github.com/madjiebimaa/dsa/tree/node"
	"github.com/stretchr/testify/assert"
)

func TestBST_New(t *testing.T) {
	t.Run("should return nil when initialized bst", func(t *testing.T) {
		bst := binarySearchTree.New()
		assert.Equal(t, (*node.Node)(nil), bst.Root)
	})
}

func TestBST_Insert(t *testing.T) {
	t.Run("should return 1 root value when insert bst with 1", func(t *testing.T) {
		bst := binarySearchTree.New()
		bst.Insert(1, bst.Root)
		assert.Equal(t, 1, bst.Root.Value)
	})

	t.Run("should add value to left node when adding value smaller than root", func(t *testing.T) {
		bst := binarySearchTree.New()
		bst.Insert(2, bst.Root)
		bst.Insert(1, bst.Root)
		assert.Equal(t, 1, bst.Root.Left.Value)
	})

	t.Run("should add value to right node when adding bigger than root", func(t *testing.T) {
		bst := binarySearchTree.New()
		bst.Insert(1, bst.Root)
		bst.Insert(2, bst.Root)
		assert.Equal(t, 2, bst.Root.Right.Value)
	})

	t.Run("should add value to left grandchild node when adding smaller than ancestor left node", func(t *testing.T) {
		bst := binarySearchTree.New()
		bst.Insert(3, bst.Root)
		bst.Insert(2, bst.Root)
		bst.Insert(1, bst.Root)
		assert.Equal(t, 1, bst.Root.Left.Left.Value)
	})

	t.Run("should add value to right grandchild node when adding greater than ancestor right node", func(t *testing.T) {
		bst := binarySearchTree.New()
		bst.Insert(1, bst.Root)
		bst.Insert(2, bst.Root)
		bst.Insert(3, bst.Root)
		assert.Equal(t, 3, bst.Root.Right.Right.Value)
	})
}
