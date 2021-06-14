package array

/*
1. 找出数组中的重复数字
   在一个长度为n的数组中，所有数字都在0~n-1的范围内。数组中的某些数字是重复的，
但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的
数字。
例如：如果输入长度为7的数组 [2, 3, 1, 0, 2, 5, 3], 那么对应的输出是重复数字2或者3
*/

// 使用map
// 时间： O(n), 空间: O(n)
func FindRepeatItem(array []int) int {
	record := make([]int, len(array))
	for _, v := range array {
		if v >= len(array) {
			return -1
		}
		record[v]++
		if record[v] != 1 {
			return v
		}
	}
	return -1
}

// 元素交换, 修改数组
// 时间： O(n), 空间: O(1)
func FindRepeatItemV2(array []int) int {
	// {{2, 3}, {2, 3, 1, 0, 2, 5, 3}},
	for index := 0; index < len(array); index++ {
		for array[index] != index {
			if array[index] >= len(array) {
				return -1
			}
			// 下标index 和 v 都出现了数字v
			if array[index] == array[array[index]] {
				return array[index]
			}
			tmp := array[index]
			array[index] = array[tmp]
			array[tmp] = tmp
		}
	}
	return -1
}

// 不修改数组
// 类似于二分查找 时间：O(nlogn), 空间 O(1)
func FindRepeatItemV3(array []int) int {
	left := 1
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		count := Counter(array, left, mid)
		if left == right {
			if count > 1 {
				if array[left-1] <= 0 || array[left-1] > len(array) {
					return -1
				}
				return array[left]
			}
			break
		}
		if count > (mid - left + 1) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

func Counter(array []int, left, right int) (c int) {
	for _, v := range array {
		if v >= left && v <= right {
			c++
		}
	}
	return
}
