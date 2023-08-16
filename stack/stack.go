package stack

// Generic Stack implementation
type Stack[T any] struct {
	items []T
}

func New() *Stack[any] {
	return &Stack[any]{
		items: make([]any, 0),
	}
}

func (s *Stack[T]) Put(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() *T {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &item
}

func (s *Stack[T]) Peek() *T {
	if len(s.items) == 0 {
		return nil
	}
	return &s.items[len(s.items)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
