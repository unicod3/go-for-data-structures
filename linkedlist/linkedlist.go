package linkedlist

import "fmt"

//import "log"

type Node struct {
	next *Node
	data interface{}
}

type LinkedList struct {
	head *Node
}

func NewNode(val interface{}) *Node {
	return &Node{
		next: nil,
		data: val,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) GetLength() int {
	length := 0

	if l.head == nil {
		return length
	}

	last := l.head
	if last != nil {
		length++
	}
	for last.next != nil {
		last = last.next
		length++
	}
	return length
}

func (l *LinkedList) Prepend(d interface{}) {

	node := NewNode(d)
	if l.head == nil {
		l.head = node
		return
	}

	node.next = l.head
	l.head = node
}

func (l *LinkedList) Append(d interface{}) {

	node := NewNode(d)
	if l.head == nil {
		l.head = node
		return
	}

	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = node
}

func (l *LinkedList) Pop() interface{} {
	elem := l.head
	l.head = l.head.next
	return elem.data
}

func (l *LinkedList) Display() {
	fmt.Print("List: [")
	if l.head == nil {
		fmt.Print("]\n")
		return
	}
	node := l.head
	for node.next != nil {
		fmt.Print(node.data)
		fmt.Print(" ---> ")
		node = node.next
	}

	fmt.Print(node.data)
	fmt.Print("]\n")
}
