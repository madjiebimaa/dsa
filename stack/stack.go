package stack

import "fmt"

var NoSuchElementException = fmt.Errorf("No such element")

func New() *Stack {
	return &Stack{arr: []int{}}
}

type Stack struct {
	arr []int
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Push(val int) error {
	s.arr = append(s.arr, val)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, NoSuchElementException
	}

	lastIdx := s.Size() - 1
	val := s.getLastElement()
	s.arr = s.arr[:lastIdx]
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, NoSuchElementException
	}

	val := s.getLastElement()
	return val, nil
}

func (s *Stack) getLastElement() int {
	return s.arr[len(s.arr)-1]
}
