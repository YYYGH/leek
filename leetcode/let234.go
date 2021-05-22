package leetcode

import (
	al "leek/algorithm"
)

/*
234. 回文链表
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
// IsPalindromeV1 快慢指针 + 反转链表, 时间复杂度: O(n), 空间复杂度O(1)
func IsPalindromeV1(head *al.ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head
	// 查找中点  O(n/2)
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 当链表节点为奇数个时, fast 刚好指向最后一个元素, slow刚好指向中间节点，slow需要指向next节点, 中间节点无需比较
	if fast != nil {
		slow = slow.Next
	}
	//O(n/2)
	newHead := reverse(slow)
	left := head
	right := newHead

	// O(n/2)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse(head *al.ListNode) *al.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &al.ListNode{
		Val:  -1,
		Next: head,
	}
	prev := dummyNode
	cur := prev.Next

	// 1,2,3,4,5
	for prev.Next != nil && cur.Next != nil {
		removed := cur.Next
		cur.Next = removed.Next
		removed.Next = prev.Next
		prev.Next = removed
	}
	return dummyNode.Next
}

var left *al.ListNode

// IsPalindromeV2 递归判断, 时间复杂度O(n), 空间复杂度:O(n)
func IsPalindromeV2(head *al.ListNode) bool {
	left = head
	return traverse(left)
}

func traverse(right *al.ListNode) bool {
	if right == nil {
		return true
	}

	res := traverse(right.Next)
	res = res && left.Val == right.Val
	left = left.Next
	return res
}

// IsPalindromeV3, 时间复杂度O(n), 空间复杂度O(1)
// 1. 快慢指针, 遍历的同时反转节点, 一直到中间节点(相当于反转链表中间节点之前的节点)
// 2. left=newhead  right=middle.Next 判断元素是否相等
func IsPalindromeV3(head *al.ListNode) bool {
	if head == nil {
		return false
	}

	dummyNode := al.NewListNode(-1)
	dummyNode.Next = head
	prev := dummyNode
	cur := prev.Next
	fast := head
	// 查找中点
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		removed := cur.Next
		cur.Next = removed.Next
		removed.Next = prev.Next
		prev.Next = removed
	}
	left := dummyNode.Next
	right := cur.Next
	al.OuputListNode(left)
	al.OuputListNode(right)
	// 奇数个节点, 中间节点被反转了, 中间节点变成了头结点, 头结点无需比较
	// 例如：1->2->3->2->1   ->  3->2->1->2->1
	if fast.Next == nil {
		left = left.Next
	}
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
