package bst

import "leek/base"

/*
700. 二叉搜索树中的搜索, https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
例如，
给定二叉搜索树:

        4
       / \
      2   7
     / \
    1   3

和值: 2
你应该返回如下子树:

      2
     / \
    1   3
*/
func SearchBST(root *base.BinTree, val int) *base.BinTree {
	if root == nil {
		return nil
	}
	if root.Data.(int) == val {
		return root
	} else if val < root.Data.(int) {
		return SearchBST(root.Left, val)
	}
	return SearchBST(root.Right, val)
}
