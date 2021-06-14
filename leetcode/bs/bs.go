package bs

func BinarySearch(list []int, target int) int {
	left := 0
	right := len(list) - 1
	// [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == target {
			return mid
		} else if list[mid] < target {
			left = mid + 1
		} else if list[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchV2(list []int, target int) int {
	left := 0
	right := len(list)
	// [left, right)
	for left < right {
		mid := left + (right-left)/2
		if list[mid] == target {
			return mid
		} else if list[mid] < target {
			left = mid + 1
		} else if list[mid] > target {
			right = mid
		}
	}
	return -1
}

// 查找左边界
func BinarySearchLeftRound(list []int, target int) int {
	left := 0
	right := len(list) - 1
	// [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == target {
			right = mid - 1
		} else if list[mid] < target {
			left = mid + 1
		} else if list[mid] > target {
			right = mid - 1
		}
	}
	if left >= len(list) || list[left] != target {
		return -1
	}
	return left
}

func BinarySearchLeftRoundV2(list []int, target int) int {
	left := 0
	right := len(list)
	// [left, right)
	for left < right {
		mid := left + (right-left)/2
		if list[mid] == target {
			right = mid
		} else if list[mid] < target {
			left = mid + 1
		} else if list[mid] > target {
			right = mid
		}
	}
	if left >= len(list) || list[left] != target {
		return -1
	}
	return left
}

func BinarySearchRightRound(list []int, target int) int {
	left := 0
	right := len(list) - 1
	// [left, right]
	for left <= right {
		mid := (right-left)>>1 + left
		if list[mid] == target {
			left = mid + 1
		} else if list[mid] < target {
			left = mid + 1
		} else if list[mid] > target {
			right = mid - 1
		}
	}
	if right < 0 || list[right] != target {
		return -1
	}
	return right
}

func BinarySearchRightRoundV2(list []int, target int) int {
	left := 0
	right := len(list)
	// [left, right)
	for left < right {
		mid := (right-left)>>1 + left
		if list[mid] == target {
			left = mid + 1
		} else if list[mid] < target {
			left = mid + 1
		} else if list[mid] > target {
			right = mid
		}
	}
	if left == 0 || list[right-1] != target {
		return -1
	}
	return right - 1
}
