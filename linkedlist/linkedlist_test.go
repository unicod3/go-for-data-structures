package linkedlist

import (
	"testing"
)

func TestPrepend(t *testing.T) {
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

func TestAppend(t *testing.T) {
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

func TestPop(t *testing.T) {
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
