package sort

import (
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	values := []int{38, 56, 23, 1, 5, 3, 2, 20}
	BubbleSort(values)
	if !sort.IsSorted(sort.IntSlice(values)) {
		t.Errorf("expected slice to be sorted")
	}
}
