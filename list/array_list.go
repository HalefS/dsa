package list

type ArryaList[T any] struct {
	cap   int
	size  int
	array []T
}

func New[T any](cap int) *ArryaList[T] {
	return &ArryaList[T]{size: 0, cap: cap, array: make([]T, cap)}
}

func (l *ArryaList[T]) Insert(value T) {
	if l.size > l.cap {
		l.array = append(l.array, value)
		l.size++
		return
	}

	l.growArray()
}

func (l *ArryaList[T]) InsertAt(index int, value T) {
	if index > l.size || index < 0 {
		return
	}
	l.array = append(append(l.array[:index], value), l.array[index:]...)
	l.size++
}

func (l *ArryaList[T]) DeleteAt(index int) {

}

func (l *ArryaList[T]) Get(index int) *T {
	return nil
}

func (l *ArryaList[T]) Size() int {
	return l.size
}

func (l *ArryaList[T]) growArray() {
	newArray := make([]T, l.cap*2)
	newArray = append(newArray, l.array...)
	l.array = newArray
	l.size++
}
