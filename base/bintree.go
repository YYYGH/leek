package base

// BinTree binary tree
type BinTree struct {
	Data  interface{}
	Left  *BinTree
	Right *BinTree
	// 以当前节点为根节点的树的总结点数
	Count int
}

func NewBinTreeNode(data interface{}, left, right *BinTree) *BinTree {
	return &BinTree{
		Data:  data,
		Left:  left,
		Right: right,
	}
}
