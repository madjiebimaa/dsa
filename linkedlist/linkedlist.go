package linkedlist

import "fmt"

type Node struct {
	key  interface{}
	next *Node
}

func (n Node) Key() interface{} {
	return n.key
}

func (n Node) Next() *Node {
	return n.next
}

type LinkedList struct {
	head *Node
}

func New(k interface{}) *LinkedList {
	return &LinkedList{head: &Node{key: k}}
}

func (l LinkedList) Head() *Node {
	return l.head
}

func (l LinkedList) Tail() *Node {
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	return curr
}

func (l LinkedList) Display() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%+v -> ", curr.key)
		curr = curr.next
	}

	fmt.Println(nil)
}

func (l LinkedList) CountNodes() int {
	curr := l.head
	count := 1
	for curr.next != nil {
		curr = curr.next
		count++
	}

	return count
}

func (l *LinkedList) InsertForward(k interface{}) {
	node := &Node{
		key:  k,
		next: nil}

	front := l.head
	l.head = node
	l.head.next = front
}

func (l *LinkedList) InsertBackward(k interface{}) {
	node := &Node{
		key:  k,
		next: nil}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = node
}

func (l *LinkedList) DeletionForward() {
	l.head = l.head.next
}

func (l *LinkedList) DeletionBackward() {
	curr := l.head
	for curr.next.next != nil {
		curr = curr.next
	}

	curr.next = nil
}

func (l *LinkedList) Delete(k interface{}) {
	slow := l.head
	fast := l.head.next

	for fast != nil {
		if fast.key == k {
			slow.next = fast.next
			fast = slow.next
		}

		slow = slow.next
		fast = fast.next
	}
}

func (l *LinkedList) Reverse() {
	curr := l.head
	var slow *Node

	for curr != nil {
		fast := curr.next
		curr.next = slow
		slow = curr
		curr = fast
	}

	l.head = slow
}
