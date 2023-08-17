package list

import "errors"

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

type LinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func (l *LinkedList[T]) Prepend(value T) {
	if l.head == nil {
		l.head = &Node[T]{Value: value}
	} else {
		newNode := &Node[T]{Value: value}
		newNode.Next = l.head
		l.head.Prev = newNode
		l.head = newNode
	}

	l.size++
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
		l.size++
	} else {
		l.tail.Next = newNode
		newNode.Prev = l.tail
		l.tail = newNode
		l.size++
	}
}

func (l *LinkedList[T]) InsertAt(index int, value T) {
	if index > l.size || index < 0 {
		return
	}

	newNode := &Node[T]{Value: value}

	if index == 0 {
		l.head.Prev = newNode
		newNode.Next = l.head
		l.head = newNode
		l.size++
		return
	}

	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	newNode.Prev = prev
	newNode.Next = prev.Next
	prev.Next = newNode

	l.size++
}

func (l *LinkedList[T]) Get(index int) *T {
	if l.head == nil || index > l.size || index < 0 {
		return nil
	}

	current := l.head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}

	return &current.Value
}

func (l *LinkedList[T]) Delete(index int) error {
	if index > l.size || index < 0 {
		return errors.New("invalid index")
	}

	if index == 0 {
		l.head = l.head.Prev
		if l.head != nil {
			l.head.Prev = nil
		}
		l.size--
		return nil
	}

	if index == l.size-1 {
		l.tail = l.tail.Prev
		l.size--
		return nil
	}

	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	prev.Next = prev.Next.Next
	prev.Next.Prev = prev
	l.size--
	return nil
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
