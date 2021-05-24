package algorithm

import (
	"fmt"
	"leek/base"
)

// RecurseReverseListNode 递归反转单链表
func RecurseReverseListNode(header *base.ListNode) *base.ListNode {
	if header == nil || header.Next == nil {
		return header
	}
	last := RecurseReverseListNode(header.Next)
	header.Next.Next = header
	header.Next = nil
	return last
}

// ReverseListNode 非递归反转链表
func ReverseListNode(header *base.ListNode) *base.ListNode {
	if header == nil || header.Next == nil {
		return header
	}
	var cursor *base.ListNode
	p := header
	for header.Next != nil {
		cursor = header.Next
		header.Next = cursor.Next
		cursor.Next = p
		p = cursor
	}
	return p
}

// ReverseListNodeV2 非递归反转链表
func ReverseListNodeV2(head *base.ListNode) *base.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &base.ListNode{
		Val:  -1,
		Next: head,
	}
	prev := dummyNode
	cur := prev.Next

	//       1,     2,   3,   4,   5
	// prev  cur   re
	for prev.Next != nil && cur.Next != nil {
		removed := cur.Next
		cur.Next = removed.Next
		removed.Next = prev.Next
		prev.Next = removed
	}
	return dummyNode.Next
}

var Cursor *base.ListNode

// RecurseReverseListNodeN 递归反转链表的前n个节点
func RecurseReverseListNodeN(header *base.ListNode, n int) *base.ListNode {
	// fmt.Printf("%d, %p\n", n, Cursor)
	if header == nil || header.Next == nil {
		Cursor = nil
		return header
	}

	if n <= 1 {
		Cursor = header.Next
		return header
	}

	last := RecurseReverseListNodeN(header.Next, n-1)
	header.Next.Next = header
	header.Next = Cursor
	return last
}

// ReverseListNodeN 非递归反转链表前n个节点
func ReverseListNodeN(header *base.ListNode, n int) *base.ListNode {
	if header == nil || header.Next == nil || n <= 1 {
		return header
	}

	cursor := header
	p := header
	i := 0

	for next := header.Next; i < n-1 && next != nil; i++ {
		cursor = next.Next
		next.Next = p
		p = next
		next = cursor
	}

	header.Next = cursor
	return p
}

// RecurseReverseListNodeBetween 递归反转链表的第m个至第n个节点
func RecurseReverseListNodeBetween(header *base.ListNode, m, n int) *base.ListNode {
	if header == nil || m >= n {
		return header
	}

	if m == 1 {
		return RecurseReverseListNodeN(header, n)
	}

	header.Next = RecurseReverseListNodeBetween(header.Next, m-1, n-1)
	return header
}

// ReverseListNodeBetween 非递归反转链表中的第m至第n个节点
func ReverseListNodeBetween(header *base.ListNode, m, n int) *base.ListNode {
	if header == nil || header.Next == nil || m >= n {
		return header
	}

	dummyNode := &base.ListNode{
		Val:  -1,
		Next: header,
	}

	prev := dummyNode
	p := prev.Next
	for i := 0; i < m-1 && p != nil; i++ {
		prev = prev.Next
		p = p.Next
	}

	// dum    1,   2,   3,   4,  5
	// prev   p   cur
	for i := 0; i < n-m; i++ {
		curr := p.Next
		p.Next = curr.Next
		curr.Next = prev.Next
		prev.Next = curr
	}

	return dummyNode.Next
}

// ReverseBetweenV2  反转指定起始节点之间的节点
func ReverseBetweenV2(head, start, end *base.ListNode) *base.ListNode {
	if head == nil || head.Next == nil || start == nil {
		return head
	}
	// 定义一个假的头结点, 方便解决 start == head 的情况
	dummyNode := &base.ListNode{
		Val:  -1,
		Next: head,
	}
	// prev 保存start 的前一个节点
	prev := dummyNode
	// 移动prev, 使之指向start 的前一个节点
	for prev.Next != start {
		prev = prev.Next
	}
	// 反转 [start, end)区间内的节点
	//  1  ->  2   ->  3  ->  4  ->  5  ->  6  ->  1, 3, 2, 4, 5, 6
	//  prev  start          end
	for prev.Next != end && start.Next != nil {
		removed := start.Next
		start.Next = removed.Next
		removed.Next = prev.Next
		prev.Next = removed
	}
	return dummyNode.Next
}

// RecurseReverseKGroup k 个一组反转链表, 递归算法
func RecurseReverseKGroup(head *base.ListNode, k int) *base.ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	start := head
	end := head
	i := 0
	for ; i < k-1; i++ {
		// 不够k 个 直接返回
		if end == nil {
			return head
		}
		end = end.Next
	}
	// 说明剩余的不足k 个, 则无需反转

	if i != k-1 || end == nil {
		return head
	}
	fmt.Println("end i:", i)
	fmt.Println("start: ", start.Val)
	// 因为反转之后，end的next 不在指向之前的节点, 所以需要先保存之前的节点，以便下一次调用resverseKGroup
	next := end.Next
	fmt.Println("end0:", start.Val)
	// 1 2 3
	newheader := ReverseBetweenV2(head, start, end)
	base.OuputListNode(newheader)
	// time.Sleep(time.Second * 2)
	// fmt.Println("newheader: ", end.Val)
	start.Next = RecurseReverseKGroup(next, k)
	return newheader
}
