package redblack

import (
	"fmt"
)

// Represents the color of a RedBlackTree node.  A node can be either Red or Black.
const (
	Red = 	true
	Black = false
)

type RedBlackTree struct {
	Value 		int
	Data 		interface{}
	color		bool
	parent		*RedBlackTree
	left		*RedBlackTree
	right		*RedBlackTree
}


/** Data Structure Basics **/

func NewRedBlackTree(value int, data interface{}) (*RedBlackTree){
	rbt := new(RedBlackTree)
	rbt.Value = value
	rbt.Data = data
	return rbt
}

// Sets the current nodes color to Red.
func (rbt *RedBlackTree) SetRed() {
	rbt.color = Red
}

// Sets the current nodes color to Black.
func (rbt *RedBlackTree) SetBlack() {
	rbt.color = Black
}

// Returns the left child of the current node.
func (rbt *RedBlackTree) GetLeftChild() (*RedBlackTree) {
	return rbt.left
}

// Returns the right chid of the current node.
func (rbt *RedBlackTree) GetRightChild() (*RedBlackTree) {
	return rbt.right
}

// Returns the parent of the current node.
func (rbt *RedBlackTree) GetParent() (*RedBlackTree) {
	return rbt.parent
}

// Sets the item as the left child of the current RedBlackTree object
// Also sets teh current RedBlackTree object as teh given childs parent
func (rbt *RedBlackTree) SetLeftChild(child *RedBlackTree) {
	rbt.left = child
	child.parent = rbt
}

// Sets the item as the right child of current RedBlackTree object
// Also sets the current RedBlackTree object as the given childs parent
func (rbt *RedBlackTree) SetRightChild(child *RedBlackTree) {
	rbt.right = child
	child.parent = rbt
}

// Returns a textual representation of a given color.
func (rbt *RedBlackTree) GetColorString() string{
	if rbt.color == Red{
		return "Red"
	} else {
		return "Black"
	}
}

// 
func (rbt *RedBlackTree) String() string{
	return fmt.Sprintf("%d", rbt.Value)
}