package bst

import (
	"leek/base"
)

/*
230. 二叉搜索树中第K小的元素, https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
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

var rank int
var dstRank int
var p *base.BinTree

func KthSmallest(root *base.BinTree, k int) int {
	dstRank = k
	rank = 0
	p = nil
	traverseBSTInorder(root)
	if p == nil {
		return -1
	}
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

func KthSmallestV2(root *base.BinTree, k int) int {
	dstRank = k
	rank = 0
	p = nil
	if root == nil || root.Count < k {
		// 找不到
		return -1
	}
	leftCount := 0
	if root.Left != nil {
		if root.Left.Count >= k {
			return KthSmallestV2(root.Left, k)
		}
		leftCount = root.Left.Count
	}
	// 找到了
	if leftCount+1 == k {
		return root.Data.(int)
	}
	// 去右边找
	if root.Count-leftCount >= (k - leftCount - 1) {
		return KthSmallestV2(root.Right, k-leftCount-1)
	}

	return -1
}
