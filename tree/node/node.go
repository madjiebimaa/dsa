package node

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func New(val int) *Node {
	return &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
}
