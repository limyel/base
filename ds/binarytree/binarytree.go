package binarytree

type node struct {
	value interface{}
	left  *node
	right *node
}

type BinaryTree struct {
	Root *node
}

// PreOrder 前序遍历，根->左->右
func (bt *BinaryTree) PreOrder() []interface{} {
	r := make([]interface{}, 0)

	return doPreOrder(bt.Root, r)
}

func doPreOrder(n *node, r []interface{}) []interface{} {
	if n == nil {
		return r
	}

	r = append(r, n.value)

	r = doPreOrder(n.left, r)
	r = doPreOrder(n.right, r)

	return r
}

// InOrder 中序遍历，左->根->右
func (bt *BinaryTree) InOrder() []interface{} {
	r := make([]interface{}, 0)

	return doInOrder(bt.Root, r)
}

func doInOrder(n *node, r []interface{}) []interface{} {
	if n == nil {
		return r
	}

	r = doInOrder(n.left, r)
	r = append(r, n.value)
	r = doInOrder(n.right, r)

	return r
}

// PostOrder 后序插入，左->右->根
func (bt *BinaryTree) PostOrder() []interface{} {
	r := make([]interface{}, 0)

	return doPostOrder(bt.Root, r)
}

func doPostOrder(n *node, r []interface{}) []interface{} {
	if n == nil {
		return r
	}

	r = doPostOrder(n.left, r)
	r = doPostOrder(n.right, r)
	r = append(r, n.value)

	return r
}
