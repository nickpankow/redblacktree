package redblack

import (
	"testing"
)

func TestInsert(t *testing.T) {
	rbt := NewRedBlackTree(5, nil)
	rbt1 := NewRedBlackTree(3, nil)
	rbt2 := NewRedBlackTree(7, nil)
	rbt3 := NewRedBlackTree(1, nil)
	rbt4 := NewRedBlackTree(6, nil)
	rbt5 := NewRedBlackTree(8, nil)
	rbt6 := NewRedBlackTree(4, nil)

	rbt.Insert(rbt1)
	rbt.Insert(rbt2)
	rbt.Insert(rbt3)
	rbt.Insert(rbt4)
	rbt.Insert(rbt5)
	rbt.Insert(rbt6)

	chkFunc := func(node *RedBlackTree, val int) {
		if node == nil {
			t.Error("Node doesn't exist for value: ", val)
		} else if node.Value != val {
			t.Error("Node has incorrect value.  Expecting: ", val, " Found: ", node.Value)
		}
	}

	if rbt.Value != 5 {
		t.Error("Root has incorrect value.  Expecting: 5, Found: ", rbt.Value)
	}

	check := rbt.GetLeftChild()
	chkFunc(check, 3)

	check = rbt.GetRightChild()
	chkFunc(check, 7)	

	check = rbt.GetLeftChild()
	if (check != nil) {
		check = check.GetLeftChild()
		chkFunc(check, 1)
	} else {
		t.Error("Node has nil child.  Expecting child")
	}

	check = rbt.GetLeftChild()
	if (check != nil) {
		check = check.GetRightChild()
		chkFunc(check, 4)
	} else {
		t.Error("Node has nil child.  Expecting child")
	}
	
	check = rbt.GetRightChild()
	if (check != nil) {
		check = check.GetLeftChild()
		chkFunc(check, 6)
	} else {
		t.Error("Node has nil child.  Expecting child")
	}

	check = rbt.GetRightChild()
	if (check != nil) {
		check = check.GetRightChild()
		chkFunc(check, 8)
	} else {
		t.Error("Node has nil child.  Expecting child")
	}

}

func TestFind(t *testing.T) {
	rbt := NewRedBlackTree(5, nil)
	rbt1 := NewRedBlackTree(3, nil)
	rbt2 := NewRedBlackTree(7, nil)
	rbt3 := NewRedBlackTree(1, nil)
	rbt4 := NewRedBlackTree(6, nil)
	rbt5 := NewRedBlackTree(8, nil)
	rbt6 := NewRedBlackTree(4, nil)

	rbt.Insert(rbt1)
	rbt.Insert(rbt2)
	rbt.Insert(rbt3)
	rbt.Insert(rbt4)
	rbt.Insert(rbt5)
	rbt.Insert(rbt6)

	chkFunc := func(expected *RedBlackTree, found *RedBlackTree){
		if expected == nil  {
			t.Error("Expected node is nil.")
		} else if found == nil{
			t.Error("Did not find ", expected.Value)
		} else if expected != found {
			t.Error("Find mismatch.  Expected: ", expected, " Found: ", found)
		} else {
			t.Log("Found ", expected.Value)
		}
	}

	check := rbt.Find(rbt.Value)
	chkFunc(rbt, check)

	check = rbt.Find(rbt1.Value)
	chkFunc(rbt1, check)

	check = rbt.Find(rbt2.Value)
	chkFunc(rbt2, check)

	check = rbt.Find(rbt3.Value)
	chkFunc(rbt3, check)

	check = rbt.Find(rbt4.Value)
	chkFunc(rbt4, check)

	check = rbt.Find(rbt5.Value)
	chkFunc(rbt5, check)

	check = rbt.Find(rbt6.Value)
	chkFunc(rbt6, check)
}