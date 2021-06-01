package main

import (
	"fmt"
	al "leek/algorithm"
	"leek/base"
	"leek/common"
	iv "leek/interview"
	lc "leek/leetcode"
	bt "leek/leetcode/bintree"
	"leek/leetcode/bst"
	list "leek/leetcode/list"
)

func Myslice() {
	var arr []interface{}
	arr = append(arr, 10)
	arr = append(arr, "sds")
	fmt.Println(arr)
}

func main001() {
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
	r := iv.AddNumString(num1, num2)
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

	ListNodearray := make([]*base.ListNode, len(multiArray))
	for i := 0; i < len(multiArray); i++ {
		ListNodearray[i] = base.CreateListNode(multiArray[i])
	}

	// test reverse
	for i := range ListNodearray {
		base.OuputListNode(ListNodearray[i])
		// newHeader := al.RecurseReverseListNode(ListNodearray[i])
		// newHeader := al.ReverseListNode(ListNodearray[i])
		// newHeader := al.RecurseReverseListNodeN(ListNodearray[i], 1)
		// newHeader := al.ReverseListNodeN(ListNodearray[i], 2)
		// newHeader := al.ReverseBetweenV2(ListNodearray[i], ListNodearray[i], ListNodearray[i].Next.Next)
		// newHeader := al.RecurseReverseKGroup(ListNodearray[i], 2)
		newHeader := al.ReverseListNodeV2(ListNodearray[i])
		base.OuputListNode(newHeader)
		fmt.Println("-----------")
	}
}

func main007() {
	PalindromeArray := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 2, 2, 1},
		{1, 2, 3, 3, 2, 1},
		{1, 2, 3, 4, 3, 2, 1},
		{1, 2, 3, 4, 3, 3, 1},
	}

	ListNodearray := make([]*base.ListNode, len(PalindromeArray))
	for i := 0; i < len(PalindromeArray); i++ {
		ListNodearray[i] = base.CreateListNode(PalindromeArray[i])
	}

	// test reverse
	for i := range ListNodearray {
		// result := lc.IsPalindromeV1(ListNodearray[i])
		// result := lc.IsPalindromeV2(ListNodearray[i])
		fmt.Printf("row: ")
		base.OuputListNode(ListNodearray[i])
		result := list.IsPalindromeV3(ListNodearray[i])
		fmt.Println(result)
		fmt.Println("------------")
	}
}

func main008() {
	stack := common.NewStack()
	array := []int{0, 10, 21, 13, 44}
	for _, v := range array {
		stack.Put(v)
	}
	fmt.Println(stack.Empty())
	fmt.Println("----------------")
	for !stack.Empty() {
		v, _ := stack.Get()
		// fmt.Printf("get: %v", v)
		fmt.Printf("%v - ", v)
		// fmt.Println("peek: ", stack.Size())
	}

	queue := common.NewQueue()
	for _, v := range array {
		queue.Push(v)
	}
	fmt.Println(queue.Empty())
	fmt.Println("----------------")
	for !queue.Empty() {
		v, _ := queue.Pop()
		fmt.Printf("%v - ", v)
		// fmt.Println("peek: ", queue.Size())
	}
}

func main009() {
	array := []interface{}{4, 2, 7, 1, 3, 6, 9}
	// 原始树
	root := al.CreateBinTree(array)
	info := al.LevelOrder(root)
	common.PrintItems(info)
	fmt.Println("----------------")
	// 翻转树
	// newroot := lc.InvertTree(root)
	newroot := bt.InvertTreeV2(root)
	info = al.LevelOrder(newroot)
	common.PrintItems(info)
	fmt.Println("----------------")
	array1 := []int{3, 2, 1, 6, 0, 5}
	fmt.Println(array1)
	r := bt.ConstructMaximumBinaryTree(array1)
	info = al.LevelOrder(r)
	common.PrintItems(info)
}

func main0010() {
	preorder := []int{1, 2, 5, 4, 6, 7, 3, 8, 9}
	inorder := []int{5, 2, 6, 4, 7, 1, 8, 3, 9}
	postorder := []int{5, 6, 7, 4, 2, 8, 9, 3, 1}
	root := bt.BuildTreeFromPreorderAndInorder(preorder, inorder)
	info := al.LevelOrder(root)
	common.PrintItems(info)
	fmt.Println("----------------")
	root = bt.BuildTreeFromInorderAndPostorder(inorder, postorder)
	info = al.LevelOrder(root)
	common.PrintItems(info)
	fmt.Println("----------------")
}

func main() {
	/*
		//test KthSmallest
		// array := []interface{}{3, 1, 4, nil, 2}
		// array := []interface{}{5, 3, 6, 2, 4, nil, nil, 1}
		// array := []interface{}{7, 5, 8, 3, 6, nil, nil, 2, 4, nil, nil, nil, nil, nil, nil, 1}
		array := []interface{}{10, 7, 13, 5, 8, 11, 14, 3, 6, nil, 9, nil, 12, nil, nil, 2, 4, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 1}

		root := al.CreateBinTree(array)
		info := al.LevelOrder(root)
		al.ReCountBinTreeNode(root)
		common.PrintItems(info)
		k := 6
		val := bst.KthSmallest(root, k)
		fmt.Println(val)
		val = bst.KthSmallestV2(root, k)
		fmt.Println(val)
	*/

	/*
		//test ConvertBST
		array := []interface{}{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}
		root := al.CreateBinTree(array)
		info := al.LevelOrder(root)
		common.PrintItems(info)
		fmt.Println("----------------")
		root = bst.ConvertBST(root)
		info = al.LevelOrder(root)
		common.PrintItems(info)
	*/
	/*
		array := []interface{}{4, 2, 7, 1, 3}
		root := al.CreateBinTree(array)
		info := al.LevelOrder(root)
		common.PrintItems(info)
		fmt.Println("----------------")
		node := bst.SearchBST(root, 2)
		info = al.LevelOrder(node)
		common.PrintItems(info)
	*/
	/*
		// array := []interface{}{2, 1, 3}
		// array := []interface{}{5, 1, 4, nil, nil, 3, 6}
		// array := []interface{}{1, 1}
		array := []interface{}{0}
		root := al.CreateBinTree(array)
		// info := al.LevelOrder(root)
		ok := bst.IsValidBST(root)
		fmt.Println(ok)
	*/
	/*
		// array := []interface{}{4, 2, 7, 1, 3}
		array := []interface{}{40, 20, 60, 10, 30, 50, 70}
		root := al.CreateBinTree(array)
		info := al.LevelOrder(root)
		common.PrintItems(info)
		fmt.Println("----------------")
		// root = bst.InsertIntoBSTV1(root, 5)
		bst.InsertIntoBSTV2(root, 25)
		info = al.LevelOrder(root)
		common.PrintItems(info)
	*/

	// /*
	// array := []interface{}{0}
	// array := []interface{}{2, 1}
	// array := []interface{}{2, 1, 3}
	// array := []interface{}{4, 2, 7, 1, 3}
	// array := []interface{}{5, 3, 6, 2, 4, nil, 7}

	array := []interface{}{40, 25, 60, 10, 30, 50, 70, nil, 20, nil, nil, 45}
	root := al.CreateBinTree(array)
	info := al.LevelOrder(root)
	common.PrintItems(info)
	key := 10
	// root = bst.InsertIntoBSTV1(root, 5)
	fmt.Println("----------------:", key)
	// root = bst.DeleteNodeFromBST(root, key)
	root = bst.DeleteNodeFromBSTV2(root, key)
	fmt.Println("----------------:")
	info = al.LevelOrder(root)
	common.PrintItems(info)
	// */
}
