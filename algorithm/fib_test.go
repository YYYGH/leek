package algorithm

import "testing"

var (
	expected = map[int]int{
		-1: 0,
		0:  0,
		1:  1,
		2:  1,
		5:  5,
		7:  13,
	}
)

func TestFibonacciRecurse(t *testing.T) {
	for k, v := range expected {
		actual := FibonacciRecurse(k)
		if actual != v {
			t.Errorf("Fib(%d) = %d; expected %d", k, actual, v)
		}
	}
}

func TestFibonacciRecurseV2(t *testing.T) {
	for k, v := range expected {
		actual := FibonacciRecurseV2(k)
		if actual != v {
			t.Errorf("Fib(%d) = %d; expected %d", k, actual, v)
		}
	}
}
func TestFibonacciForLoop(t *testing.T) {
	for k, v := range expected {
		actual := FibonacciForLoop(k)
		if actual != v {
			t.Errorf("Fib(%d) = %d; expected %d", k, actual, v)
		}
	}
}
func TestFibonacciForLoopV2(t *testing.T) {
	for k, v := range expected {
		actual := FibonacciForLoopV2(k)
		if actual != v {
			t.Errorf("Fib(%d) = %d; expected %d", k, actual, v)
		}
	}
}
