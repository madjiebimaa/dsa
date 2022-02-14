package linkedlist

type NodeD struct {
	Data int
	Prev *NodeD
	Next *NodeD
}

func NewNodeD(data int) *NodeD {
	return &NodeD{
		Data: data,
		Prev: nil,
		Next: nil,
	}
}

type DoubleLinkedList struct {
	Head   *NodeD
	Length int
}

func NewDoubleLinkedListWithoutHead() *DoubleLinkedList {
	return &DoubleLinkedList{
		Head:   nil,
		Length: 0,
	}
}

func NewDoubleLinkedListWithHead(head *NodeD) *DoubleLinkedList {
	return &DoubleLinkedList{
		Head:   head,
		Length: 0,
	}
}

func (l *DoubleLinkedList) Tail() *NodeD {
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	return curr
}

func (l *DoubleLinkedList) InsertForward(n int) {
	node := NewNodeD(n)
	if l.Head != nil {
		node.Next = l.Head
	}
	l.Head = node
	l.Length += 1
}
