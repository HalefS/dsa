package trees

type Number interface {
	int | int32 | int8 | int64 | float32 | float64
}

type TreeNode[T Number] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func newNode[T Number](value T) *TreeNode[T] {
	return &TreeNode[T]{Value: value}
}

type BST[T Number] struct {
	Root *TreeNode[T]
}

func (b *BST[T]) Insert(value T) {
	root := b.Root
	if root == nil {
		root = &TreeNode[T]{Value: value}
	}

	if root.Value >= value {
		if root.Left == nil {
			root.Left = newNode[T](value)
		} else {
			insert[T](root.Left, value)
		}
	} else {
		if root.Right == nil {
			root.Right = newNode[T](value)
		} else {
			insert[T](root.Right, value)
		}
	}
}

func (b *BST[T]) Delete(value T) {
	// root := b.Root
}

func delete[T Number](root *TreeNode[T], value T) {
	if root == nil {
		return
	}
}

func insert[T Number](root *TreeNode[T], value T) {
	if root == nil {
		root = &TreeNode[T]{Value: value}
	}

	if root.Value >= value {
		root.Left = newNode[T](value)
	} else {
		root.Right = newNode[T](value)
	}
}

func preOrderWalk[T Number](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	*path = append(*path, node.Value)
	preOrderWalk[T](node.Left, path)
	preOrderWalk[T](node.Right, path)

	return path
}

func inOrderWalk[T Number](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	inOrderWalk[T](node.Left, path)
	*path = append(*path, node.Value)
	inOrderWalk[T](node.Right, path)

	return path
}

func postOrderWalk[T Number](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	inOrderWalk[T](node.Left, path)
	inOrderWalk[T](node.Right, path)
	*path = append(*path, node.Value)

	return path
}

func PreOrderSeach[T Number](root *TreeNode[T]) []T {
	return *preOrderWalk(root, &[]T{})
}

func InOrderSeach[T Number](root *TreeNode[T]) []T {
	return *inOrderWalk(root, &[]T{})
}

func PostOrderSearch[T Number](root *TreeNode[T]) []T {
	return *postOrderWalk(root, &[]T{})
}

func BinarySearch[T Number](root *TreeNode[T], value T) bool {
	if root == nil {
		return false
	}

	if root.Value == value {
		return true
	}

	if value > root.Value {
		return BinarySearch[T](root.Right, value)
	} else {
		return BinarySearch[T](root.Left, value)
	}
}

func BFS[T Number](root *TreeNode[T], value T) bool {
	q := []*TreeNode[T]{root}

	for len(q) != 0 {
		curr := q[0]
		q = q[1:]
		if curr.Value == value {
			return true
		}

		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		if curr.Right != nil {
			q = append(q, curr.Right)
		}
	}

	return false
}

func Compare[T Number](a, b *TreeNode[T]) bool {
	if a == nil && b == nil {
		return true
	}

	// if one is null and the other is not
	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return Compare[T](a.Left, b.Left) && Compare[T](a.Right, b.Right)
}
