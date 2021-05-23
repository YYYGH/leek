package algorithm

import (
	"fmt"
	"leek/base"
)

// BinTree binary tree
type BinTree struct {
	Data  interface{}
	Left  *BinTree
	Right *BinTree
}

type Manager interface {
	PreOrder()
	Inorder()
	PostOrder()
	LevelOrder()
}

func NewBinTreeNode(data interface{}, left, right *BinTree) *BinTree {
	return &BinTree{
		Data:  data,
		Left:  left,
		Right: right,
	}
}

// CreateBinTree create binary tree from slice
// 广度优先遍历
func CreateBinTree(list []interface{}) (root *BinTree) {
	if len(list) == 0 {
		return nil
	}
	queue := base.NewQueue()
	root = NewBinTreeNode(list[0], nil, nil)
	queue.Push(root)
	i := 1
	for i < len(list) {
		item, ok := queue.Pop()
		// 为空跳出
		if !ok {
			break
		}
		node := item.(*BinTree)
		fmt.Println("node: ", node)
		node.Left = NewBinTreeNode(list[i], nil, nil)
		i++
		node.Right = NewBinTreeNode(list[i], nil, nil)
		i++
		fmt.Println("node out: ", node.Left, node.Right)
		queue.Push(node.Left)
		queue.Push(node.Right)
	}
	return
}

// PreOrder 先序遍历，先根遍历，
// 先遍历根节点，再遍历左子节点，最后遍历右子节点
func (p *BinTree) PreOrder() {
	if p == nil {
		return
	}
	stack := make([]*BinTree, 0)
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

// Inorder 中序遍历二叉树，中根遍历
// 先遍历左子节点，再遍历根节点，最后遍历右子节点
func (p *BinTree) Inorder() {
	if p == nil {
		return
	}
	stack := make([]*BinTree, 0)
	t := p

	for t != nil || len(stack) > 0 {
		// 左孩子入栈
		if t != nil {
			stack = append(stack, t)
			t = p.Left
		} else {
			// 出栈
			t = stack[len(stack)-1]
			fmt.Printf("preorder:[%v]\n", t.Data)
			stack[len(stack)-1] = nil
			stack = stack[:len(stack)-1]
			// 进入右孩子
			t = t.Right
		}
	}
}

// PostOrder 后续遍历，后根遍历
// 先遍历左子节点，再遍历右子节点，最后遍历根节点
func (p *BinTree) PostOrder() {
	if p == nil {
		return
	}
	stack := make([]*BinTree, 0)
	t := p
	var prv *BinTree
	for len(stack) > 0 || t != nil {
		// 进入左子树最底层
		for t != nil {
			stack = append(stack, t)
			t = t.Left
		}

		t := stack[len(stack)-1]
		if t.Right == nil || prv == t.Right {
			fmt.Printf("preorder:[%v]\n", t.Data)
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
func (p *BinTree) LevelOrder() []interface{} {
	list := make([]*BinTree, 0)
	list = append(list, p)
	result := make([]interface{}, 0)
	for len(list) > 0 {
		tmp := make([]*BinTree, 0)
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
