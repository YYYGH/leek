package bintree

import "leek/base"

/* 114. 二叉树展开为链表, https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

示例 1：
        1             1
      /   \            \
    2      5    ->      2
  /  \      \            \
 3    4      6            3
                           \
						    4
						     \
							  5
							   \
							    6
*/
// Flatten 展开二叉树
func Flatten(root *base.BinTree) {
	if root == nil {
		return
	}
	Flatten(root.Left)
	Flatten(root.Right)

	right := root.Right
	left := root.Left
	root.Right = left
	root.Left = nil

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
