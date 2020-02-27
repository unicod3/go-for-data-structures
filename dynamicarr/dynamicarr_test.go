package dynamicarr

import (
	"testing"
)

func TestGetLength(t *testing.T) {
	want := 1
	darr := NewDynamicArray()
	darr.Push(want)
	got := darr.GetLength()
	if got != want {
		t.Errorf("got: %d  want: %d", got, want)
	}
}

func TestGetCapacity(t *testing.T) {
	want := 2
	darr := NewDynamicArray()
	darr.Push("something")
	darr.Push("something 2")
	got := darr.GetCapacity()
	if got != want {
		t.Errorf("got: %d  want: %d", got, want)
	}
}

func TestSet(t *testing.T) {
	darr := NewDynamicArray()
	darr.Push("something")
	darr.Push("something 2")
	for i := -100; i < 100; i = i + 10 {
		if i < 0 {
			err := darr.Set(i, "value")
			if err == nil {
				t.Errorf("got: nil  want: %s", err)
			}
			continue
		}

		err := darr.Set(i, "value")
		if err != nil {
			t.Errorf("got: %s  want: nil", err)
		}

	}
}

func TestGrowSize(t *testing.T) {
	want := 1
	darr := NewDynamicArray()
	for i := 0; i < 100; i++ {
		darr.Push("something 2")
		if i == want {
			want *= 2
		}
		if darr.GetCapacity() != want {
			t.Errorf("got: %d  want: %d", darr.GetCapacity(), want)
		}
	}
}

func TestShrinkSize(t *testing.T) {
	darr := NewDynamicArray()
	for i := 0; i < 100; i++ {
		darr.Push("something 2")
	}

	want := darr.GetCapacity()
	for i := 100; i > 0; i-- {
		darr.Pop()
		if i == want/2 {
			want /= 2
		}

		if darr.GetCapacity() != want {
			t.Errorf("got: %d  want: %d", darr.GetCapacity(), want)
		}
	}

	// test the empty state
	val, err := darr.Pop()
	if err == nil {
		t.Errorf("got: %s  want: nil", err)
	}
	if val != nil {
		t.Errorf("got: %s  want: nil", val)
	}

}

func TestPush(t *testing.T) {
	want := "abc"
	darr := NewDynamicArray()
	darr.Push(want)
	val, err := darr.Get(0)
	if err != nil || val != want {
		t.Errorf("got: %s  want: %s", err, want)
	}
}

func TestPop(t *testing.T) {
	darr := NewDynamicArray()
	darr.Push("value 1")
	darr.Push("value 2")
	darr.Push("value 3")
	val, err := darr.Pop()
	if err != nil {
		t.Errorf("got: %s want: nil", err)
	}
	if val != "value 3" {
		t.Errorf("got: %s want: value 3", val)
	}
}
