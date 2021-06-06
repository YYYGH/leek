package algorithm

import "fmt"

// 01冒泡， O(n^2) 稳定
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		swap := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j] = nums[j] + nums[j+1]
				nums[j+1] = nums[j] - nums[j+1]
				nums[j] = nums[j] - nums[j+1]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

// SelectSort
// 02选择排序 O(n^2)  不稳定
// 比如 10, 2, 34, 4, 6, 1, 2, 2, 9, for 循环中下标为8 的2 会被交换到前面
func SelectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[i] = nums[i] + nums[min]
			nums[min] = nums[i] - nums[min]
			nums[i] = nums[i] - nums[min]
		}
	}
}

// 03 插入排序  O(n^2)  稳定
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		index := i
		value := nums[index]
		for index > 0 && value < nums[index-1] {
			nums[index] = nums[index-1]
			index--
		}
		nums[index] = value
	}
}

// 04 归并排序 O(nlogn) 稳定
func MergeSort(nums []int) {
	temp := make([]int, len(nums))
	internalMergeSort(nums, temp, 0, len(nums)-1)
}

func internalMergeSort(nums, tmp []int, left, right int) {
	if left < right {
		middle := (left + right) / 2
		fmt.Println("internal: ", left, middle, right)
		internalMergeSort(nums, tmp, left, middle)          // 排序左子数组
		internalMergeSort(nums, tmp, middle+1, right)       // 排序右子数组
		mergeSortChildArray(nums, tmp, left, middle, right) // 合并两个子序列
	}
}

func mergeSortChildArray(nums, tmp []int, left, middle, right int) {
	i := left
	j := middle + 1
	k := 0
	fmt.Println(i, j, k, len(nums), len(tmp))
	for i <= middle && j <= right {
		minn := nums[i]
		if minn < nums[j] {
			i++
		} else {
			minn = nums[j]
			j++
		}
		tmp[k] = minn
		k++
	}

	for ; i <= middle; i++ {
		tmp[k] = nums[i]
		k++
	}

	for ; j <= right; j++ {
		tmp[k] = nums[j]
		k++
	}

	for i := 0; i < k; i++ {
		nums[left+i] = tmp[i]
	}
}

// 快速排序  O(nlogn) 不稳定
func QuickSort(nums []int) {
	qsort(nums, 0, len(nums)-1)
}

func qsort(nums []int, low, high int) {
	if low >= high {
		return
	}
	pivotIndex := partition(nums, low, high)
	qsort(nums, low, pivotIndex-1)
	qsort(nums, pivotIndex+1, high)
}

func partition(nums []int, low, high int) int {
	// 设置一个基准
	pivot := nums[low]
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}
		// 交换比基准大的数据到基准左边
		nums[low] = nums[high]
		fmt.Println("swap1: ", nums)

		for low < high && nums[low] <= pivot {
			low++
		}
		// 交换比基准小的数据到基准右边
		nums[high] = nums[low]
		fmt.Println("swap2: ", nums)
	}
	nums[low] = pivot
	fmt.Println("pivotIndex: ", low)
	return low
}

// Array array 结构体
// 堆排序 O(nlogn), 不稳定
type Array struct {
	Arr []int
}

// GetParentIndex 根据子节点下标获取父节点下标
func (p *Array) GetParentIndex(child int) int {
	return (child - 1) / 2
}

// GetLeftChildIndex 根据左子节点下标获取父节点上下标
func (p *Array) GetLeftChildIndex(parent int) int {
	return parent*2 + 1
}

// Swap 交换值
func (p *Array) Swap(i, j int) {
	p.Arr[i] += p.Arr[j]
	p.Arr[j] = p.Arr[i] - p.Arr[j]
	p.Arr[i] = p.Arr[i] - p.Arr[j]
}

// AdjustHeap 调整堆结构
func (p *Array) AdjustHeap(i, len int) {
	left, right, j := p.GetLeftChildIndex(i), 0, 0
	fmt.Println("out: ", left, right, len)
	for left <= len {
		right = left + 1
		j = left
		if j < len && p.Arr[left] < p.Arr[right] {
			j++
		}
		fmt.Println("x: ", i, j)
		if p.Arr[i] < p.Arr[j] {
			p.Swap(i, j)
			i = j
			left = p.GetLeftChildIndex(i)
		} else {
			break
		}
	}
}

// HeapSort 堆排序
func (p *Array) HeapSort() {
	last := len(p.Arr) - 1
	for i := p.GetParentIndex(last); i >= 0; i-- {
		fmt.Println("f: ", p.Arr)
		p.AdjustHeap(i, last)
		fmt.Println("--------------------")
	}

	for last > 0 {
		p.Swap(0, last)
		last--
		fmt.Println("l: ", p.Arr)
		p.AdjustHeap(0, last)
		fmt.Println("********************")
	}
}
