package bintree

import (
	"leek/base"
	"leek/common"
)

// InvertTree 翻转一个二叉树, 时间复杂度: O(n), 空间复杂度: O(n)
// leetcode 226. 翻转二叉树
func InvertTree(root *base.BinTree) *base.BinTree {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}

/*
     4                    4
   /   \                /   \
  2     7      ->      7     2
 / \   / \           /  \   / \
1   3 6   9         9    6 3   1
*/

// InvertTreeV2 翻转二叉树，版本2, 时间复杂度： O(n), 空间复杂度O(2^(h-1)), n=2^h - 1 -> h = log2^(n+1)
func InvertTreeV2(root *base.BinTree) *base.BinTree {
	if root == nil {
		return nil
	}
	q := common.NewQueue()
	q.Push(root)
	for !q.Empty() {
		node, ok := q.Pop()
		if !ok {
			break
		}
		pt := node.(*base.BinTree)
		temp := pt.Left
		pt.Left = pt.Right
		pt.Right = temp
		if pt.Left != nil {
			q.Push(pt.Left)
		}
		if pt.Right != nil {
			q.Push(pt.Right)
		}
	}
	return root
}
