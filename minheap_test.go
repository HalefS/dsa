package dsa

import (
	"testing"
)

func TestHeapInsert(t *testing.T) {
	h := NewHeap[int](50)

	h.Insert(200)
	h.Insert(50)

	if h.size != 2 {
		t.Errorf("expected one element in the heap, go %d", h.size)
		return
	}
}

func TestHeapPop(t *testing.T) {
	h := NewHeap[int](50)

	h.Insert(200)
	h.Insert(50)
	h.Insert(300)
	h.Insert(10)

	minValue := h.Pop()
	if minValue != 10 {
		t.Errorf("expected min value to be 50, got %d", minValue)
		return
	}

	if h.size != 3 {
		t.Errorf("expected size of the heap to be 1, got %d", h.size)
	}

}
