package linkedlist

import (
	"fmt"
)

type NodeS struct {
	Data int
	Next *NodeS
}

func NewNodeS(data int) *NodeS {
	return &NodeS{
		Data: data,
		Next: nil,
	}
}

type SingleLinkedList struct {
	Head   *NodeS
	Length int
}

func NewSingleLinkedListWithoutHead() *SingleLinkedList {
	return &SingleLinkedList{
		Head:   nil,
		Length: 0,
	}
}

func NewSingleLinkedListWithHead(head *NodeS) *SingleLinkedList {
	return &SingleLinkedList{
		Head:   head,
		Length: 1,
	}
}

func (l *SingleLinkedList) Tail() *NodeS {
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	return curr
}

func (l *SingleLinkedList) InsertForward(n int) {
	node := NewNodeS(n)
	if l.Head != nil {
		node.Next = l.Head
	}

	l.Head = node
	l.Length += 1
}

func (l *SingleLinkedList) InsertBackward(n int) {
	curr := l.Head
	if curr == nil {
		l.InsertForward(n)
	} else {
		node := NewNodeS(n)
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
		l.Length += 1
	}
}

func (l *SingleLinkedList) InsertAt(idx int, n int) {
	curr := l.Head
	if curr == nil || idx == 1 {
		l.InsertForward(n)
	} else {
		if idx > l.Length {
			l.InsertBackward(n)
		} else {
			node := NewNodeS(n)
			for i := 2; i < idx; i++ {
				curr = curr.Next
			}

			next := curr.Next
			node.Next = next
			curr.Next = node
			l.Length += 1
		}
	}
}

func (l *SingleLinkedList) DeleteForward() {
	curr := l.Head
	l.Head = curr.Next
	curr = nil
	l.Length -= 1
}

func (l *SingleLinkedList) DeleteBackward() {
	curr := l.Head
	for curr != nil {
		curr = curr.Next
	}

	curr = nil
	l.Length -= 1
}

func (l *SingleLinkedList) DeleteAt(idx int) {
	curr := l.Head
	if idx > l.Length {
		idx = l.Length
	}

	if idx == 1 {
		l.DeleteForward()
	} else if idx > 1 {
		for i := 2; i < idx; i++ {
			curr = curr.Next
		}

		next := curr.Next
		curr.Next = next.Next
		next = nil
		l.Length -= 1
	}
}

func (l *SingleLinkedList) Tranverse() {
	curr := l.Head
	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
}

func (l *SingleLinkedList) Reverse() {
	var (
		prev *NodeS
		curr *NodeS
		next *NodeS
	)

	prev = nil
	curr = l.Head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l.Head = prev
}
