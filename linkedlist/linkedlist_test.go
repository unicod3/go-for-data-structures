package linkedlist

import (
	"testing"
)

func TestLinkedListPrepend(t *testing.T) {
	want := []string{
		"foo",
		"bar",
		"baz",
	}
	list := NewLinkedList()
	for i := 0; i < 3; i++ {
		list.Prepend(want[i])
	}

	got := list.GetLength()
	if got != len(want) {
		t.Errorf("got: %d  want: %d", got, len(want))
	}
	list.Display()
}

func TestLinkedListAppend(t *testing.T) {
	want := []string{
		"foo",
		"bar",
		"baz",
	}
	list := NewLinkedList()
	for i := 0; i < 3; i++ {
		list.Append(want[i])
	}

	got := list.GetLength()
	if got != len(want) {
		t.Errorf("got: %d  want: %d", got, len(want))
	}
	list.Display()
}

func TestLinkedListPop(t *testing.T) {
	want := []string{
		"foo",
		"bar",
		"baz",
	}
	list := NewLinkedList()
	for i := 0; i < 3; i++ {
		list.Append(want[i])
	}

	for i := 0; i < 3; i++ {
		got := list.Pop()
		if got != want[i] {
			t.Errorf("got: %s  want: %s", got, want[i])
		}
		// check remaining element count
		remaining := (2 - i)
		if list.GetLength() != remaining {
			t.Errorf("got: %d want: %d", list.GetLength(), remaining)
		}
	}
	list.Display()
}

func TestLinkedListInsertAt(t *testing.T) {
	want := []string{
		"foo",
		"bar",
		"baz",
	}
	list := NewLinkedList()
	for i := 0; i < 3; i++ {
		list.Append(want[i])
	}

	val := "inserted"
	list.InsertAt(2, val)
	got := list.GetLength()
	if got != 4 {
		t.Errorf("got: %d  want: %d", got, 4)
	}
	list.Display()
}

func TestLinkedListGetAt(t *testing.T) {
	want := []string{
		"foo",
		"bar",
		"baz",
	}
	list := NewLinkedList()
	for i := 0; i < 3; i++ {
		list.Append(want[i])
	}

	got := list.GetAt(2)
	if got != want[2] {
		t.Errorf("got: %s  want: %s", got, want[2])
	}
	list.Display()
}
