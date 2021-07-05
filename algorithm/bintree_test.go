package algorithm

import (
	"leek/common"
	"testing"
)

func TestPreOrder(t *testing.T) {
	array := []interface{}{40, 25, 60, 10, 30, 50, 70, nil, 20, nil, nil, 45}
	root := CreateBinTree(array)
	info := LevelOrder(root)
	common.PrintItems(info)
	PreOrder(root)
}

func TestInOrder(t *testing.T) {
	array := []interface{}{40, 25, 60, 10, 30, 50, 70, nil, 20, nil, nil, 45}
	root := CreateBinTree(array)
	info := LevelOrder(root)
	common.PrintItems(info)
	InOrder(root)
}

func TestPostOrder(t *testing.T) {
	array := []interface{}{40, 25, 60, 10, 30, 50, 70, nil, 20, nil, nil, 45}
	root := CreateBinTree(array)
	info := LevelOrder(root)
	common.PrintItems(info)
	PostOrder(root)
}
