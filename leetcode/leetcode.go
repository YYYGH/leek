package leetcode

import (
	"fmt"
	"leek/base"
)

// MinWindow 76题.最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
func MinWindow(s, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	needCnt := len(t)

	for _, v := range []byte(t) {
		need[v]++
	}

	left, right := 0, 0
	var result string

	fmt.Println(need)
	for ; right < len(s); right++ {
		fmt.Println("loop: ", needCnt, left, right, need)
		// 如果需要的元素个数大于0, 则说明找到了一个需要的元素, 总需要的数量-1
		if need[s[right]] > 0 {
			needCnt--
		}
		need[s[right]]--

		if needCnt == 0 {
			// 移动左下标
			// continue
			fmt.Println("s:", s)
			fmt.Println("t:", t)
			fmt.Println("all", left, right, need, needCnt)
			// 跳过不需要的字符
			for need[s[left]] != 0 {
				fmt.Println("xx", left, right, need, needCnt)
				need[s[left]]++
				left++
			}

			fmt.Println("xx2", left, right, need, needCnt)
			if len(result) == 0 || right-left < len(result) {
				result = s[left : right+1]
			}

			fmt.Println("compare:", left, right, string(result), needCnt)
			need[s[left]]++
			needCnt++
			left++
		}
	}

	return result
}

// MinPathSum 64题. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
func MinPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = base.Min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

//
func FindDiagonalOrder(mat [][]int) []int {
	count := len(mat) * len(mat[0])
	result := make([]int, count)
	x := 0
	y := 0
	xturn := false
	yturn := false
	status := false
	for index := 0; index < count; index++ {
		fmt.Println(index, mat[x][y], x, y, xturn, yturn, status)
		result[index] = mat[x][y]
		// 第一行和最后一行
		// step := 1
		if (x == 0 || (x == len(mat)-1 && y != 0)) && !xturn {
			fmt.Println("00000")
			if y != len(mat[0])-1 {
				y++
			} else {
				// 移动到了最后一列, 此时x 坐标要增加
				x++
			}
			xturn = true
			if x == 0 {
				status = true
			} else {
				status = false
			}
		} else if (y == 0 || (y == len(mat[0])-1 && x != len(mat)-1)) && !yturn && status {
			// 第一列和最后一列
			fmt.Println("11111")
			if x != len(mat)-1 {
				x++
			} else {
				y++
			}
			yturn = true
			if y == 0 {
				status = false
			} else {
				status = true
			}
		} else {
			if status {
				fmt.Println("x...")
				x++
				y--
				if y == 0 || x == len(mat)-1 {
					xturn = false
					status = false
				} else {
					yturn = true
				}
			} else {
				fmt.Println("y...")
				x--
				y++
				if x == 0 || y == len(mat[0])-1 {
					yturn = false
					status = true
				} else {
					xturn = true
				}
			}
		}
	}
	return result
}
