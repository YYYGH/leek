package bst

import "fmt"

/*
96. 不同的二叉搜索树, https://leetcode-cn.com/problems/unique-binary-search-trees/
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
示例 1：
	1      1             2             3             3
	 \      \          /   \          /            /
	  3      2        1     3        2            1
	 /        \                     /              \
	2          3                   1                2

输入：n = 3
输出：5
示例 2：

输入：n = 1
输出：1
*/

var matrix [][]int
var matrixv2 map[string]int

func NumTrees(n int) int {
	matrix = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		matrix[i] = make([]int, n+1)
	}
	matrixv2 = make(map[string]int)
	return CountChildTree(1, n)
}

func CountChildTree(low, hi int) int {
	if low > hi {
		return 1
	}
	if matrix[low][hi] != 0 {
		return matrix[low][hi]
	}

	res := 0
	for k := low; k <= hi; k++ {
		// low 作为根节点的情况
		leftCount := CountChildTree(low, k-1)
		rightCount := CountChildTree(k+1, hi)
		res += leftCount * rightCount
	}
	matrix[low][hi] = res
	return res
}

func CountChildTreeV2(low, hi int) int {
	if low > hi {
		return 1
	}
	k := fmt.Sprintf("%d%d", low, hi)
	if v, ok := matrixv2[k]; ok {
		return v
	}

	res := 0
	for k := low; k <= hi; k++ {
		// low 作为根节点的情况
		leftCount := CountChildTreeV2(low, k-1)
		rightCount := CountChildTreeV2(k+1, hi)
		res += leftCount * rightCount
	}
	matrixv2[k] = res
	return res
}
