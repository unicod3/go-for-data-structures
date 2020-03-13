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

func (l *LinkedList) GetAt(k int) interface{} {

	if k < 0 || k > l.GetLength() {
		fmt.Printf("index is out of bounds, specify between: %d, %d", 0, l.GetLength())
		return nil
	}

	counter := 0
	elem := l.head
	for elem.next != nil {
		if k == counter {
			return elem.data
		}
		elem = elem.next
		counter++
	}
	return elem.data
}

func (l *LinkedList) InsertAt(k int, d interface{}) {

	if k < 0 || k > l.GetLength() {
		fmt.Printf("index is out of bounds, specify between: %d, %d", 0, l.GetLength())
		return
	}

	node := NewNode(d)
	if l.head == nil {
		l.head = node
		return
	}

	if k == 0 {
		l.Prepend(d)
		return
	}

	if k == l.GetLength() {
		l.Append(d)
		return
	}

	counter := 1
	elem := l.head
	for elem.next != nil {

		if k == counter {
			tmp := elem.next
			elem.next = node
			elem.next.next = tmp
			break
		}
		elem = elem.next
		counter++
	}
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
