package binarytree

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

type BinaryTree struct {
	Root *Node
}

// PreOrder 前序遍历，根->左->右
func (bt *BinaryTree) PreOrder() []interface{} {
	r := make([]interface{}, 0)
	doPreOrder(bt.Root, &r)
	return r
}

// doPreOrder
// 规则：
// 1、先访问当前节点（根）
// 2、然后递归地对左子树进行一次完整的前序遍历
// 3、最后递归地对右子树进行一次完整的前序遍历
func doPreOrder(node *Node, r *[]interface{}) {
	if node == nil {
		return
	}

	*r = append(*r, node.value)
	doPreOrder(node.left, r)
	doPreOrder(node.right, r)
}

// InOrder 中序遍历，左->根->右
func (bt *BinaryTree) InOrder() []interface{} {
	r := make([]interface{}, 0)
	doInOrder(bt.Root, &r)
	return r
}

// doInOrder
// 规则：
// 1、
func doInOrder(node *Node, r *[]interface{}) {
	if node == nil {
		return
	}

	doInOrder(node.left, r)
	*r = append(*r, node.value)
	doInOrder(node.right, r)

	return
}

// PostOrder 后序插入，左->右->根
func (bt *BinaryTree) PostOrder() []interface{} {
	r := make([]interface{}, 0)
	doPostOrder(bt.Root, &r)
	return r
}

func doPostOrder(node *Node, r *[]interface{}) {
	if node == nil {
		return
	}

	doPostOrder(node.left, r)
	doPostOrder(node.right, r)
	*r = append(*r, node.value)
}
