package redblack

import (
	"fmt"
)

const (
	Red = 	true
	Black = false
)

type RedBlackTree struct {
	Root 		*RedBlackNode
}

type RedBlackNode struct {
	Value 		int
	Data 		interface{}
	redblack	bool
	parent		*RedBlackNode
	left		*RedBlackNode
	right		*RedBlackNode
}

func NewTree() (*RedBlackTree){
	return new(RedBlackTree)
}

func (rb *RedBlackTree) Insert(element *RedBlackNode) {
	fmt.Println("Insert ", element)
	if rb.Root == nil {
		rb.Root = element
		// Root is always black
		rb.Root.SetBlack()
	} else {
		// Start by inserting node as you would as binary search tree and color it red
		recursiveInsert(&rb.Root, element, nil)
	 	return
	} 
}

func recursiveInsert(n **RedBlackNode, i *RedBlackNode, p *RedBlackNode) {
	if *n == nil {
		*n = i
		(*n).parent = p
		(*n).SetRed()
	} else if i.Value < (*n).Value {
		fmt.Println("Insert Left")
		recursiveInsert(&((*n).left), i, *n)
	} else {
		fmt.Println("Insert Right")
		recursiveInsert(&((*n).right), i, *n)
	}
}


func leftRotate(t *RedBlackTree, x *RedBlackNode) {
	var y *RedBlackNode
	y = x.right
	// Turn y's left sub-tree into x's right sub-tree
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	// y's new parent was x's parent
	y.parent = x.parent
	// Set parent to point to y instead of x
	
	// First see whether we're at the root
	if x.parent == nil {
		t.Root = y
	} else {
		if (x == x.parent.left) {
			// x was on the left of its parent
			x.parent.left = y
		} else {
			// x must have been on the right
			x.parent.right = y
		}
	}
	// Finally, put x on y's left
	y.SetLeftChild(x)
}

func rightRotate(t *RedBlackTree, y *RedBlackNode) {
	var x *RedBlackNode
	x = y.left

	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent
	if y.parent == nil {
		t.Root = x
	} else {
		if (y == y.parent.right) {
			y.parent.right = x
		} else {
			y.parent.left = x
		}
	}
	x.SetRightChild(y)
}

func (rb RedBlackTree) Delete(element int) (*RedBlackNode) {
	fmt.Println("Delete ", element)
	return nil
}

func (rb RedBlackTree) Find(value int) (*RedBlackNode){
	// fmt.Println("Find: ", value)
	return recursiveFind(rb.Root, value)
}

func recursiveFind(node *RedBlackNode, value int) (*RedBlackNode){
	// If matches current node, we have a match
	if node.Value == value {
		return node
	}

	var retNode *RedBlackNode
	// Search the left side of the tree
	if value < node.Value && node.left != nil {
		retNode = recursiveFind(node.left, value)
		if retNode != nil{
			// Return if found
			return retNode
		}
	}

	// Search the right side of the tree
	if value > node.Value && node.right != nil{
		retNode = recursiveFind(node.right, value)
	}

	// Return result from right side of the search; nil or node.
	return retNode
}

func (rb *RedBlackTree) Height() int {
	return height(rb.Root, 0)
}

func height(n *RedBlackNode, h int) int {
	fmt.Println("Node: ", n)
	if n == nil {
		return 0
	}
	l := height(n.left, h + 1)
	r := height(n.right, h + 1)
	if l > r {
		return l
	} else {
		return r
	}
}

/** Node Child Methods **/

func (rbn *RedBlackNode) SetLeftChild(x *RedBlackNode) {
	rbn.left = x
	x.parent = rbn
}

func (rbn *RedBlackNode) SetRightChild(x *RedBlackNode) {
	rbn.right = x
	x.parent = rbn
}

// func (rbn *RedBlackNode) Left() (l *RedBlackNode) {
// 	return rbn.left
// }

// func (rbn *RedBlackNode) Right() {} 

/** Node Methods **/

func (rbn *RedBlackNode) SetRed() {
	rbn.redblack = Red
}

func (rbn *RedBlackNode) SetBlack() {
	rbn.redblack = Black
}

/** Utility **/

func (rbn *RedBlackTree) String() {

}