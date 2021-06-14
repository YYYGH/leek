package array

import "testing"

var (
	data = [][][]int{
		{{2, 3}, {2, 3, 1, 0, 2, 5, 3}},
		{{-1}, {2, 4, 1, 0, 3, 5, 6}},
		{{-1}, {2, 4, 1, 0, 3, 8, 6}},
	}
)

func TestFindRepeatItem(t *testing.T) {
	for _, v := range data {
		actual := FindRepeatItem(v[1])
		ok := false
		for _, expect := range v[0] {
			if actual == expect {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("expect: %v, actual: %v, input: %v", v[0], actual, v[1])
		}
	}
}

func TestFindRepeatItemV2(t *testing.T) {
	for _, v := range data {
		actual := FindRepeatItemV2(v[1])
		ok := false
		for _, expect := range v[0] {
			if actual == expect {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("expect: %v, actual: %v, input: %v", v[0], actual, v[1])
		}
	}
}

func TestFindRepeatItemV3(t *testing.T) {
	D := [][][]int{
		{{2, 3}, {2, 3, 1, 4, 2, 5, 3}},
		{{-1}, {2, 4, 1, 3, 5, 6}},
		{{-1}, {2, 4, 1, 0, 3, 8, 6}},
		{{2, 3}, {2, 3, 5, 4, 3, 2, 6, 7}},
	}
	for _, v := range D {
		actual := FindRepeatItemV3(v[1])
		ok := false
		for _, expect := range v[0] {
			if actual == expect {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("expect: %v, actual: %v, input: %v", v[0], actual, v[1])
		}
	}
}
