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
func (n *Node) IsRed() bool {
	if n == nil {
		return BLACK
	}
	return n.color
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
func (tree *RedBlackTree) Add(key, val int)  {
	isAdd, nd := tree.root.add(key, val)
	tree.size += isAdd
	tree.root = nd
	tree.root.color = BLACK
}

func (n *Node)add(key,val int)(int, *Node){
	if n==nil{
		return 1, NewNode(key,val)
	}
	Added:=0
	if key<n.key{
		Added,n.left=n.left.add(key,val)
	}else if key > n.key{
		Added,n.right=n.right.add(key,val)
	}else{
		n.value=val
	}
	n=n.updateRedBlackTree(Added)
	return Added,n
}




func (n *Node) updateRedBlackTree(Added int) *Node {
	if Added == 0 {
		return n
	}
	if n.right.IsRed() == RED && n.left.IsRed() != RED {
		n = n.leftRotate()
	}

	if n.left.IsRed() == RED && n.left.left.IsRed() == RED {
		n = n.rightRotate()
	}
	if n.left.IsRed() == RED && n.right.IsRed() == RED {
		n.changeColors()
	}

	return n
}