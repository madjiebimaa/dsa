package stack

import "fmt"

var NoSuchElementException = fmt.Errorf("No such element")

func New() *Stack {
	return &Stack{Arr: []int{}}
}

type Stack struct {
	Arr []int
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s Stack) Size() int {
	return len(s.Arr)
}

func (s *Stack) Push(val int) error {
	s.Arr = append(s.Arr, val)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, NoSuchElementException
	}

	lastIdx := s.Size() - 1
	val := s.getLastElement()
	s.Arr = s.Arr[:lastIdx]
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
	return s.Arr[len(s.Arr)-1]
}
