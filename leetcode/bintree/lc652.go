package bintree

import (
	"fmt"
	"leek/base"
)

/*
652. 寻找重复的子树
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

两棵树重复是指它们具有相同的结构以及相同的结点值。

示例 1：

1
/ \
2   3
/   / \
4   2   4
/
4
下面是两个重复的子树：

2
/
4
和

4
因此，你需要以列表的形式返回上述重复子树的根结点。
*/
var list []*base.BinTree
var record map[string]int

// FindDuplicateSubtrees  寻找重复的子树
func FindDuplicateSubtrees(root *base.BinTree) []*base.BinTree {
	list = make([]*base.BinTree, 0)
	record = make(map[string]int, 0)
	traverse(root)
	return list
}

func traverse(root *base.BinTree) string {
	if root == nil {
		return "#"
	}
	left := traverse(root.Left)
	right := traverse(root.Right)

	subModel := fmt.Sprintf("%v,%v,%d", left, right, root.Data)
	// 已经存在了
	if c, ok := record[subModel]; ok && c == 1 {
		list = append(list, root)
	}
	record[subModel]++
	return subModel
}
