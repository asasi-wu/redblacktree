package RedBlackTree
const(
	RED bool =true
	BLACK bool =false
)
type Node struct{
	key int
	value interface{}

	left *Node
	right *Node
	parent *Node

	color bool
}
type RedBlackTree struct{
	size int
	root *Node
}
func NewNode(key,val int)*Node{
	return &Node{
		key:    key,
		value:  val,
		left:   nil,
		right:  nil,
		parent: nil,
		color:  RED,

	}
}
func NewRedBlackTree()*RedBlackTree{
	return &RedBlackTree{}
}
func (n *Node)leftRotate()*Node{
	retNode := n.right
	n.right = retNode.left
	retNode.left = n
	retNode.color = n.color
	n.color = RED
	return retNode
}
func (n *Node) rightRotate() *Node {
	//右旋转
	retNode := n.left
	n.left = retNode.right

	retNode.right = n
	retNode.color = n.color
	n.color = RED

	return retNode
}
func (n *Node) changeColors() {
	n.color = RED
	n.left.color = BLACK
	n.right.color = BLACK
}
