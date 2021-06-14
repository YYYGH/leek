package slidewindow

import "math"

/*
76. 最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：

输入：s = "a", t = "a"
输出："a"
提示：
1 <= s.length, t.length <= 105
s 和 t 由英文字母组成
*/

func MinWindow(s, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, v := range []byte(t) {
		need[v]++
	}
	valid := 0
	start, length := 0, math.MaxInt32
	left, right := 0, 0
	bs := []byte(s)
	for ; right < len(s); right++ {
		// 找到需要的字符
		if v, ok := need[bs[right]]; ok {
			// 在窗口中记录元素出现次数
			window[bs[right]]++
			// 窗口范围内的 当前字符出现次数等于 需要的次数, 说明这个字符已经找全了
			if v == window[bs[right]] {
				valid++
			}
		}
		// ADOBECODEBANC", "ABC"
		// 找到的有效字符数和need 长度相等
		for valid == len(need) {
			// 记录当前找到的子串位置
			if (right - left) < length {
				start = left
				length = right - left
			}
			// 从需要的字符记录中去掉left 下标对应的元素, left 向右移动
			if v, ok := need[bs[left]]; ok {
				if v == window[bs[left]] {
					valid--
				}
				// 窗口中记录的元素数量减1
				window[bs[left]]--
			}
			left++
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length+1]
}

func MinWindowV2(s, t string) string {
	need := make(map[byte]int)
	valid := 0
	for _, v := range []byte(t) {
		need[v]++
		valid++
	}
	// valid 初始化为需要的有效字符总数量
	left, right := 0, 0
	start, length := 0, math.MaxInt32
	bs := []byte(s)
	for ; right < len(s); right++ {
		// 找到一个需要的字符,
		if _, ok := need[bs[right]]; ok {
			// 字符出现的次数减 1
			need[bs[right]]--
			// 需要的有效字符总数 减 1
			valid--
		}
		// valid == 0 说明已经找到所有需要的字符, 这个时候可以执行left右移, 即窗口右移继续找
		for ; valid == 0; left++ {
			// 记录当前找到的子串位置
			if length > (right - left) {
				start = left
				length = right - left
			}
			// 当前字符是need 中需要的
			if _, ok := need[bs[left]]; ok {
				// 需要的字符总数 加 1, 因为left 右移, 少了一个字符
				valid++
				// 对应需要的字符个数加 1
				need[bs[left]]++
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length+1]
}
