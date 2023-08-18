package sort

type Number interface {
	int8 | int | int32 | int64 | float32 | float64
}

func partition[T Number](arr []T, lo, hi int) int {
	pivot := T(hi)
	index := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			index++
			temp := arr[i]
			arr[i] = arr[index]
			arr[index] = temp
		}
	}
	// we need to perform the last swap so everything to the
	// left of the pivot is less and everything to the right is greater
	index++
	arr[hi] = arr[index]
	arr[index] = pivot

	return index // position of the pivot
}

func QuickSort[T Number](arr []T, lo, hi int) {
	// base case (we have reached a 1 element sorted array)
	if lo >= hi {
		return
	}

	pivotIndex := partition[T](arr, lo, hi)
	QuickSort(arr, lo, pivotIndex-1) // week sort left side excluding pivot
	QuickSort(arr, pivotIndex+1, hi) // week sort right side excluding pivot
}
