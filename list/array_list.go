package list

// slice based implementation of an array list since go does not
// permit declaring normal arrays on structs
type ArryaList[T any] struct {
	cap   int
	size  int
	array []T
}

func New[T any](cap int) *ArryaList[T] {
	return &ArryaList[T]{size: 0, cap: cap, array: make([]T, cap)}
}

func (l *ArryaList[T]) Insert(value T) {
	if l.size < l.cap {
		l.array[l.size] = value
	} else {
		l.growArray()
		l.array[l.size] = value
		l.size--
	}

	l.size++
}

func (l *ArryaList[T]) InsertAt(index int, value T) {
	if index > l.size || index < 0 {
		return
	}

	left := append(l.array[:index], value)
	right := l.array[index:]
	l.array = append(left, right...)

	l.size++
}

func (l *ArryaList[T]) DeleteAt(index int) {
	if index >= l.size || index < 0 {
		return
	}
	// split the internal array in two using the index provided
	// to get rid of it
	l.array = append(l.array[:index], l.array[index+1:]...)
	l.size--
}

func (l *ArryaList[T]) Get(index int) *T {
	if index >= l.size || index < 0 {
		return nil
	}

	return &l.array[index]
}

func (l *ArryaList[T]) Size() int {
	return l.size
}

func (l *ArryaList[T]) growArray() {
	newArray := make([]T, l.cap*2)
	newArray = append(newArray, l.array...)
	l.array = newArray
}
