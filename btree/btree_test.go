package btree

import "testing"

func TestInsert(t *testing.T) {

	expected := 10
	left := 5
	right := 20
	root := New(expected)
	root.Insert(right)
	root.Insert(left)

	if root.Get() != expected {
		t.Errorf("expected %d, got %d", expected, root.Get())
	}

	if root.Left.Value != left {
		t.Errorf("expected root.Left to be %d got %d", left, root.Left.Value)
	}

	if root.Right.Value != right {
		t.Errorf("expected root.Left to be %d got %d", left, root.Left.Value)
	}

}

func TestSearch(t *testing.T) {

	exp := 500

	root := New(30)
	root.Insert(400)
	root.Insert(10)

	root.Insert(600)
	root.Insert(200)
	root.Insert(12)
	root.Insert(44)

	root.Insert(exp)

	if !root.Search(exp) {
		t.Errorf("expected %d to exist on tree", exp)
	}

	if root.Search(1000) {
		t.Errorf("%d should not be present on the tree", 1000)
	}
}
