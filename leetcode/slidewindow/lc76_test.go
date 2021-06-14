package slidewindow

import "testing"

func TestMinWindow(t *testing.T) {
	strmap := map[string][]string{
		"BANC": {"ADOBECODEBANC", "ABC"},
		"a":    {"a", "a"},
		"":     {"abFGhrGoP", "cpf"},
	}
	for k, v := range strmap {
		actual := MinWindow(v[0], v[1])
		if actual != k {
			t.Errorf("input:%v, expected:%s, got:%s", v, k, actual)
		}
	}
}

func TestMinWindowV2(t *testing.T) {
	strmap := map[string][]string{
		"BANC": {"ADOBECODEBANC", "ABC"},
		"a":    {"a", "a"},
		"":     {"abFGhrGoP", "cpf"},
	}
	for k, v := range strmap {
		actual := MinWindowV2(v[0], v[1])
		if actual != k {
			t.Errorf("input:%v, expected:%s, got:%s", v, k, actual)
		}
	}
}
