package stack

type Node[T any] struct {
	Value T
	Prev  *Node[T]
}

// LinkedList based implementation of a generic Stack
type Stack[T any] struct {
	head *Node[T]
	size int
}

func New() *Stack[any] {
	return &Stack[any]{
		head: nil,
		size: 0,
	}
}

func (s *Stack[T]) Push(value T) {
	node := &Node[T]{Value: value}
	s.size++

	if s.head == nil {
		s.head = node
		return
	}

	node.Prev = s.head
	s.head = node
}

func (s *Stack[T]) Pop() *T {
	if s.head == nil {
		return nil
	}

	head := s.head
	s.head = head.Prev
	s.size--

	return &head.Value
}

func (s *Stack[T]) Peek() *T {
	if s.head == nil {
		return nil
	}

	return &s.head.Value
}

func (s *Stack[T]) Size() int {
	return s.size
}
