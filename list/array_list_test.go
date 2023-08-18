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
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	l := prepareArrayList(t, values)

	want := 10
	l.InsertAt(5, want)

	if l.Size() != len(values)+1 {
		t.Errorf("expected list size to be %d got %d", len(values)+1, l.Size())
		return
	}

	got := *l.Get(5)
	if want != got {
		t.Errorf("expected element at index 5 to be %d, got %d", want, got)
	}
}

func prepareArrayList(t *testing.T, values []int) *ArryaList[int] {
	t.Helper()
	l := New[int](10)
	for i, v := range values {
		if i >= l.cap {
			break
		}
		l.Insert(v)
	}

	return l
}
