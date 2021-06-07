package bfs

import (
	"testing"
)

func TestOpenLock(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	// ["8887","8889","8878","8898","8788","8988","7888","9888"] "8888"
	target := "0202"
	actual := OpenLock(deadends, target)
	if actual != 6 {
		t.Errorf("OpenLock(%v)= %v; expected %d, actual: %d", target, deadends, 6, actual)
	}
	deadends = []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target = "8888"
	actual = OpenLock(deadends, target)
	if actual != -1 {
		t.Errorf("OpenLock(%v)= %v; expected %d, actual: %d", target, deadends, -1, actual)
	}
}

func TestOpenLockV2(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	// ["8887","8889","8878","8898","8788","8988","7888","9888"] "8888"
	target := "0202"
	actual := OpenLockV2(deadends, target)
	if actual != 6 {
		t.Errorf("OpenLock(%v)= %v; expected %d, actual: %d", target, deadends, 6, actual)
	}
	deadends = []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target = "8888"
	actual = OpenLockV2(deadends, target)
	if actual != -1 {
		t.Errorf("OpenLock(%v)= %v; expected %d, actual: %d", target, deadends, -1, actual)
	}
}
