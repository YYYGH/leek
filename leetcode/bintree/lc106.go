package bintree

import "leek/base"

/*
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/

func BuildTreeFromInorderAndPostorder(inorder []int, postorder []int) *base.BinTree {
	var buildChild func(instart, inend, poststart, postend int) *base.BinTree
	buildChild = func(instart, inend, poststart, postend int) *base.BinTree {
		if postend < poststart {
			return nil
		}

		rootVal := postorder[postend]
		index := instart
		for ; index <= inend; index++ {
			if inorder[index] == rootVal {
				break
			}
		}
		root := &base.BinTree{
			Data: rootVal,
		}
		lefSize := index - instart
		root.Left = buildChild(instart, index-1, poststart, poststart+lefSize-1)
		root.Right = buildChild(index+1, inend, poststart+lefSize, postend-1)
		return root
	}
	return buildChild(0, len(inorder)-1, 0, len(postorder)-1)
}

/*
func buildTree(inorder []int, postorder []int) *TreeNode {
    // 递归
	inorderMap := map[int]int{}
	for i, v := range inorder {
		inorderMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft int, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]

		root := &TreeNode{Val: val}
		root.Right = build(inorderMap[val]+1, inorderRight)
		root.Left = build(inorderLeft, inorderMap[val]-1)
		return root
	}
	return build(0, len(postorder)-1)
}
*/
