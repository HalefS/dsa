package list

import (
	"errors"
	"fmt"
)

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

// Prepend adds a new element to the beginning o the list
func (l *LinkedList[T]) Prepend(value T) {
	node := &Node[T]{Value: value}
	l.size++

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	node.Next = l.head
	l.head.Prev = node
	l.head = node
}

// Append adds a new element to the end of the list
func (l *LinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value}
	l.size++

	if l.tail == nil {
		l.tail = node
		l.head = node
		return
	}

	l.tail.Next = node
	node.Prev = l.tail
	l.tail = node
}

// InsertAt adds a new element the the specified index on the list
func (l *LinkedList[T]) InsertAt(index int, value T) {
	if index > l.size || index < 0 {
		return
	}

	if index == 0 {
		l.Prepend(value)
		return
	} else if index == l.size {
		l.Append(value)
		return
	}

	node := &Node[T]{Value: value}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	prev := curr.Prev
	prev.Next = node
	node.Next = curr
	node.Prev = prev

	l.size++
}

// Get returns the element at index
func (l *LinkedList[T]) Get(index int) *T {
	if l.head == nil || index > l.size || index < 0 {
		return nil
	}

	curr := l.head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}

	return &curr.Value
}

// Delete removes the element at index
func (l *LinkedList[T]) Delete(index int) error {
	if index >= l.size || index < 0 {
		return errors.New(fmt.Sprintf("invalid index: usind %d on list of size %d", index, l.size))
	}

	if index == 0 {
		l.head = l.head.Next
		l.size--
		return nil
	} else if index == l.size-1 {
		l.tail = l.tail.Prev
		l.tail.Next = nil
		l.size--
		return nil
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	prev := curr.Prev
	curr = curr.Next
	curr.Prev = prev
	prev.Next = curr

	l.size--
	return nil
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
