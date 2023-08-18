package list

import "testing"

func TestArrayInsert(t *testing.T) {
	l := New[int](10)

	want := 10
	l.Insert(want)
	if l.Size() != 1 {
		t.Errorf("expected l.Size() to be %d, got %d", 1, l.Size())
		return
	}

	if *l.Get(0) != want {
		t.Errorf("expected element at index 0 to be %d, got %d", want, *l.Get(0))
	}
}

func TestInsertAt(t *testing.T) {
	l := New[int](10)

	want := 10
	l.InsertAt(5, want)

	if l.Size() != 1 {
		t.Errorf("expected list size to be %d got %d", 1, l.Size())
		return
	}

	got := *l.Get(5)
	if want != got {
		t.Errorf("expected element at index 5 to be %d, got %d", want, got)
	}
}
