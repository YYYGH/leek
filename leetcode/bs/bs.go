package bs

/*
704. 二分查找
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
提示：

你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间
*/
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

// 二分查找左边界
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

// 二分查找左边界
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

// 二分查找右边界
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

// 二分查找右边界
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

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
进阶：
你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
提示：
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/
func SearchItemRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	// [0, len(nums))
	mid := 0
	result := []int{-1, -1}
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left > right {
		return result
	}
	// 找左边界
	result[0] = searchLeft(nums, target, left, mid)
	// 找右边界
	result[1] = searchRight(nums, target, mid+1, right)
	return result
}

func searchLeft(nums []int, target, left, right int) int {
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else { // 除了等于 就只有小于了, 即 nums[mid] < target
			left = mid + 1
		}
	}
	return right
}

func searchRight(nums []int, target, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else { // 除了等于 就只有大于了, 即 nums[mid] > target
			right = mid - 1
		}
	}
	return left - 1
}
