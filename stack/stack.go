package stack

type Stack struct {
	Items []int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	length := s.Len() - 1
	poppedItem := s.Peek()
	s.Items = s.Items[:length]
	return poppedItem
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	return s.Items[s.Len()-1]
}

func (s *Stack) Len() int {
	return len(s.Items)
}
