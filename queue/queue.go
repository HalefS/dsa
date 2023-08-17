package queue

// Queue node
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// Linked list based implementation of a Queue
type Queue[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{length: 0, head: nil, tail: nil}
}

func (q *Queue[T]) Enqueue(value T) {
	node := &Node[T]{Value: value}
	q.length++

	if q.tail == nil {
		q.tail = node
		q.head = q.tail
		return
	}

	q.tail.Next = node
	q.tail = node
}

func (q *Queue[T]) Dequeue() *T {
	if q.head == nil {
		return nil
	}

	q.length--
	head := q.head
	q.head = q.head.Next

	return &head.Value
}

func (q *Queue[T]) Peek() *T {
	if q.length == 0 {
		return nil
	}

	return &q.head.Value
}
