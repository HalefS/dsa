package search

import "testing"

func TestSearch(t *testing.T) {
	values := []int{20, 40, 60, 80, 100, 120, 140, 160, 180, 200}
	if !BinarySearch[int](values, 180) {
		t.Errorf("Expected 54 to be found")
	}
}
