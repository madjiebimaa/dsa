package binarysearchtree

import "github.com/madjiebimaa/dsa/tree/node"

type BST struct {
	Root *node.Node
}

func New() *BST {
	return &BST{}
}

func (bst *BST) Insert(val int, currentNode *node.Node) {
	if bst.Root == nil {
		bst.Root = node.New(val)
	} else if val < currentNode.Value {
		if currentNode.Left == nil {
			currentNode.Left = node.New(val)
		} else {
			bst.Insert(val, currentNode.Left)
		}
	} else if val > currentNode.Value {
		if currentNode.Right == nil {
			currentNode.Right = node.New(val)
		} else {
			bst.Insert(val, currentNode.Right)
		}
	}
}
