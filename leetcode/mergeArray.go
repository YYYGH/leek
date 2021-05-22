package leetcode

import (
	"fmt"
	"leek/base"
	"sort"
)

type DataList [][]int

func (d DataList) Len() int           { return len(d) }
func (d DataList) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d DataList) Less(i, j int) bool { return d[i][0] < d[j][0] }

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。


提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/
// MergeArray 合并数组区间
func MergeArray(intervals [][]int) [][]int {
	data := DataList(intervals)
	result := make([][]int, 0)
	sort.Sort(data)
	fmt.Println(data)
	result = append(result, data[0])
	length := len(result)
	for i := 1; i < len(data); i++ {
		if result[length-1][1] >= data[i][0] {
			result[length-1][1] = base.Max(result[length-1][1], data[i][1])
		} else {
			result = append(result, data[i])
			length++
		}
	}
	return result
}
