package btree

/*
 * 对于一棵 m（表示一个节点最多可以拥有多少个子节点） 阶 B 树来说：
 * * 每个节点最多包含 m-1 个键
 * * 除根节点外，每个节点至少包含 m/2-1（向上取整） 个键
 * * 根节点至少包含 1 个键（除非树为空）
 *
 * * 每个非叶子节点最多有 m 个子节点
 * * 除根节点外，每个非叶子节点至少有 m/2 个子节点
 * * 如果根节点不是叶子节点，则它至少有两个子节点
 *
 * * 节点内的键必须按升序排列
 * * 对于任意节点中的键 k，左子树中所有的键都小于 k，右子树中所有的键都大于 k，中间子树中所有的键必须在相邻两个键之间
 *
 * * 所有叶子节点必须在同一层（从根到任何子节点的路径长度都相同）
 */

// Item B 树的键接口
type Item interface {
	Less(Item) bool
}

// Node B 树节点
type Node struct {
	items    []Item  // 节点的键
	children []*Node // 子节点
	isLeaf   bool    // 是否为叶子节点
}

// BTree B 树
type BTree struct {
	root *Node // 根节点
	t    int   // t，t = m/2（向上取整）所以 m = 2t
}

func equal(a, b Item) bool {
	return !a.Less(b) && !b.Less(a)
}

// createNode 创建一个节点
func createNode(isLeaf bool) *Node {
	return &Node{
		items:    make([]Item, 0),
		children: make([]*Node, 0),
		isLeaf:   isLeaf,
	}
}

// initBTree 初始化 B 树
func initBTree(t int) *BTree {
	return &BTree{
		root: createNode(true),
		t:    t,
	}
}

// search 在节点 n 中寻找 item
func (n *Node) search(item Item) int {
	i := 0
	// 找到第一个大于等于 item 的键的位置
	for i < len(n.items) && (n.items[i]).Less(item) {
		i++
	}
	return i
}

// search 在 B 树中寻找 item
func (bt *BTree) search(n *Node, item Item) (*Node, int) {
	i := n.search(item)

	// 如果找到了，返回节点和索引
	if i < len(n.items) && equal(item, n.items[i]) {
		return n, i
	}

	// 如果没有找到，返回子节点和索引
	if n.isLeaf {
		return nil, -1
	}

	// 递归调用，
	return bt.search(n.children[i], item)
}

// insert 在 B 树中插入 item
func (bt *BTree) insert(item Item) {
	if bt.root == nil {
		bt.root = createNode(false)
		bt.root.items = append(bt.root.items, item)
	} else {
		newRoot := createNode(false)
		newRoot.children = append(newRoot.children, bt.root)

	}
}

func (n *Node) splitChild(i int, t int) {
	// 要分裂的子节点
	node := n.children[i]

	// todo ? 创建新节点
	newNode := createNode(node.isLeaf)
	newNode.items = node.items[t:]

	// 如果不是叶子节点，移动子节点
	if !node.isLeaf {
		newNode.children = node.children[t:]
	}

	// 将中间的键移动到父节点
	n.items = append(n.items, node.items[t-1])

}

func insertNonFull() {

}

func splitChild() {

}
