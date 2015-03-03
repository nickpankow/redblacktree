package redblack

import (
	"testing"
	"fmt"
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

	rbNode = RedBlackNode{}
	rbNode.Value = 25
	rbTree.Insert(&rbNode)
	if rbTree.Root.left != &rbNode || rbTree.Root.left.Value != 25{
		if rbTree.Root.left == nil {
			t.Error("Insert value mismatch.  No node found")
		} else {
			t.Error("Insert value mismatch.  Expected 25, found: ", rbTree.Root.left.Value)
		}
	}
	// t.Error("Insert not implemented")
}

func TestDelete(t *testing.T) {
	t.Error("Delete not implemented")
}

func createNodeWithValue(value int) (*RedBlackNode){
	return &RedBlackNode{ value, 
						  fmt.Sprintf("%d", value),
						  false, nil, nil, nil}
}

func TestFind(t *testing.T) {
	tree := NewTree()
	// var tmpNode RedBlackNode
	tree.Root = &RedBlackNode{}; tree.Root.Value = 50
	tree.Root.left = createNodeWithValue(25)
	tree.Root.left.left = createNodeWithValue(12)
	tree.Root.left.right = createNodeWithValue(37)
	tree.Root.right = createNodeWithValue(75)
	tree.Root.right.left = createNodeWithValue(62)
	tree.Root.right.right = createNodeWithValue(87)

	if tree.Root.right.right.Value != 87 || tree.Root.left.left.Value != 12 {
		t.Error("Tree setup does not match expected.")
	}

	for _, v := range []int{87,62,75,37,12,25} {
		x := tree.Find(v)
		if x == nil {
			t.Error("Find failed to find expected node: ", v)
		}
	}

	for _, w := range []int{1,20,55,69,88,100} {
		y := tree.Find(w)
		if y != nil {
			t.Error("Find erroneously found node: ", w)
		}
	}
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

func TestHeight(t *testing.T) {
	rbt := NewTree()
	// d := &RedBlackNode{ 0, "D", false, nil, nil, nil}
	// e := &RedBlackNode{ 1, "E", false, nil, nil, nil}
	a := &RedBlackNode{3, "A", false, nil, nil, nil}
	b := &RedBlackNode{4, "B", false, nil, nil, nil}
	// c := &RedBlackNode{2, "C", false, nil, nil, nil}

	if h := rbt.Height(); h != 0 {
		t.Error("Incorrect height.  Height should be 0, but is ", h)
	}

	rbt.Insert(a)
	if h := rbt.Height(); h != 1 {
		t.Error("Incorrect height.  Height should be 1, but is ", h)
	}

	rbt.Insert(b)
	if h := rbt.Height(); h != 2 {
		t.Error("Incorrect height.  Height should be 2, but is ", h)
	}
}