package bst

import (
	"leek/base"
)

/*
701. 二叉搜索树中的插入操作
给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
示例 1：
			4                    4
		  /   \                /   \
		2      7    ->        2     7
	  /  \                   / \   /
	 1    3                 1   3 5
输入：root = [4,2,7,1,3], val = 5
输出：[4,2,7,1,3,5]
解释：另一个满足题目要求可以通过的树是：[5, 2, 7, 1, 3, null, null, null, 4]

示例 2：
				5
			  /   \
			 2     7
		   /  \
          1	   3
		        \
				 4
输入：root = [40,20,60,10,30,50,70], val = 25
输出：[40,20,60,10,30,50,70,null,null,25] or [40, 25, 60, 10, 30, 50, 70, null, 20]
示例 3：

输入：root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
输出：[4,2,7,1,3,5]
*/
func InsertIntoBSTV1(root *base.BinTree, val int) *base.BinTree {
	if root == nil {
		return base.NewBinTreeNode(val, nil, nil)
	}
	if root.Data.(int) > val {
		root.Left = InsertIntoBSTV1(root.Left, val)
	} else if root.Data.(int) < val {
		root.Right = InsertIntoBSTV1(root.Right, val)
	}
	return root
}

func InsertIntoBSTV2(root *base.BinTree, val int) *base.BinTree {
	if root == nil {
		return base.NewBinTreeNode(val, nil, nil)
	}

	if root.Left == nil && root.Data.(int) > val {
		root.Left = InsertIntoBSTV2(root.Left, val)
	} else if root.Right == nil && root.Data.(int) < val {
		root.Right = InsertIntoBSTV2(root.Right, val)
	} else {
		if val < root.Data.(int) {
			return InsertIntoBSTV2(root.Left, val)
		} else if val > root.Data.(int) && val > root.Right.Data.(int) {
			return InsertIntoBSTV2(root.Right, val)
		}
		moveVal := root.Data.(int)
		root.Data = val
		p := root.Left
		for p.Right != nil {
			p = p.Right
		}
		p.Right = base.NewBinTreeNode(moveVal, nil, nil)
	}
	return root
}
