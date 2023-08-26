package dsa

type Number interface {
	int | int8 | int32 | int64 | float32
}

type MinHeap[T Number] struct {
	elements []T
	size     int
	cap      int
}

func NewHeap[T Number](cap int) *MinHeap[T] {
	return &MinHeap[T]{
		elements: make([]T, cap),
		size:     0,
		cap:      cap,
	}
}

func (h *MinHeap[T]) Insert(value T) {
	h.elements[h.size] = value
	h.heapifyUp(h.size)
	h.size++
	// expand array if needed
	if h.size == h.cap {
		arr := make([]T, h.cap*2)
		arr = append(arr, h.elements...)
		h.elements = arr
		h.cap *= 2
	}
}

func (h *MinHeap[T]) Pop() T {
	if h.size == 0 {
		return -1
	}

	out := h.elements[0]
	h.size--
	if h.size == 0 {
		h.elements = make([]T, 0)
		return out
	}

	h.elements[0] = h.elements[h.size]
	h.heapifyDown(0)

	return out
}

func (h *MinHeap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parent := parent(idx)
	parentV := h.elements[parent]
	v := h.elements[idx]

	if parentV > v {
		h.elements[parent] = v
		h.elements[idx] = parentV
		h.heapifyUp(parent)
	}
}

func (h *MinHeap[T]) heapifyDown(idx int) {
	// we have reached the end of our heap and are ready to stop
	// we need this indexes to select which child to swap with
	lIdx := h.leftChild(idx)
	rIdx := h.rightChild(idx)
	if idx >= h.size || lIdx >= h.size {
		return
	}

	// parent and children values
	lv := h.elements[lIdx]
	rv := h.elements[rIdx]
	v := h.elements[idx]

	if lv > rv && v > lv { // heapify left
		h.elements[lIdx] = v
		h.elements[idx] = lv
		h.heapifyDown(rIdx)
	} else if rv > lv && v > rv { // heapify right
		h.elements[rIdx] = v
		h.elements[idx] = rv
		h.heapifyDown(lIdx)
	}

}

func parent(idx int) int {
	return (idx - 1) / 2
}

func (h *MinHeap[T]) leftChild(idx int) int {
	return (idx*2 + 1) % h.size
}

func (h *MinHeap[T]) rightChild(idx int) int {
	return (idx*2 + 2) % h.size
}

func (h *MinHeap[T]) Size() int {
	return h.size
}
