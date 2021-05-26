package bst

import (
	"leek/base"
)

/*
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
示例 1：
          3
        /   \
	   1     4
	    \
		 2
输入：root = [3,1,4,null,2], k = 1
输出：1
示例 2：
					5
				  /   \
				 3     6
			   /   \
			  2     4
			/
		   1
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3
*/

var rank int = 0
var dstRank int = 0
var p *base.BinTree

func KthSmallest(root *base.BinTree, k int) int {
	dstRank = k
	traverseBSTInorder(root)
	return p.Data.(int)
}

func traverseBSTInorder(root *base.BinTree) {
	if root == nil {
		return
	}
	traverseBSTInorder(root.Left)
	rank++
	if rank == dstRank {
		p = root
		return
	}
	traverseBSTInorder(root.Right)
}
