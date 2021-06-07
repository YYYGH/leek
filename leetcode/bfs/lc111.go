package bfs

import "leek/base"

/*
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。
*/
func MinDepth(root *base.BinTree) int {
	if root == nil {
		return 0
	}
	q := make([]*base.BinTree, 0)
	q = append(q, root)
	depth := 1
	for len(q) != 0 {
		n := len(q)
		for index := 0; index < n; index++ {
			node := q[index]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
		q = q[n:]
	}
	return depth
}
