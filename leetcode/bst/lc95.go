package bst

import (
	"fmt"
	"leek/base"
)

/**
95. 不同的二叉搜索树 II, https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
示例 1：
	1        1             2              3       3
	 \        \          /   \           /       /
	  3        2        1     3         2       1
	 /          \                      /         \
    2            3                    1           2
输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
示例 2：

输入：n = 1
输出：[[1]]
*/

var record map[string][]*base.BinTree

func GenerateTrees(n int) []*base.BinTree {
	record = make(map[string][]*base.BinTree)
	return BuildTreesFromInterval(1, n)
}

func BuildTreesFromInterval(low, hi int) []*base.BinTree {
	list := make([]*base.BinTree, 0)
	if low > hi {
		list = append(list, nil)
		return list
	}
	k := fmt.Sprintf("%d%d", low, hi)
	if v, ok := record[k]; ok {
		return v
	}

	for k := low; k <= hi; k++ {
		leftTrees := BuildTreesFromInterval(low, k-1)
		rightTrees := BuildTreesFromInterval(k+1, hi)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := base.NewBinTreeNode(k, left, right)
				list = append(list, root)
			}
		}
	}
	record[k] = list
	return list
}
