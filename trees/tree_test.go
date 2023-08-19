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

func TestBFS(t *testing.T) {
	root := TreeNode[int]{Value: 50}

	left := TreeNode[int]{Value: 44}
	left.Left = &TreeNode[int]{Value: 10}
	left.Right = &TreeNode[int]{Value: 12}

	right := TreeNode[int]{Value: 13}
	right.Left = &TreeNode[int]{Value: 65}
	right.Right = &TreeNode[int]{Value: 28}

	root.Left = &left
	root.Right = &right

	if !Bfs[int](&root, 28) {
		t.Errorf("expected 28 to be found on the tree")
		return
	}

	if Bfs[int](&root, 100) {
		t.Errorf("did not expect 100 to be present in the tree")
	}
}
