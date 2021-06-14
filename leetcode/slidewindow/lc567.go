package slidewindow

/*
567. 字符串的排列
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
示例 1：

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").
示例 2：

输入: s1= "ab" s2 = "eidboaoo"
输出: False
*/
func CheckInclusion(s, t string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, v := range []byte(t) {
		need[v]++
	}
	valid := 0
	left, right := 0, 0

	bs := []byte(s)
	for ; right < len(s); right++ {
		if v, ok := need[bs[right]]; ok {
			window[bs[right]]++
			if v == window[bs[right]] {
				valid++
			}
		}
		for valid == len(need) {
			if (right - left + 1) == len(need) {
				return true
			}
			if v, ok := need[bs[left]]; ok {
				if v == window[bs[left]] {
					valid--
				}
				window[bs[left]]--
			}
			left++
		}
	}
	return false
}

func CheckInclusionV2(s, t string) bool {
	need := make(map[byte]int)
	valid := 0
	for _, v := range []byte(t) {
		need[v]++
		valid++
	}
	left, right := 0, 0
	bs := []byte(s)
	for ; right < len(s); right++ {
		if _, ok := need[bs[right]]; ok {
			need[bs[right]]--
			valid--
		}

		for ; valid == 0; left++ {
			if (right - left + 1) == len(t) {
				return true
			}
			if _, ok := need[bs[left]]; ok {
				need[bs[left]]++
				valid++
			}
		}
	}
	return false
}
