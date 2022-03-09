/*
 * @Author: 姚杨煜
 * @Email: yangyu.yao@shopee.com
 * @Date: 2022-03-09 20:50:20
 * @LastEditTime: 2022-03-09 20:52:07
 * @LastEditors: 姚杨煜
 * @Description:
 * @FilePath: /leek/base/bintree.go
 */
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

type BinTreeV2 struct {
	Val    interface{}
	Left   *BinTree
	Right  *BinTree
	Parent *BinTree
}
