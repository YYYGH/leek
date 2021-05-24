package base

// BinTree binary tree
type BinTree struct {
	Data  interface{}
	Left  *BinTree
	Right *BinTree
}

func NewBinTreeNode(data interface{}, left, right *BinTree) *BinTree {
	return &BinTree{
		Data:  data,
		Left:  left,
		Right: right,
	}
}
