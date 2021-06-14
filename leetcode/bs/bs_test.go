package bs

import "testing"

var (
	dataBS = map[int][][]int{
		3: {{8}, {2, 3, 5, 8, 9}},
		2: {{3}, {1, 2, 3, 4, 5}},
		0: {{2}, {2, 3, 5, 8, 9}},
		// -1: {{6}, {2, 3, 5, 8, 9}},
		-1: {{10}, nil},
	}

	dataLeftRound = map[int][][]int{
		3: {{8}, {2, 3, 5, 8, 8, 8, 9, 10}},
		1: {{3}, {1, 3, 3, 4, 5}},
		0: {{2}, {2, 3, 5, 8, 9}},
		4: {{9}, {2, 3, 5, 8, 9}},
		// -1: {{6}, {2, 3, 5, 8, 9}},
		-1: {{10}, nil},
	}

	dataRightRound = map[int][][]int{
		5: {{8}, {2, 3, 5, 8, 8, 8, 9, 10}},
		2: {{3}, {1, 3, 3, 4, 5}},
		0: {{2}, {2, 3, 5, 8, 9}},
		4: {{9}, {2, 3, 5, 8, 9}},
		// -1: {{6}, {2, 3, 5, 8, 9}},
		// -1: {{10}, {2, 3, 5, 8, 9}},
		-1: {{10}, nil},
	}
)

func TestBinarySearch(t *testing.T) {
	for k, v := range dataBS {
		index := BinarySearch(v[1], v[0][0])
		if index != k {
			t.Errorf("expect: %d, actual: %d, s:%d, v: %v", k, index, v[0][0], v[1])
		}
	}
}

func TestBinarySearchV2(t *testing.T) {
	for k, v := range dataBS {
		index := BinarySearchV2(v[1], v[0][0])
		if index != k {
			t.Errorf("expect: %d, actual: %d, s:%d, v: %v", k, index, v[0][0], v[1])
		}
	}
}

func TestBinarySearchLeftRound(t *testing.T) {
	for k, v := range dataLeftRound {
		index := BinarySearchLeftRound(v[1], v[0][0])
		if index != k {
			t.Errorf("expect: %d, actual: %d, s:%d, v: %v", k, index, v[0][0], v[1])
		}
	}
}

func TestBinarySearchLeftRoundV2(t *testing.T) {
	for k, v := range dataLeftRound {
		index := BinarySearchLeftRoundV2(v[1], v[0][0])
		if index != k {
			t.Errorf("expect: %d, actual: %d, s:%d, v: %v", k, index, v[0][0], v[1])
		}
	}
}

func TestBinarySearchRightRound(t *testing.T) {
	for k, v := range dataRightRound {
		index := BinarySearchRightRound(v[1], v[0][0])
		if index != k {
			t.Errorf("expect: %d, actual: %d, s:%d, v: %v", k, index, v[0][0], v[1])
		}
	}
}
func TestBinarySearchRightRoundV2(t *testing.T) {
	for k, v := range dataRightRound {
		index := BinarySearchRightRoundV2(v[1], v[0][0])
		if index != k {
			t.Errorf("expect: %d, actual: %d, s:%d, v: %v", k, index, v[0][0], v[1])
		}
	}
}
