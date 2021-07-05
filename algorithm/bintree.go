package algorithm

import (
	"fmt"
	"leek/base"
	"leek/common"
)

// CreateBinTree create binary tree from slice
// 广度优先遍历
func CreateBinTree(list []interface{}) (root *base.BinTree) {
	if len(list) == 0 {
		return nil
	}
	queue := common.NewQueue()
	root = base.NewBinTreeNode(list[0], nil, nil)
	queue.Push(root)
	i := 1
	var node *base.BinTree
	for i < len(list) {
		item, ok := queue.Pop()
		// 为空跳出
		if !ok {
			break
		}
		if node, ok = item.(*base.BinTree); !ok || node == nil {
			queue.Push(nil)
			queue.Push(nil)
			i += 2
			continue
		}
		fmt.Println("node: ", node)
		if list[i] == nil {
			node.Left = nil
		} else {
			node.Left = base.NewBinTreeNode(list[i], nil, nil)
		}
		i++
		if i >= len(list) {
			break
		}
		if list[i] == nil {
			node.Right = nil
		} else {
			node.Right = base.NewBinTreeNode(list[i], nil, nil)
		}
		i++
		fmt.Println("node out: ", node.Left, node.Right)
		queue.Push(node.Left)
		queue.Push(node.Right)
	}
	return
}

// ReCountBinTreeNode 计算二叉树中的节点个数
func ReCountBinTreeNode(root *base.BinTree) int {
	if root == nil {
		return 0
	}
	// 后续遍历
	leftCount := ReCountBinTreeNode(root.Left)
	rightCount := ReCountBinTreeNode(root.Right)
	root.Count = leftCount + rightCount + 1
	return root.Count
}

// PreOrder 先序遍历，先根遍历，
// 先遍历根节点，再遍历左子节点，最后遍历右子节点
func PreOrder(p *base.BinTree) {
	if p == nil {
		return
	}
	stack := make([]*base.BinTree, 0)
	stack = append(stack, p)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack[len(stack)-1] = nil
		fmt.Printf("preorder:[%v]\n", top.Data)
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
}

func RecursePreOrder(p *base.BinTree) {
	if p == nil {
		return
	}
	fmt.Printf("preorder: %d\n", p.Data.(int))
	RecursePreOrder(p.Left)
	RecursePreOrder(p.Right)
}

// InOrder 中序遍历二叉树，中根遍历
// 先遍历左子节点，再遍历根节点，最后遍历右子节点
func InOrder(p *base.BinTree) {
	if p == nil {
		return
	}
	stack := make([]*base.BinTree, 0)
	t := p

	for t != nil || len(stack) > 0 {
		// 左孩子入栈
		if t != nil {
			stack = append(stack, t)
			t = t.Left
		} else {
			// 出栈
			t = stack[len(stack)-1]
			fmt.Printf("inorder:[%v]\n", t.Data)
			stack[len(stack)-1] = nil
			stack = stack[:len(stack)-1]
			// 进入右孩子
			t = t.Right
		}
	}
}

// PostOrder 后续遍历，后根遍历
// 先遍历左子节点，再遍历右子节点，最后遍历根节点
func PostOrder(p *base.BinTree) {
	if p == nil {
		return
	}
	stack := make([]*base.BinTree, 0)
	t := p
	var prv *base.BinTree
	for len(stack) > 0 || t != nil {
		// 进入左子树最底层
		for t != nil {
			stack = append(stack, t)
			t = t.Left
		}

		t = stack[len(stack)-1]
		// t 的右子树为空 或者 t 的右子树是prv(上次遍历的记录), 说明上次刚遍历了右子树, 现在应该遍历根节点
		if t.Right == nil || prv == t.Right {
			fmt.Printf("postorder:[%v]\n", t.Data)
			stack[len(stack)-1] = nil
			stack = stack[:len(stack)-1]
			prv = t
			t = nil
		} else {
			t = t.Right
		}
	}
}

// LevelOrder 二叉树层次遍历
func LevelOrder(p *base.BinTree) []interface{} {
	if p == nil {
		return nil
	}
	list := make([]*base.BinTree, 0)
	list = append(list, p)
	result := make([]interface{}, 0)
	for len(list) > 0 {
		tmp := make([]*base.BinTree, 0)
		for i := 0; i < len(list); i++ {
			result = append(result, list[i].Data)
			// fmt.Printf("levelorder:[%v]\n", list[i].Data)
			if list[i].Left != nil {
				tmp = append(tmp, list[i].Left)
			}
			if list[i].Right != nil {
				tmp = append(tmp, list[i].Right)
			}
		}
		list = tmp
	}
	return result
}
