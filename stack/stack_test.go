package stack

import "testing"

func TestPut(t *testing.T) {
	s := New()

	s.Put(200)
	s.Put(300)
	s.Put(1)

	if s.Size() != 3 {
		t.Errorf("stack size not %d. got=%d", 3, s.Size())
	}
}

func TestPop(t *testing.T) {
	s := New()

	item := s.Pop()
	if item != nil {
		t.Errorf("expected nil on Pop() with empty stack. got %q", *item)
	}

	expected := 500
	s.Put(expected)

	got := s.Pop()
	if expected != *got {
		t.Errorf("expected %q, got %q", expected, *got)
	}

	if s.Size() != 0 {
		t.Errorf("expected size %d. got %d", 0, s.Size())
	}

	s = New()
	if s.Peek() != nil {
		t.Errorf("Pop in empty stack should return nil")
	}
}

func TestPeek(t *testing.T) {
	s := New()

	expected := 500
	s.Put(expected)

	got := s.Peek()

	if expected != *got {
		t.Errorf("expected %q, got %q", expected, *got)
	}
}
