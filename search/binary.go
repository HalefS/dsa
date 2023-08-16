package search

type Number interface {
	int8 | int | int32 | int64 | float32 | float64
}

func BinarySearch[T Number](haystack []T, needle T) bool {
	high := len(haystack)
	low := 0
	for low < high {
		m := low + (high-low)/2
		midValue := haystack[m]
		if midValue == needle {
			return true
		} else if haystack[m] > needle {
			high = m
		} else {
			low = m + 1
		}
	}
	return false
}
