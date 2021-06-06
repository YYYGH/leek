package algorithm

import "testing"

func TestLFU(t *testing.T) {
	lfu := NewLFUCache(3)
	lfu.Put(1, 2)
	lfu.Put(2, 3)
	lfu.Put(3, 4)
	if len(lfu.KeyToVal) != 3 {
		t.Errorf("LFUSize: %d; expected %d", len(lfu.KeyToVal), 3)
	}

	r := lfu.Get(3)
	if r != 4 {
		t.Errorf("key: %d; expected %d, actual:%d", 3, 4, r)
	}
	lfu.Put(4, 5)
	// 已经被淘汰了
	r = lfu.Get(1)
	if r != -1 {
		t.Errorf("key: %d; expected %d, actual:%d", 3, 4, r)
	}
	lfu.Get(2)
	lfu.Put(5, 6)
	// 已经被淘汰了
	r = lfu.Get(4)
	if r != -1 {
		t.Errorf("key: %d; expected %d, actual:%d", 3, 4, r)
	}
}
