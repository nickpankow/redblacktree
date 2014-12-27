package redblack

import (
	"fmt"
)

type RedBlackNode struct {
	value		interface{}
	redblack	bool
	left		*RedBlackNode
	right		*RedBlackNode
}

func (rb RedBlackNode) Insert(element interface{}) (*RedBlackNode) {
	fmt.Println("Insert ", element)
	return nil
}

func (rb RedBlackNode) Delete(element interface{}) (*RedBlackNode) {
	fmt.Println("Delete ", element)
	return nil
}

func (rb RedBlackNode) Find(value interface{}) (*RedBlackNode){
	fmt.Println("Find: ", value)
	return nil
}