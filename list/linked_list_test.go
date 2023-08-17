package list

import "testing"

func TestAppend(t *testing.T) {
	l := &LinkedList[int]{}

	values := []int{20, 50, 60, 30}
	for _, v := range values {
		l.Append(v)
	}

	if l.Size != len(values) {
		t.Errorf("Expected s.Size to be %d, got %d", len(values), l.Size)
	}
}
