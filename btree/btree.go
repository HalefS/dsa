package btree

type Number interface {
	int8 | int | int32 | int64 | float32 | float64
}

type Node[T Number] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func New[T Number](value T) *Node[T] {
	return &Node[T]{Value: value, Left: new(Node[T]), Right: new(Node[T])}
}

func (n *Node[T]) Insert(value T) {
	if n == nil {
		n.Value = value
	} else if value >= n.Value {
		n.InsertRight(value)
	} else {
		n.InsertLeft(value)
	}
}

func (n *Node[T]) InsertLeft(value T) {
	if value < n.Value {
		n.Left.Value = value
	} else {
		n.InsertLeft(value)
	}
}

func (n *Node[T]) InsertRight(value T) {
	if value >= n.Value {
		n.Right.Value = value
	} else {
		n.InsertRight(value)
	}
}

func (n *Node[T]) Search(value T) bool {
	current := n
	for current != nil {
		if current.Value == value {
			return true
		} else if value > current.Value {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return false
}
