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

func TestSearchItemRange(t *testing.T) {
	data := map[int][][]int{
		8: {{3, 4}, {5, 7, 7, 8, 8, 10}},
		6: {{-1, -1}, {5, 7, 7, 8, 8, 10}},
		0: {{-1, -1}, {}},
		2: {{0, 1}, {2, 2}},
		7: {{4, 6}, {0, 1, 2, 5, 7, 7, 7, 8, 9, 10}},
	}

	for k, v := range data {
		actual := SearchItemRange(v[1], k)
		if actual[0] != v[0][0] || actual[1] != v[0][1] {
			t.Errorf("input: %v, target:%d, expected %v, got %v", v[1], k, v[0], actual)
		}
	}
}
