package bintree

import (
	"leek/base"
)

/*
105. 从前序与中序遍历序列构造二叉树
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
通过次数196,416提交次数281,722
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BuildTreeFromPreorderAndInorder(preorder []int, inorder []int) *base.BinTree {
	return buildChild(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildChild(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *base.BinTree {
	if preEnd < preStart {
		return nil
	}
	// 前序遍历 preorder = [3,   9,   20,   15,   7]
	//
	// 中序遍历 inorder =  [9,    3,   15,   20,   7]
	//                          index
	rootVal := preorder[preStart]
	index := inStart
	for ; index <= inEnd; index++ {
		if rootVal == inorder[index] {
			break
		}
	}

	leftSize := index - inStart
	root := &base.BinTree{
		Data: rootVal,
	}
	root.Left = buildChild(preorder, inorder, preStart+1, preStart+leftSize, inStart, index-1)
	root.Right = buildChild(preorder, inorder, preStart+leftSize+1, preEnd, index+1, inEnd)
	return root
}
