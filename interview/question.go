package interview

import (
	"fmt"
	"leek/base"
	"sort"
	"strconv"
)

// MinPathWay xx
/*
[
   //0  1  2  3  4  5  6
  0  1  2  1  1  1  1  1
  1  1  1  1  1  3  D  1
  2  1  1  #  1  1  #  1
  3  1  1  1  #  1  1  1
  4  1  S  1  1  2  1  1
  5  1  1  1  2  1  1  1
  6  1  1  1  1  1  1  1
]
s[4, 1]
d[1, 5]
3 + 4
在一快2D平面上，求从一个点移动到另一个点的最快移动路径。
注意：起始点随机，路况随机，有的路无法通过，有点路是普通路时间的两倍或三倍。
格子中的数字表示通过该点的耗时, # 表示无法通过
S点为起始位置, D点为目标位置
*/
func MinPathWay(grids [][]int, start, end []int) [][]int {
	preGrids := make([][]int, 0)
	for _, row := range grids {
		arry := make([]int, len(row), len(row))
		for index := range row {
			arry[index] = row[index]
		}
		preGrids = append(preGrids, arry)
	}

	xstep := 1
	if end[0]-start[0] < 0 {
		xstep = -1
	}
	ystep := 1
	if end[1]-start[1] < 0 {
		ystep = -1
	}

	xindex := start[0]
	for i := 0; i <= base.Abs(start[0]-end[0]); i++ {
		yindex := start[1]
		for j := 0; j <= base.Abs(start[1]-end[1]); j++ {
			// 起始位置
			if xindex == start[0] && yindex == start[1] {
			} else if xindex == start[0] && yindex != start[1] {
				// 假设-1 是路不通的
				grids[xindex][yindex] = grids[xindex][yindex-ystep] + grids[xindex][yindex]
			} else if xindex != start[0] && yindex == start[1] {
				grids[xindex][yindex] = grids[xindex-xstep][yindex] + grids[xindex][yindex]
			} else {
				grids[xindex][yindex] = base.Min(grids[xindex-xstep][yindex], grids[xindex][yindex-ystep]) + grids[xindex][yindex]
			}
			yindex += ystep
		}
		xindex += xstep
	}

	for _, v := range grids {
		fmt.Println(v)
	}
	fmt.Println("customer_time: ", grids[end[0]][end[1]])
	return findPath(preGrids, grids, start, end)
}

func findPath(before, after [][]int, start, end []int) [][]int {
	xstep := 1
	if end[0]-start[0] < 0 {
		xstep = -1
	}
	ystep := 1
	if end[1]-start[1] < 0 {
		ystep = -1
	}
	path := make([][]int, 0)
	xindex := end[0]
	yindex := end[1]
	for i := 0; i <= base.Abs(start[0]-end[0]); i++ {
		for j := 0; j <= base.Abs(start[1]-end[1]); j++ {
			// 当前位置的值等于 原始数组位置的值加上x坐标移动的值
			path = append(path, []int{xindex, yindex})
			if ystep > 0 {
				if (yindex - ystep) < start[1] {
					break
				}
			} else {
				if (yindex - ystep) > start[1] {
					break
				}
			}
			if before[xindex][yindex] == (after[xindex][yindex] - after[xindex][yindex-ystep]) {
				yindex -= ystep
			} else {
				break
			}
		}
		xindex -= xstep
	}

	for i := 0; i < len(path)/2; i++ {
		temp := path[i]
		path[i] = path[len(path)-i-1]
		path[len(path)-i-1] = temp
	}
	return path
}

