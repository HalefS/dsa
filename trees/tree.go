package trees

type TreeNode[T comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func preOrderWalk[T comparable](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	*path = append(*path, node.Value)
	preOrderWalk[T](node.Left, path)
	preOrderWalk[T](node.Right, path)

	return path
}

func inOrderWalk[T comparable](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	inOrderWalk[T](node.Left, path)
	*path = append(*path, node.Value)
	inOrderWalk[T](node.Right, path)

	return path
}

func postOrderWalk[T comparable](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	inOrderWalk[T](node.Left, path)
	inOrderWalk[T](node.Right, path)
	*path = append(*path, node.Value)

	return path
}

func PreOrderSeach[T comparable](root *TreeNode[T]) []T {
	return *preOrderWalk(root, &[]T{})
}

func InOrderSeach[T comparable](root *TreeNode[T]) []T {
	return *inOrderWalk(root, &[]T{})
}

func PostOrderSearch[T comparable](root *TreeNode[T]) []T {
	return *postOrderWalk(root, &[]T{})
}

func Bfs[T comparable](root *TreeNode[T], value T) bool {
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
