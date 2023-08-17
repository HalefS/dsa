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

func prepareList(t *testing.T) *LinkedList[int] {
	t.Helper()

	l := &LinkedList[int]{}

	values := []int{10, 20, 30, 40, 50}
	for _, v := range values {
		l.Append(v)
	}

	return l
}
