package redblack

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	rbTree := NewTree()
	if (rbTree == nil || rbTree.Root != nil) {
		t.Error("New tree not initialized.")		
	}
}

func TestInsert(t *testing.T) {
	rbTree := NewTree()

	rbNode := RedBlackNode{}
	rbNode.Value = 50

	rbTree.Insert(&rbNode)

	if rbTree.Root == nil || rbTree.Root.Value != 50 {
		t.Error("Root node value mismatch.  Root: ", rbTree.Root)
	}
	// t.Error("Insert not implemented")
}

func TestDelete(t *testing.T) {
	t.Error("Delete not implemented")
}

func TestFind(t *testing.T) {
	t.Error("Find not implemented")
}

func TestRotateLeft(t *testing.T) {
	// Test with element as the root
	x := &RedBlackNode{ 0, "x", false, nil, nil, nil}
	y := &RedBlackNode{ 1, "y", false, nil, nil, nil}
	a := &RedBlackNode{3, "A", false, nil, nil, nil}
	b := &RedBlackNode{4, "B", false, nil, nil, nil}
	c := &RedBlackNode{2, "C", false, nil, nil, nil}

	x.SetLeftChild(a)
	x.SetRightChild(y)
	y.SetLeftChild(b)
	y.SetRightChild(c)

	tree := NewTree()
	tree.Root = x

	leftRotate(tree, x)
	if tree.Root != y || y.left != x || x.right != b || x.left != a || y.right != c {
		t.Error("Incorrect Left-Rotate result")
	}

	// Test with element not as the root
	z := &RedBlackNode{ 5, "z", false, nil, nil, nil}
	x.SetLeftChild(a)
	x.SetRightChild(y)
	y.SetLeftChild(b)
	y.SetRightChild(c)
	z.SetLeftChild(x)

	tree.Root = z
	leftRotate(tree, x)
	if tree.Root != z || y.left != x || x.right != b || x.left != a || y.right != c {
		t.Error("Incorrect Left-Rotate result")
	}
}

func TestRotateRight(t *testing.T) {
	// Test with element as the root
	x := &RedBlackNode{ 0, "x", false, nil, nil, nil}
	y := &RedBlackNode{ 1, "y", false, nil, nil, nil}
	a := &RedBlackNode{3, "A", false, nil, nil, nil}
	b := &RedBlackNode{4, "B", false, nil, nil, nil}
	c := &RedBlackNode{2, "C", false, nil, nil, nil}

	y.SetLeftChild(x)
	y.SetRightChild(c)
	x.SetLeftChild(a)
	x.SetRightChild(b)

	tree := NewTree()
	tree.Root = y

	rightRotate(tree, y)
	if tree.Root != x || x.left != a || x.right != y || y.left != b || y.right != c {
		t.Error("Incorrect Right-Rotate result")
	}

	// Test with element not as the root
	z := &RedBlackNode{ 5, "z", false, nil, nil, nil}
	y.SetLeftChild(x)
	y.SetRightChild(c)
	x.SetLeftChild(a)
	x.SetRightChild(b)
	z.SetRightChild(y)

	tree.Root = z
	rightRotate(tree, y)
	if tree.Root != z || x.left != a || x.right != y || y.left != b || y.right != c {
		t.Error("Incorrect Right-Rotate result")
	}
}
