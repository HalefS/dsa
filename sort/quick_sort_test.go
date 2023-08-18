package sort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{30, 40, 50, 12, 44, 17, 2, 9, 78}
	QuickSort[int](arr, 0, len(arr)-1)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("expected slice to be sorted")
	}
}
