package btree

// Item B 树键的接口
type Item interface {
	Less(item Item) bool
}

type items []Item

type nodes []*Node

type Node struct {
	keys     items // 键
	isLean   bool
	children nodes // 子节点
}

// BTree B 树
// 除了根节点之外：
// 1、任意节点至少有 t-1 个键，最多有 2t-1 个键。这保证了节点空间既不会太浪费，也不会无限膨胀。
// 2、任意节点只要不是叶子节点，如果它有 k 个键，那么它必须有 k+1 个子节点。
// 3、节点内所有的键都是从小到大排序。
// 4、所有叶子节点都必须在同一层。这保证查找任何数据经过的路径长度都是一样的。
type BTree struct {
	root *Node // 根节点
	t    int   // B 树的阶，t >= 2
}

func newNode(isLeaf bool) *Node {
	return &Node{
		keys:     make(items, 0),
		children: make(nodes, 0),
		isLean:   isLeaf,
	}
}

func New(t int) *BTree {
	return &BTree{
		root: newNode(true),
		t:    t,
	}
}

func equals(i1, i2 Item) bool {
	return !i1.Less(i2) && !i2.Less(i1)
}

func (bt *BTree) Search(item Item) (*Node, int) {
	return bt.search(bt.root, item)
}

func (bt *BTree) search(currentNode *Node, target Item) (*Node, int) {
	i := 0
	// 在当前节点中找到第一个不小于 target 的 key
	for i < len(currentNode.keys) && currentNode.keys[i].Less(target) {
		i += 1
	}

	// 检查 key 和 target 是否相等
	if i < len(currentNode.keys) && equals(target, currentNode.keys[i]) {
		return currentNode, i
	}

	// 如果没找到，且当前是叶子节点，则不存在
	if currentNode.isLean {
		return nil, -1
	}

	// 如果不是叶子节点，则递归查询
	return bt.search(currentNode.children[i], target)
}

// splitChild 节点分裂
// n 要分裂的节点
// childIdx
func (bt *BTree) splitChild(n *Node, childIdx int) {

}

// Insert 插入
// B 树永远不会向一个满的节点插入新的键。向下遍历时如果遇到满的节点，就马上分裂它。
func (bt *BTree) Insert(item Item) {
	if len(bt.root.keys) == bt.t*2-1 {
		newRoot := newNode(false)
		newRoot.children = append(newRoot.children)
		// 分裂旧的root
	}

}

func (bt *BTree) insertNonfull(n *Node, item Item) {

}
