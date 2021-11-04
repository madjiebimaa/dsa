package stack_test

import (
	"testing"

	"github.com/madjiebimaa/go-ds/stack"
	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := stack.New()
	t.Run(
		"A stack is empty on construction", func(t *testing.T) {
			assert.Equal(t, true, s.IsEmpty())
		})

	t.Run(
		"A stack has size 0 on construction", func(t *testing.T) {
			assert.Equal(t, 0, s.Size())
		})
}

func TestInser(t *testing.T) {
	t.Run("After n pushes to an empty stack, n > 0, the stack is not empty and its size is n", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		assert.Equal(t, 4, s.Size())
		assert.Equal(t, false, s.IsEmpty())
	})

	t.Run("If one pushes x then pops, the value popped is x.", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(13)
		assert.Equal(t, 3, s.Size())
		val, _ := s.Pop()
		assert.Equal(t, 13, val)
		assert.Equal(t, 2, s.Size())
	})

	t.Run("If one pushes x then peeks, the value returned is x, but the size stays the same", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(13)
		val, _ := s.Peek()
		assert.Equal(t, 13, val)

	})

	t.Run("If the size is n, then after n pops, the stack is empty and has a size 0", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(13)
		assert.Equal(t, 3, s.Size())
		s.Pop()
		s.Pop()
		s.Pop()
		assert.Equal(t, true, s.IsEmpty())
		assert.Equal(t, 0, s.Size())
	})
}

func TestError(t *testing.T) {
	t.Run("Popping from an empty stack does throw a NoSuchElementException", func(t *testing.T) {
		s := stack.New()
		_, err := s.Pop()
		if err == nil {
			t.Fail()
			t.Logf("Expect error is not nil, but got `%v`", err)
		}
		assert.Equal(t, stack.NoSuchElementException, err)
	})

	t.Run("Peeking into an empty stack does throw a NoSuchElementException", func(t *testing.T) {
		s := stack.New()
		_, err := s.Peek()
		if err == nil {
			t.Fail()
			t.Logf("Expect error is not nil, but got `%v`", err)
		}
		assert.Equal(t, stack.NoSuchElementException, err)
	})
}

//
//
