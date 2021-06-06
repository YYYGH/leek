package backtrack

import "sort"

/*
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]
*/
var res [][]int

func PermuteV1(nums []int) [][]int {
	res = make([][]int, 0)
	record := make([]int, 0)
	subPermute(nums, record)
	return res
}

func subPermute(nums []int, record []int) {
	if len(nums) == len(record) {
		res = append(res, append([]int{}, record...))
		return
	}

	for _, v := range nums {
		// 选过的不能选
		if hasExists(record, v) {
			continue
		}
		record = append(record, v)
		subPermute(nums, record)
		record = record[0 : len(record)-1]
	}
}

func hasExists(list []int, val int) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

/*
1. 对nums 进行排序
2.
*/
func PermuteV2(nums []int) [][]int {
	res := make([][]int, 0)
	var backTracking func([]int)
	sort.Ints(nums)
	tempRecord := make([]int, nums[len(nums)-1]+1)
	backTracking = func(temp []int) {
		if len(temp) == len(nums) {
			res = append(res, append([]int{}, temp...))
		}
		for i := 0; i < len(nums); i++ {
			if tempRecord[nums[i]] != 0 {
				continue
			}

			temp = append(temp, nums[i])
			tempRecord[nums[i]] = 1
			backTracking(temp)
			tempRecord[nums[i]] = 0
			temp = temp[:len(temp)-1]

		}
	}
	backTracking([]int{})
	return res
}
