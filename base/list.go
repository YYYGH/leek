package base

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// CreateListNode 创建一个链表
func CreateListNode(array []int) *ListNode {
	var header *ListNode
	var p *ListNode
	for _, v := range array {
		node := NewListNode(v)
		if header == nil {
			header = node
		} else {
			p.Next = node
		}
		p = node
	}
	return header
}

// OuputListNode 输出链表
func OuputListNode(header *ListNode) {
	for header != nil && header.Next != nil {
		fmt.Printf("%d -> ", header.Val)
		header = header.Next
	}
	if header != nil {
		fmt.Printf("%d\n", header.Val)
	}
}
