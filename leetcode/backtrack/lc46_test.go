package backtrack

import "testing"

func BenchmarkPermuteV1(b *testing.B) {
	array := []int{1, 2, 3, 4, 5}
	PermuteV1(array)
}

func TestPermuteV2(b *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	PermuteV2(array)
}
