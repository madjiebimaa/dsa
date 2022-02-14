package stacks

import (
	"errors"
)

type Stacks struct {
	Items []int
}

func NewStacks() *Stacks {
	return &Stacks{}
}

func (s *Stacks) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stacks) Peek() int {
	return s.Items[len(s.Items)-1]
}

func (s *Stacks) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("can't pop because stacks is empty")
	}

	toRemove := s.Peek()
	s.Items = s.Items[:len(s.Items)-1]
	return toRemove, nil
}

func (s *Stacks) Push(n int) {
	s.Items = append(s.Items, n)
}
