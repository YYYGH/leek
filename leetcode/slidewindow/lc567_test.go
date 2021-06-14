package slidewindow

import "testing"

var datav1 = map[bool][]string{
	true:  {"ab", "eidbaooo"},
	false: {"ab", "eidboaoo"},
}

func TestCheckInclusion(t *testing.T) {
	for k, v := range datav1 {
		actual := CheckInclusion(v[1], v[0])
		if actual != k {
			t.Errorf("input: %v, expected: %v, actual: %v", v, k, actual)
		}
	}
}

func TestCheckInclusionV2(t *testing.T) {
	for k, v := range datav1 {
		actual := CheckInclusionV2(v[1], v[0])
		if actual != k {
			t.Errorf("input: %v, expected: %v, actual: %v", v, k, actual)
		}
	}
}
