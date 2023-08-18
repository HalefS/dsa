package trees

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func preOrderWalk[T any](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	*path = append(*path, node.Value)
	preOrderWalk[T](node.Left, path)
	preOrderWalk[T](node.Right, path)

	return path
}

func inOrderWalk[T any](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	inOrderWalk[T](node.Left, path)
	*path = append(*path, node.Value)
	inOrderWalk[T](node.Right, path)

	return path
}

func postOrderWalk[T any](node *TreeNode[T], path *[]T) *[]T {
	if node == nil {
		return path
	}

	inOrderWalk[T](node.Left, path)
	inOrderWalk[T](node.Right, path)
	*path = append(*path, node.Value)

	return path
}

func PreOrderSeach[T any](root *TreeNode[T]) []T {
	return *preOrderWalk(root, &[]T{})
}

func InOrderSeach[T any](root *TreeNode[T]) []T {
	return *inOrderWalk(root, &[]T{})
}

func PostOrderSearch[T any](root *TreeNode[T]) []T {
	return *postOrderWalk(root, &[]T{})
}
