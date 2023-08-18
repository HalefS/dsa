package list

import "testing"

func TestAppend(t *testing.T) {
	l := &LinkedList[int]{}

	values := []int{20, 50, 60, 30}
	for _, v := range values {
		l.Append(v)
	}

	if l.Size() != len(values) {
		t.Errorf("Expected s.Size to be %d, got %d", len(values), l.Size())
	}
}

func TestPrepend(t *testing.T) {
	l := &LinkedList[int]{size: 0}
	l.Prepend(10)
	l.Prepend(20)

	if l.Size() != 2 {
		t.Errorf("expected l.Size() to be %d, got %d", 2, l.Size())
	}

	if *l.Get(0) != 20 {
		t.Errorf("expected head to be %d, got %d", 20, *l.Get(0))
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name  string
		index int
		exp   int
	}{
		{name: "InsertAtZero", index: 0, exp: 40},
		{name: "InsertAtMiddle", index: 3, exp: 100},
		{name: "InsertAtEnd", index: 4, exp: 150},
	}

	for _, tc := range tests {
		l := prepareList(t)
		t.Run(tc.name, func(t *testing.T) {
			l.InsertAt(tc.index, tc.exp)
			if *l.Get(tc.index) != tc.exp {
				t.Fatalf("expected value to be %d, got %d", tc.exp, *l.Get(tc.index))
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		expValue int
		expSize  int
	}{
		{"DeleteHead", 0, 2, 4},
		{"DeleteAtMiddle", 2, 4, 4},
		{"DeleteTail", 4, 4, 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := prepareList(t)
			l.Delete(tc.index)

			if l.Size() != tc.expSize {
				t.Errorf("expected list size to be %d, got %d", tc.expSize, l.Size())
				return
			}

			if tc.index == l.Size() {
				tc.index--
			}

			if *l.Get(tc.index) != tc.expValue {
				t.Errorf("expected value at index %d to be %d, got %d", tc.index, tc.expValue, *l.Get(tc.index))
			}
		})
	}
}

func prepareList(t *testing.T) *LinkedList[int] {
	t.Helper()

	l := &LinkedList[int]{}

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		l.Append(v)
	}

	return l
}
