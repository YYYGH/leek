package algorithm

/*
509. 斐波那契数
*/
func FibonacciRecurse(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return FibonacciRecurse(n-1) + FibonacciRecurse(n-2)
}

var dp []int

func FibonacciRecurseV2(n int) int {
	if n < 1 {
		return 0
	}
	dp = make([]int, n)
	return fibRecurseV2(n, dp)
}

func fibRecurseV2(n int, dp []int) int {
	if dp[n-1] != 0 {
		return dp[n-1]
	}
	if n == 1 || n == 2 {
		dp[n-1] = 1
	} else {
		dp[n-1] = fibRecurseV2(n-1, dp) + fibRecurseV2(n-2, dp)
	}
	return dp[n-1]
}

func FibonacciForLoop(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	array := make([]int, n)
	array[0] = 1
	array[1] = 1
	for i := 3; i <= n; i++ {
		array[i-1] = array[i-2] + array[i-3]
	}
	return array[n-1]
}

func FibonacciForLoopV2(n int) int {
	if n < 1 {
		return 0
	}
	prev := 1
	cur := 1
	for i := 3; i <= n; i++ {
		tmp := prev + cur
		prev = cur
		cur = tmp
	}
	return cur
}
