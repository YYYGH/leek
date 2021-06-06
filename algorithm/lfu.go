package algorithm

import "sync"

/*
LFU 策略，淘汰使用频率最低的数据
1. get, 获取 key, 存在则返回val, 不存在返回-1
2. 调用 get or put, 对应freq(频率)+1
3. put, 容量没满直接插入，容量满了之后删除使用频率最低的元素
如果最小的使用频率对应多个key 就删除最旧的key
解决思路：
需要在O(1) 时间内进行以上操作
1. 使用 Map (key, val)可以快速解决 get(key)
var KeyToVal map[int]int
2. 使用 Map(key, freq)可以快速key对应的freq
var KeyToFreq map[int]int
3.
3.1 需要key->freq 映射， 需要找到最小的freq
3.2 删除最小的freq 对应的key, 需要找到系统的中的最下的freq, 需要快速找到最小的freq 可以用 minfreq 记录最小的freq
3.3 freq 可能对应多个key, freq->list(key)
3.4 freq 对应的key 列表需要按时序存储，能够快速删除最旧的
3.5 希望能够快速删除key列表中的key, 频率为freq 的key被访问一次后，频率+1，所以key应该保存在 freq+1 对应的列表中
var FreqToKeys map[int][]int
*/
type LFUCache struct {
	KeyToVal  map[int]int
	KeyToFreq map[int]int
	// 最新的key添加到slice末尾, 最旧的为第一个元素
	FreqToKeys map[int][]int
	MinFrq     int
	Capacity   int
	sync.Mutex
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		KeyToVal:   make(map[int]int),
		KeyToFreq:  make(map[int]int),
		FreqToKeys: make(map[int][]int),
		Capacity:   capacity,
	}
}

func (p *LFUCache) Get(key int) int {
	p.Lock()
	defer p.Unlock()
	if _, ok := p.KeyToVal[key]; !ok {
		return -1
	}
	// 从之前freq对应的key列表中删除key
	p.delete(key, p.KeyToFreq[key])
	// 将 key 添加到新的freq 对应的列表中
	p.increaseFreq(key)
	return p.KeyToVal[key]
}

func (p *LFUCache) Put(key int, value int) {
	p.Lock()
	defer p.Unlock()
	if _, ok := p.KeyToVal[key]; ok {
		p.KeyToVal[key] = value
		// 从freq对应的列表中找到key 并删除
		p.delete(key, p.KeyToFreq[key])
		p.increaseFreq(key)
		return
	}
	// 如果容量满了
	if len(p.KeyToVal) >= p.Capacity {
		// 执行lfu 淘汰, 删除使用频率最低的
		// 拿到要被淘汰的key
		delKey := p.FreqToKeys[p.MinFrq][0]
		// 从freq 对应的列表中删除下标为0的元素, 因为新元素放在切片末尾
		p.FreqToKeys[p.MinFrq] = p.FreqToKeys[p.MinFrq][1:]
		// 删除之后列表为空，则删除freq 这个key
		if len(p.FreqToKeys[p.MinFrq]) == 0 {
			delete(p.FreqToKeys, p.MinFrq)
		}
		delete(p.KeyToVal, delKey)
		delete(p.KeyToFreq, delKey)
	}
	// 更新key->val 表
	p.KeyToVal[key] = value
	// 将 key 添加到新的freq 对应的列表中
	p.increaseFreq(key)
	p.MinFrq = 1
}

// 删除指定freq对应列表中的指定key, 在访问key 操作时用到
func (p *LFUCache) delete(key, freq int) {
	index := 0
	for ; index < len(p.FreqToKeys[freq]); index++ {
		if p.FreqToKeys[freq][index] == key {
			break
		}
	}
	p.FreqToKeys[freq] = append(p.FreqToKeys[freq][:index], p.FreqToKeys[freq][index+1:]...)
}

//更新使用频率
func (p *LFUCache) increaseFreq(key int) {
	p.KeyToFreq[key]++
	newFreq := p.KeyToFreq[key]
	// 更新频率
	if v, ok := p.FreqToKeys[newFreq]; !ok || v == nil {
		p.FreqToKeys[newFreq] = make([]int, 0)
	}
	// 添加key到prevFreq 对应的列表中
	p.FreqToKeys[newFreq] = append(p.FreqToKeys[newFreq], key)
}
