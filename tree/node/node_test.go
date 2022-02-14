package node_test

import (
	"testing"

	"github.com/madjiebimaa/dsa/tree/node"
	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	t.Run("should return 1 value when initialized node", func(t *testing.T) {
		n := node.New(1)
		assert.Equal(t, n.Value, 1)
	})

	t.Run("should return nil value in left node when initialized node", func(t *testing.T) {
		n := node.New(1)
		assert.Equal(t, n.Left, (*node.Node)(nil))
	})

	t.Run("should return nil value in right node when initialized node", func(t *testing.T) {
		n := node.New(1)
		assert.Equal(t, n.Right, (*node.Node)(nil))
	})
}
