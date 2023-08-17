package list

import "errors"

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Size int
	Head *Node[T]
	Tail *Node[T]
}

func (l *LinkedList[T]) Prepend(value T) {
	if l.Head == nil {
		l.Head = &Node[T]{Value: value}
	} else {
		newNode := &Node[T]{Value: value}
		newNode.Next = l.Head
		l.Head.Prev = newNode
		l.Head = newNode
	}

	l.Size++
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = l.Head
		l.Size++
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail
		l.Tail = newNode
		l.Size++
	}
}

func (l *LinkedList[T]) InsertAt(index int, value T) {
	if index > l.Size || index < 0 {
		return
	}

	newNode := &Node[T]{Value: value}

	if index == 0 {
		l.Head.Prev = newNode
		newNode.Next = l.Head
		l.Head = newNode
		l.Size++
		return
	}

	prev := l.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	newNode.Prev = prev
	newNode.Next = prev.Next
	prev.Next = newNode

	l.Size++
}

func (l *LinkedList[T]) Get(index int) *T {
	if l.Head == nil || index > l.Size || index < 0 {
		return nil
	}

	prev := l.Head
	for i := 0; i <= index; i++ {
		prev = prev.Next
	}

	return &prev.Value
}

func (l *LinkedList[T]) Delete(index int) error {
	if index > l.Size || index < 0 {
		return errors.New("invalid index")
	}

	if index == 0 {
		l.Head = l.Head.Prev
		if l.Head != nil {
			l.Head.Prev = nil
		}
		l.Size--
		return nil
	}

	if index == l.Size-1 {
		l.Tail = l.Tail.Prev
		l.Size--
		return nil
	}

	prev := l.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	prev.Next = prev.Next.Next
	prev.Next.Prev = prev
	l.Size--
	return nil
}
