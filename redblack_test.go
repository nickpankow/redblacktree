package redblack

import (
	"testing"
)

// Test the color setting abilities of the nodes.
func TestColor(t *testing.T) {
	rbt := new(RedBlackTree)

	rbt.SetRed()
	if rbt.color != Red {
		t.Error("Node color did not set correctly. Expected Red, Found ", rbt.GetColorString())
	}

	rbt.SetBlack()
	if rbt.color != Black {
		t.Error("Node color did not set correctly. Expected Black, Found ", rbt.GetColorString())
	}
}

// Test that the node sets its left child correctly.
func TestSetLeftChild(t *testing.T) {
	rbtRoot := new(RedBlackTree)
	rbtLeftChild := new(RedBlackTree)

	// Confirm setup
	if rbtRoot.GetLeftChild() != nil || rbtRoot.GetParent() != nil {
		t.Error("Node did not initialize properly.  Started with non-nil children or parent.")
	}
	if rbtLeftChild.GetLeftChild() != nil || rbtLeftChild.GetParent() != nil {
		t.Error("Node did not initialize properly.  Started with non-nil children or parent.")
	}

	// Check left child
	rbtRoot.SetLeftChild(rbtLeftChild)
	if rbtRoot.GetLeftChild() != rbtLeftChild {
		t.Error("Left child did not set properly.")
	}
	if rbtLeftChild.GetParent() != rbtRoot {
		t.Error("Left childs parent did not set correctly.")
	}
}

// Test that the node sets its right child correctly.
func TestSetRightChild(t *testing.T) {
	rbtRoot := new(RedBlackTree)
	rbtRightChild := new(RedBlackTree)

	// Confirm setup
	if rbtRoot.GetRightChild() != nil || rbtRoot.GetParent() != nil {
		t.Error("Node did not initialize properly.  Started with non-nil children or parent.")
	}
	if rbtRightChild.GetRightChild() != nil  || rbtRightChild.GetParent() != nil {
		t.Error("Node did not initialize properly.  Started with non-nil children or parent.")
	}

	// Check right child
	rbtRoot.SetRightChild(rbtRightChild)
	if rbtRoot.GetRightChild() != rbtRightChild {
		t.Error("Right child did not set properly.")
	}
	if rbtRightChild.GetParent() != rbtRoot {
		t.Error("Right childs parent did not set correctly.")
	}
}
