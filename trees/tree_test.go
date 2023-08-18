package trees

import (
	"testing"
)

func TestPreeOrderSearch(t *testing.T) {
	root := &TreeNode[int]{Value: 10}

	root.Left = &TreeNode[int]{Value: 50}
	root.Right = &TreeNode[int]{Value: 100}

	want := []int{10, 50, 100}
	got := PreOrderSeach[int](root)

	if len(want) != len(got) {
		t.Errorf("expected result to have length %d, got %d", len(want), len(got))
	}

	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Expected %q, got %q", want, got)
			return
		}
	}
}

func TestInOrderSearch(t *testing.T) {
	root := &TreeNode[int]{Value: 10}

	root.Left = &TreeNode[int]{Value: 50}
	root.Right = &TreeNode[int]{Value: 100}

	want := []int{50, 10, 100}
	got := InOrderSeach[int](root)

	if len(want) != len(got) {
		t.Errorf("expected result to have length %d, got %d", len(want), len(got))
	}

	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, got)
			return
		}
	}
}

func TestPostOrderSearch(t *testing.T) {
	root := &TreeNode[int]{Value: 10}

	root.Left = &TreeNode[int]{Value: 50}
	root.Right = &TreeNode[int]{Value: 100}

	want := []int{50, 100, 10}
	got := PostOrderSearch[int](root)

	if len(want) != len(got) {
		t.Errorf("expected result to have length %d, got %d", len(want), len(got))
	}

	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, got)
			return
		}
	}
}
