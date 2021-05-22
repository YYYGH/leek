package base

import (
	"errors"
	"reflect"
)

// Min compre a and b return min value
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max compre a and b return max value
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abs 求绝对值
func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// Reverse 反转slice
func Reverse(s []interface{}) []interface{} {
	a := make([]interface{}, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func ReverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}
	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}
