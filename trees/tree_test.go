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

	if !BFS[int](&root, 28) {
		t.Errorf("expected 28 to be found on the tree")
		return
	}

	if BFS[int](&root, 100) {
		t.Errorf("did not expect 100 to be present in the tree")
	}
}

func TestCompare(t *testing.T) {
	a := BST[int]{Root: &TreeNode[int]{
		Value: 10,
		Left:  &TreeNode[int]{Value: 20},
		Right: &TreeNode[int]{Value: 30},
	}}

	b := BST[int]{Root: &TreeNode[int]{
		Value: 10,
		Left:  &TreeNode[int]{Value: 20},
		Right: &TreeNode[int]{Value: 30},
	}}

	if !Compare[int](a.Root, b.Root) {
		t.Error("expected trees to be equsl")
	}

	b.Root.Left.Value = 100

	if Compare[int](a.Root, b.Root) {
		t.Errorf("expected trees to be not equal")
	}

}

func TestBinarySearch(t *testing.T) {
	tree := &BST[int]{
		Root: &TreeNode[int]{
			Value: 10,
			Left: &TreeNode[int]{
				Value: 5,
			},
			Right: &TreeNode[int]{
				Value: 20,
				Left: &TreeNode[int]{
					Value: 15,
				},
				Right: &TreeNode[int]{
					Value: 30,
				},
			},
		},
	}

	if !BinarySearch[int](tree.Root, 15) {
		t.Errorf("expected 15 to present on the tree")
		return
	}

	if BinarySearch[int](tree.Root, 100) {
		t.Errorf("100 should not be present on the tree")
	}
}

func TestInsert(t *testing.T) {
	tree := &BST[int]{
		Root: &TreeNode[int]{
			Value: 10,
			Right: &TreeNode[int]{
				Value: 20,
			},
			Left: &TreeNode[int]{
				Value: 5,
			},
		},
	}

	newValue := 50
	Insert[int](tree.Root, newValue)

	if !BinarySearch[int](tree.Root, newValue) {
		t.Errorf("expected %d to be in the tree", newValue)
		return
	}

	if tree.Root.Right.Right.Value != newValue {
		t.Errorf("newValue was not inserted in right place")
	}
}
