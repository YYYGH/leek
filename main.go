package main

import (
	"fmt"
	al "leek/algorithm"
	"leek/interview"
	iv "leek/interview"
	lc "leek/leetcode"
)

func Myslice() {
	var arr []interface{}
	arr = append(arr, 10)
	arr = append(arr, "sds")
	fmt.Println(arr)
}

func main001() {
	fmt.Printf("")
	list := []int{2, 4, 7, 15, 14}
	// list := []int{1, 3, 5, 6}
	num := 2
	fmt.Println(iv.CheckShuzi(list))
	fmt.Println(iv.SearchInsert(list, num))

	intervals := [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}
	// intervals := [][]int{{1, 4}, {0, 2}, {3, 5}}
	// intervals := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	// intervals := [][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}
	// [[1,4],[0,2],[3,5]]
	r := lc.MergeArray(intervals)
	fmt.Println(r)

	s := []int{10, 2, 34, 4, 6, 1, 9, 2} // [1 2 2 4 6 9 10 34]
	// al.BubbleSort(s)
	// al.SelectSort(s)
	// al.InsertSort(s)
	// al.MergeSort(s)
	// al.QuickSort(s)
	a := &al.Array{}
	a.Arr = s
	a.HeapSort()
	// fmt.Println(s)

	str := "ADOBECODEBANC"
	t := "ABC"

	// str := "acbbaca"
	// t := "aba"
	result := lc.MinWindow(str, t)
	fmt.Println(result)
	lc.MinPathSum(nil)
	// 9999 表示不通的路
	grids := [][]int{
		{1, 2, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 3, 1, 1},
		{1, 1, 9999, 1, 1, 9999, 1},
		{1, 1, 1, 9999, 1, 1, 1},
		{1, 1, 1, 1, 2, 1, 1},
		{1, 1, 1, 2, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}

	// path := iv.MinPathWay(grids, []int{4, 1}, []int{1, 5}) //customer_time:  10  [[4 1] [3 1] [2 1] [1 1] [1 2] [1 3] [1 4] [1 5]]
	// path := iv.MinPathWay(grids, []int{4, 1}, []int{1, 4}) //customer_time:  9  [[4 1] [3 1] [2 1] [1 1] [1 2] [1 3] [1 4]]
	// path := iv.MinPathWay(grids, []int{4, 1}, []int{2, 4}) //customer_time:  7   [[4 1] [4 2] [4 3] [4 4] [3 4] [2 4]]
	path := iv.MinPathWay(grids, []int{4, 1}, []int{1, 6}) //customer_time:  10 [[4 1] [4 2] [4 3] [4 4] [3 4] [3 5] [3 6] [2 6] [1 6]]
	fmt.Println("path:", path)
}

func main003() {
	// array := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// 1 2 4 7 5 3 6 8 9
	// array := [][]int{
	// 	{1, 2, 3, 4, 5},
	// 	{6, 7, 8, 9, 10},
	// 	{11, 12, 13, 14, 15},
	// } // 1 2 6 11 7 3 4 8 12 13  9  5 10 14 15
	array := [][]int{
		{1, 2, 3},
		{6, 7, 8},
		{11, 12, 13},
		{16, 17, 18},
	} // 1 2 6 11 7 3 8 12 16 17 13 18

	r := lc.FindDiagonalOrder(array)
	fmt.Println(r)
}

func main005() {
	num2 := "8321"
	num1 := "13229"
	r := interview.AddNumString(num1, num2)
	fmt.Println(r)
}

func main006() {
	multiArray := [][]int{
		// {1},
		// {1, 2},
		// {1, 2, 3},
		{1, 2, 3, 4, 5},
		// {1, 2, 3, 4, 5, 6},
		// {1, 2, 3, 4, 5, 6, 7},
	}

	ListNodearray := make([]*al.ListNode, len(multiArray))
	for i := 0; i < len(multiArray); i++ {
		ListNodearray[i] = al.CreateListNode(multiArray[i])
	}

	// test reverse
	for i := range ListNodearray {
		al.OuputListNode(ListNodearray[i])
		// newHeader := al.RecurseReverseListNode(ListNodearray[i])
		// newHeader := al.ReverseListNode(ListNodearray[i])
		// newHeader := al.RecurseReverseListNodeN(ListNodearray[i], 1)
		// newHeader := al.ReverseListNodeN(ListNodearray[i], 2)
		// newHeader := al.ReverseBetweenV2(ListNodearray[i], ListNodearray[i], ListNodearray[i].Next.Next)
		// newHeader := al.RecurseReverseKGroup(ListNodearray[i], 2)
		newHeader := al.ReverseListNodeV2(ListNodearray[i])
		al.OuputListNode(newHeader)
		fmt.Println("-----------")
	}
}

func main() {
	PalindromeArray := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 2, 2, 1},
		{1, 2, 3, 3, 2, 1},
		{1, 2, 3, 4, 3, 2, 1},
		{1, 2, 3, 4, 3, 3, 1},
	}

	ListNodearray := make([]*al.ListNode, len(PalindromeArray))
	for i := 0; i < len(PalindromeArray); i++ {
		ListNodearray[i] = al.CreateListNode(PalindromeArray[i])
	}

	// test reverse
	for i := range ListNodearray {
		// result := lc.IsPalindromeV1(ListNodearray[i])
		// result := lc.IsPalindromeV2(ListNodearray[i])
		fmt.Printf("row: ")
		al.OuputListNode(ListNodearray[i])
		result := lc.IsPalindromeV3(ListNodearray[i])
		fmt.Println(result)
		fmt.Println("------------")
	}
}
