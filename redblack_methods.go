package redblack

// Public Methods

// Find a node with a given value in a tree.
func (rbt *RedBlackTree) Find(value int) (*RedBlackTree) {
	if rbt.Value == value {
		return rbt
	}

	// Search left side of the tree
	if rbt.GetLeftChild() != nil && value < rbt.Value {
		return rbt.GetLeftChild().Find(value)
	} else if rbt.GetRightChild() != nil && value >= rbt.Value {
	// Search right side of the tree
		return rbt.GetRightChild().Find(value)
	}
	// Didn't find anything
	return nil
}

// Insert a node into an existing tree.
func (rbt *RedBlackTree) Insert(element *RedBlackTree) {
	if element.Value < rbt.Value {
		if rbt.GetLeftChild() != nil {
			rbt.GetLeftChild().Insert(element)
		} else {
			rbt.SetLeftChild(element)
		}
	} else {
		if rbt.GetRightChild() != nil {
			rbt.GetRightChild().Insert(element)
		} else {
			rbt.SetRightChild(element)
		}
	}
}