// AddNumString 字符串数字相加
func AddNumString(num1, num2 string) string {
	length1 := len(num1)
	length2 := len(num2)
	n := length1
	c := length1 - length2
	max := true
	if n < length2 {
		n = length2
		c *= -1
		max = false
	}
	result := ""
	x := uint64(0)
	temp := uint64(0)
	for i := n - 1; i >= 0; i-- {
		sum := uint64(0)
		if max {
			if i-c >= 0 {
				v, _ := strconv.ParseUint(string(num2[i-c]), 10, 64)
				sum += v
			}
			v, _ := strconv.ParseUint(string(num1[i]), 10, 64)
			sum += v
		} else {
			if i-c >= 0 {
				v := num1[i-c] - '0'
				fmt.Println("v: ", v)
				// v, _ := strconv.ParseUint(string(num1[i-c]), 10, 64)
				k, err := strconv.Atoi(fmt.Sprintf("%v", v))
				if err != nil {
					fmt.Println("err: ", err.Error())
				}
				fmt.Println("v: ", v, k)
				sum += uint64(k)
			}
			v, _ := strconv.ParseUint(string(num2[i]), 10, 64)
			sum += v
		}

		// num1 :=   "8 3 2 1"
		// num2 := "1 3 2 2 9"
		// fmt.Println(i, x, sum)
		temp = (x + sum) / 10
		result += fmt.Sprintf("%d", (x+sum)%10)
		x = temp
		// fmt.Println(i, x, result)
	}

	if x == 1 {
		result += "1"
	}

	s := ""
	for i := len(result) - 1; i >= 0; i-- {
		s += string(result[i])
	}

	return s
}

func exists(list []int, key int) (bool, int) {
	for index, v := range list {
		if key == v {
			return true, index
		}
	}
	return false, -1
}

// CheckShuzi 判断手中的牌是不是顺子, 其中大小王可以代表任意牌
func CheckShuzi(list []int) bool {
	// xiaowang = 14
	// dawang = 15
	common := 0
	if ok, index := exists(list, 14); ok {
		// list = append(list[:index], list[index+1:]...)
		list = list[:index+copy(list[index:], list[index+1:])]
		// a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
		common++
	}
	if ok, index := exists(list, 15); ok {
		list = list[:index+copy(list[index:], list[index+1:])]
		common++
	}

	sort.Ints(list)
	status := true
	for i := len(list) - 1; i > 0; i-- {
		// 大于2 则无法组成顺子
		if list[i]-list[i-1] > 3 {
			fmt.Printf("first break\n")
			// 大于 3
			// 2 . . . 6   6-2 = 4 > 3
			status = false
			break
		} else if list[i]-list[i-1] == 3 {
			fmt.Printf("value:%d , common:%d\n", 3, common)
			// 等于3 需要判断还有几张大小王， 需要判断大小王是不是至少还有2张， 有则可以替换，否则无法替换
			// 2 . . 5   5-2 = 3  差两个填空
			if common == 2 {
				common -= 2
			} else {
				status = false
				break
			}
		} else if list[i]-list[i-1] == 2 {
			fmt.Printf("value:%d , common:%d\n", 2, common)
			// 等于2 需要判断还有几张大小王， 需要判断大小王是不是至少还有一张， 有则可以替换，否则无法替换
			// 2 . 4, 5  4-2 = 2  差一张填空
			if common > 0 {
				common--
			} else {
				status = false
				break
			}
		} else if list[i]-list[i-1] == 1 { // 连续的
			fmt.Printf("continue\n")
			continue
		} else { // 等于0 则，有相同的， 直接跳出
			fmt.Printf("else break\n")
			status = false
			break
		}
	}

	return status
}

// SearchInsert 二分查找一个给定值在给定序列中插入位置
func SearchInsert(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1
	for leftIndex <= rightIndex {
		fmt.Printf("left:[%d], right:[%d]\n", leftIndex, rightIndex)

		tmpIndex := (rightIndex + leftIndex) / 2
		if nums[tmpIndex] == target {
			leftIndex = tmpIndex
			break
		} else if nums[tmpIndex] > target {
			// 说明在中间值的左边
			rightIndex = tmpIndex - 1
		} else {
			leftIndex = tmpIndex + 1
		}
	}
	return leftIndex
}
