package bfs

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
