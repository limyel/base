package bst

import (
	"github.com/limyel/base/ds"
)

type Node struct {
	value ds.Comparable
	left  *Node
	right *Node
}

// BST 二叉查找树
// 对于任意节点 N，左节点的值小于 N，右节点的值大于 N
type BST struct {
	root *Node
}

func newNode(value ds.Comparable) *Node {
	return &Node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func NewBST() *BST {
	return &BST{
		root: nil,
	}
}

func (bst *BST) Search(value ds.Comparable) bool {
	return doSearch(bst.root, value)
}

func doSearch(node *Node, value ds.Comparable) bool {
	if node == nil {
		return false
	}

	if node.value.Compare(value) == 0 {
		return true
	}

	if value.Compare(node.value) < 0 {
		return doSearch(node.left, value)
	} else {
		return doSearch(node.right, value)
	}
}

func (bst *BST) Insert(value ds.Comparable) {
	bst.root = doInsert(bst.root, value)
}

func doInsert(node *Node, value ds.Comparable) *Node {
	if node == nil {
		return newNode(value)
	}

	if value.Compare(node.value) > 0 {
		node.right = doInsert(node.right, value)
	} else if value.Compare(node.value) < 0 {
		node.left = doInsert(node.left, value)
	}

	return node
}
