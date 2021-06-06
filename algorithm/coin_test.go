package algorithm

import "testing"

var (
	Coinexpected = map[int][][]int{
		// 3: {{1, 2, 5}, {11}},
		// 20: {{186, 419, 83, 408}, {6249}},
		25: {{3, 7, 405, 436}, {8839}},
		24: {{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, {9864}},
	}
)

func TestCoinChangeV1(t *testing.T) {
	for k, v := range Coinexpected {
		actual := CoinChangeV1(v[0], v[1][0])
		if actual != k {
			t.Errorf("Coin(%d) = %d; expected %d", k, actual, v)
		}
	}
}

func TestCoinChangeV2(t *testing.T) {
	for k, v := range Coinexpected {
		actual := CoinChangeV2(v[0], v[1][0])
		if actual != k {
			t.Errorf("Coin(%d) = %d; expected %d", k, actual, v)
		}
	}
}

func TestCoinChangeV4(t *testing.T) {
	for k, v := range Coinexpected {
		actual := CoinChangeV4(v[0], v[1][0])
		if actual != k {
			t.Errorf("Coin(%d) = %d; expected %d", k, actual, v)
		}
	}
}
