package stacks_test

import (
	"fmt"
	"testing"

	"github.com/madjiebimaa/dsa/stacks"
)

func TestStack(t *testing.T) {
	s := stacks.NewStacks()
	fmt.Println(s.IsEmpty())
	s.Push(2)
	s.Push(3)
	s.Push(1)
	s.Push(100)
	fmt.Println(s.Items)
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
}

type Stack struct {
	Items []byte
}

func (s *Stack) Peek() byte {
	return s.Items[len(s.Items)-1]
}

func (s *Stack) Pop() byte {
	toRemove := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return toRemove
}

func (s *Stack) Push(w byte) {
	s.Items = append(s.Items, w)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}
