package bst

import (
	"fmt"
	"leek/base"
)

/*
450. 删除二叉搜索树中的节点, https://leetcode-cn.com/problems/delete-node-in-a-bst/
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7

*/

func DeleteNodeFromBST(root *base.BinTree, key int) *base.BinTree {
	TraverseDeleteBSTNode(&root, key)
	return root
}

func TraverseDeleteBSTNode(root **base.BinTree, key int) {
	if root == nil || *root == nil {
		return
	}
	p := *root
	if p.Data.(int) == key {
		if p.Left == nil {
			fmt.Println("left == nil", p.Data)
			*root = p.Right
		} else if p.Right == nil {
			fmt.Println("right == nil", p.Data)
			*root = p.Left
		} else {
			// min := getMin(p.Right)
			// TraverseDeleteBSTNode(&p.Right, min.Data.(int))
			min := Deleteleaf(p)
			p.Data = min.Data
			fmt.Println("left, right", p.Data, min.Data)
		}
		return
	}
	if key < p.Data.(int) {
		TraverseDeleteBSTNode(&p.Left, key)
	} else if key > p.Data.(int) {
		TraverseDeleteBSTNode(&p.Right, key)
	}
}
func getMin(root *base.BinTree) *base.BinTree {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func Deleteleaf(root *base.BinTree) *base.BinTree {
	prev := root
	child := root.Right
	for child.Left != nil {
		if prev == root {
			prev = prev.Right
		} else {
			prev = prev.Left
		}
		child = child.Left
	}
	// 要操作的就是根节点
	if prev == root {
		prev.Right = child.Right
		child.Right = nil
	} else {
		prev.Left = child.Right
	}
	return child
}

func DeleteNodeFromBSTV2(root *base.BinTree, key int) *base.BinTree {
	if root == nil {
		return root
	}
	if root.Data.(int) == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := getMin(root.Right)
		root.Data = minNode.Data
		root.Right = DeleteNodeFromBSTV2(root.Right, minNode.Data.(int))
	} else if root.Data.(int) < key {
		root.Right = DeleteNodeFromBSTV2(root.Right, key)
	} else if root.Data.(int) > key {
		root.Left = DeleteNodeFromBSTV2(root.Left, key)
	}
	return root
}
