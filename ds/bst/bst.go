package bst

import "github.com/limyel/base/ds"

type node struct {
	value ds.Comparable
	left  *node
	right *node
}

// BST 二叉查找树
// 对于任意节点 N，左节点的值小于 N，右节点的值大于 N
type BST struct {
	Root *node
}

func (bst *BST) Insert(value ds.Comparable) {

}
