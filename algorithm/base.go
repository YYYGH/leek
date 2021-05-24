package algorithm

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

// lru
// DoubleLinkedNode 双向链表结构,
// 最近访问的Node 放到链表的头部的下一个位置, 尾部表示最久没被访问的
type DoubleLinkedNode struct {
	Header *Node
	Tail   *Node
}

func NewNode(key int, value int) *Node {
	return &Node{
		Key:  key,
		Val:  value,
		Next: nil,
		Prev: nil,
	}
}

func NewDoubleLinkedNode() *DoubleLinkedNode {
	header := NewNode(0, 0)
	tail := NewNode(0, 0)
	header.Next = tail
	tail.Prev = header
	return &DoubleLinkedNode{
		Header: header,
		Tail:   tail,
	}
}
