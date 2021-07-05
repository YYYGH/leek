package bfs

/*
752. 打开转盘锁
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。

锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

示例 1:

输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。
示例 2:

输入: deadends = ["8888"], target = "0009"
输出：1
解释：
把最后一位反向旋转一次即可 "0000" -> "0009"。
示例 3:

输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。
示例 4:

输入: deadends = ["0000"], target = "8888"
输出：-1
*/
func OpenLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	deadmap := make(map[string]bool)
	for _, v := range deadends {
		deadmap[v] = true
	}
	// 记录走过的路
	visited := make(map[string]bool)
	visited["0000"] = true
	// 带选择的元素队列
	queue := make([]string, 0)
	queue = append(queue, "0000")

	step := 0
	for len(queue) != 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			curr := queue[i]
			if _, ok := deadmap[curr]; ok {
				continue
			}
			// 找到目标
			if curr == target {
				return step
			}
			dst := []byte(curr)
			// 尝试转动每一个拨轮
			for j := 0; j < len(dst); j++ {
				value := dst[j]
				// 向上转动
				dst[j] = up(dst[j])
				v := string(dst)
				if _, ok := visited[v]; !ok {
					visited[v] = true
					queue = append(queue, v)
				}
				dst[j] = value

				dst[j] = down(dst[j])
				v = string(dst)
				if _, ok := visited[v]; !ok {
					visited[v] = true
					queue = append(queue, v)
				}
				dst[j] = value
			}
		}
		queue = queue[count:]
		step++
	}
	// 找不到答案
	return -1
}

func up(cur byte) byte {
	if cur == '9' {
		return '0'
	}
	return cur + 1
}

func down(cur byte) byte {
	if cur == '0' {
		return '9'
	}
	return cur - 1
}

func OpenLockV2(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	deadmap := make(map[string]bool)
	for _, v := range deadends {
		deadmap[v] = true
	}
	// 记录走过的路
	visited := make(map[string]bool)

	visited1 := make(map[string]bool)
	visited1["0000"] = true

	visited2 := make(map[string]bool)
	visited2[target] = true

	step := 0
	for len(visited1) != 0 && len(visited2) != 0 {
		tmp := make(map[string]bool)
		for key, _ := range visited1 {
			if _, ok := deadmap[key]; ok {
				continue
			}
			// 找到目标
			if _, ok := visited2[key]; ok {
				return step
			}

			visited[key] = true
			dst := []byte(key)
			// 尝试转动每一个拨轮
			for j := 0; j < len(dst); j++ {
				value := dst[j]
				// 向上转动
				dst[j] = up(dst[j])
				v := string(dst)
				if _, ok := visited[v]; !ok {
					tmp[v] = true
				}
				dst[j] = value

				dst[j] = down(dst[j])
				v = string(dst)
				if _, ok := visited[v]; !ok {
					tmp[v] = true
				}
				dst[j] = value
			}
		}
		visited1 = visited2
		visited2 = tmp
		step++
	}
	// 找不到答案
	return -1
}
