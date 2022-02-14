package linkedlist_test

import (
	"testing"

	"github.com/madjiebimaa/dsa/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestInsertNilHead(t *testing.T) {

	t.Run("forward nil head", func(t *testing.T) {
		l := linkedlist.NewSingleLinkedListWithoutHead()

		l.InsertForward(2)
		assert.Equal(t, 2, l.Head.Data)
		assert.Equal(t, 2, l.Tail().Data)
	})

	t.Run("backward nil head", func(t *testing.T) {
		l := linkedlist.NewSingleLinkedListWithoutHead()

		l.InsertBackward(3)
		assert.Equal(t, 3, l.Head.Data)
		assert.Equal(t, 3, l.Tail().Data)
	})

	t.Run("at position nil head", func(t *testing.T) {
		l := linkedlist.NewSingleLinkedListWithoutHead()

		l.InsertAt(2, 4)
		assert.Equal(t, 4, l.Head.Data)
		assert.Equal(t, 4, l.Tail().Data)
	})
}

func TestInsertNotNilHead(t *testing.T) {
	t.Run("initial node", func(t *testing.T) {
		node := linkedlist.NewNodeS(1)
		l := linkedlist.NewSingleLinkedListWithHead(node)

		assert.Equal(t, 1, l.Head.Data)
		assert.Equal(t, 1, l.Tail().Data)
	})

	t.Run("insert forward change head to that value", func(t *testing.T) {
		node := linkedlist.NewNodeS(1)
		l := linkedlist.NewSingleLinkedListWithHead(node)

		l.InsertForward(2)
		assert.Equal(t, 2, l.Head.Data)
		assert.Equal(t, 1, l.Tail().Data)
	})

	t.Run("insert backward change tail to that value", func(t *testing.T) {
		node := linkedlist.NewNodeS(1)
		l := linkedlist.NewSingleLinkedListWithHead(node)

		l.InsertBackward(3)
		assert.Equal(t, 1, l.Head.Data)
		assert.Equal(t, 3, l.Tail().Data)
	})

	t.Run("insert at first position is equal to insert forward", func(t *testing.T) {
		node := linkedlist.NewNodeS(1)
		l := linkedlist.NewSingleLinkedListWithHead(node)

		l.InsertAt(1, 100)
		assert.Equal(t, 100, l.Head.Data)
		assert.Equal(t, 1, l.Tail().Data)
	})

	t.Run("insert at position that more than length of linkedlist is equal to insert backward", func(t *testing.T) {
		node := linkedlist.NewNodeS(1)
		l := linkedlist.NewSingleLinkedListWithHead(node)

		l.InsertAt(100, 10)
		assert.Equal(t, 1, l.Head.Data)
		assert.Equal(t, 10, l.Tail().Data)
	})

	t.Run("insert at specific position will change the node into that value and the before node will be in the next of this node", func(t *testing.T) {
		node := linkedlist.NewNodeS(1)
		l := linkedlist.NewSingleLinkedListWithHead(node)
		l.InsertBackward(2)
		l.InsertBackward(3)

		n := 2
		l.InsertAt(n, 99)
		curr := l.Head
		for i := 1; i < n; i++ {
			curr = curr.Next
		}
		assert.Equal(t, 99, curr.Data)
	})
}
