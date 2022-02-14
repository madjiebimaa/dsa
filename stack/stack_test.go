package stack_test

import (
	"testing"

	"github.com/madjiebimaa/dsa/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack_New(t *testing.T) {
	t.Run("should return nil items when initialized stack", func(t *testing.T) {
		s := stack.New()
		assert.Equal(t, []int(nil), s.Items)
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("should add 1 item to the stack when push 1", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		assert.Equal(t, 1, s.Peek())
	})

	t.Run("should add 2 items to the stack when push 1 and 2", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		assert.Equal(t, 2, s.Peek())
	})
}

func TestStack_IsEmpty(t *testing.T) {
	t.Run("should return true when initialized stack", func(t *testing.T) {
		s := stack.New()
		assert.Equal(t, true, s.IsEmpty())
	})

	t.Run("should return false when stack have 1 item", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		assert.Equal(t, false, s.IsEmpty())
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("should return 1 when popping stack that have 1 item", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		assert.Equal(t, 1, s.Pop())
	})

	t.Run("should return 2 when popping stack that have 2 items", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		assert.Equal(t, 2, s.Pop())
	})

	t.Run("should return -1 when popping empty stack", func(t *testing.T) {
		s := stack.New()
		assert.Equal(t, -1, s.Pop())
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("should return 1 when peeking stack that have 1 item", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		assert.Equal(t, 1, s.Peek())
	})

	t.Run("should return 2 when peeking stack that have 2 items", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		assert.Equal(t, 2, s.Peek())
	})

	t.Run("should return -1 when peeking empty stack", func(t *testing.T) {
		s := stack.New()
		assert.Equal(t, -1, s.Peek())
	})
}

func TestStack_Len(t *testing.T) {
	t.Run("should return 0 when stack is empty", func(t *testing.T) {
		s := stack.New()
		assert.Equal(t, 0, s.Len())
	})

	t.Run("should return 1 when stack have 1 item", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		assert.Equal(t, 1, s.Len())
	})

	t.Run("should return 2 when stack have 2 items", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		assert.Equal(t, 2, s.Len())
	})

	t.Run("should return 1 when popping stack that have 2 items", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Pop()
		assert.Equal(t, 1, s.Len())
	})
}
