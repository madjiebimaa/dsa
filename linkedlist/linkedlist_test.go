package linkedlist_test

import (
	"testing"

	"github.com/madjiebimaa/go-ds/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	l := linkedlist.New(2)
	t.Run("A Linkedlist has one value construction", func(t *testing.T) {
		assert.Equal(t, 2, l.Head().Key())
	})

	t.Run("A Linkedlist has size 1 on construction", func(t *testing.T) {
		assert.Equal(t, 1, l.CountNodes())
	})
}

func TestInsert(t *testing.T) {
	t.Run("After insert some value in forward way, value that inserted will become new head", func(t *testing.T) {
		l := linkedlist.New(2)
		l.InsertForward(3)
		assert.Equal(t, 3, l.Head().Key())
		l.InsertForward(5)
		assert.Equal(t, 5, l.Head().Key())
	})

	t.Run("After insert some value in backward way, value that inserted will become new tail", func(t *testing.T) {
		l := linkedlist.New(2)
		l.InsertBackward(3)
		assert.Equal(t, 3, l.Tail().Key())
		l.InsertBackward(5)
		assert.Equal(t, 5, l.Tail().Key())
	})
}
