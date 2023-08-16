package btree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func New(value int) *Node {
	return &Node{Value: value, Left: new(Node), Right: new(Node)}
}

func (n *Node) Insert(value int) {
	if n.Value == 0 {
		n.Value = value
	} else if value >= n.Value {
		n.InsertRight(value)
	} else {
		n.InsertLeft(value)
	}
}

func (n *Node) Get() int {
	return n.Value
}

func (n *Node) InsertLeft(value int) {
	if value < n.Value {
		n.Left.Value = value
	} else {
		n.InsertLeft(value)
	}
}

func (n *Node) InsertRight(value int) {
	if value >= n.Value {
		n.Right.Value = value
	} else {
		n.InsertRight(value)
	}
}

func (n *Node) Search(value int) bool {
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
