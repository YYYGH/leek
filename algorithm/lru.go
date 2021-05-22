package algorithm

type DoubleList interface {
	// 在链表头部添加节点p, 时间0(1)
	AddFirst(*Node)
	// 删除链表p节点(p 一定存在)
	// 由于是双向链表, 给的目标Node节点, 时间O(1)
	Remove(*Node)
	// 删除链表中的最后一个节点吗, 并返回该节点，时间O(1)
	RemoveLast() int
}

type LRUCache struct {
	Map   map[int]*Node
	Cache *DoubleLinkedNode
	Cap   int
}

type LRU interface {
	Get(key int) int
	Put(key, value int)
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Map:   make(map[int]*Node),
		Cache: NewDoubleLinkedNode(),
		Cap:   cap,
	}
}

// 在链表头部添加节点p, 时间0(1)
func (p *DoubleLinkedNode) AddFirst(n *Node) {
	n.Next = p.Header.Next
	n.Prev = p.Header
	n.Next.Prev = n
	p.Header.Next = n
}

// 删除链表p节点(p 一定存在)
// 由于是双向链表, 给的目标Node节点, 时间O(1)
func (p *DoubleLinkedNode) Remove(n *Node) int {
	key := n.Val
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	return key
}

// 删除链表中的最后一个节点吗, 并返回该节点，时间O(1)
func (p *DoubleLinkedNode) RemoveLast() int {
	if p.Header.Next == p.Tail {
		return -1
	}
	return p.Remove(p.Tail.Prev)
}

func (p *LRUCache) Get(key int) int {
	// 找不到
	if v, ok := p.Map[key]; ok {
		// 放在链表表头，保证是最近访问过的
		p.Put(v.Key, v.Val)
		return v.Val
	}
	return -1
}

func (p *LRUCache) Put(key, value int) {
	node := NewNode(key, value)
	if v, ok := p.Map[key]; ok {
		// 删除节点
		p.Cache.Remove(v)
	} else {
		// 容量不足,  删除最近最少使用的
		if p.Cap == len(p.Map) {
			k := p.Cache.RemoveLast()
			delete(p.Map, k)
		}
	}
	p.Map[key] = node
	p.Cache.AddFirst(node)
}
