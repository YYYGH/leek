package algorithm

import (
	// "fmt"
	"fmt"
	"leek/common"
	"math"
	"sort"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

/*
322. 零钱兑换, https://leetcode-cn.com/problems/coin-change/
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。
示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0
示例 4：

输入：coins = [1], amount = 1
输出：1
示例 5：

输入：coins = [1], amount = 2
输出：2
*/
func CoinChangeV1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make(map[int]int)
	count := subCoinChange(dp, coins, amount)
	if count == math.MaxInt32 {
		return -1
	}
	return count
}

func subCoinChange(dp map[int]int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if amount < 0 {
		return -1
	}
	if v, ok := dp[amount]; ok {
		return v
	}

	mincount := math.MaxInt32
	for _, v := range coins {
		count := subCoinChange(dp, coins, amount-v)
		if count < 0 {
			continue
		}
		mincount = common.Min(mincount, count+1)
	}
	if mincount == math.MaxInt32 {
		dp[amount] = -1
	} else {
		dp[amount] = mincount
	}
	return dp[amount]
}

// 1. 对零钱面值进行排序
// 2. 用大到小的面值开始组合  amount/coin_value   amount%coinvalue

func CoinChangeV2(coins []int, amount int) int {
	sort.Ints(coins)
	coindp := make(map[string]int)
	fmt.Println(coindp, &coindp)
	fmt.Printf("ss: %p, %p, %p\n", &coindp, &coins, &coins[0])
	count := subCoinChangeV2(coindp, coins, amount, len(coins)-1)
	if count == math.MaxInt32 {
		return -1
	}
	return count
}

func subCoinChangeV2(coindp map[string]int, coins []int, amount, index int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 || index < 0 {
		return -1
	}
	// fmt.Printf("start amount: %d, index: %d\n", amount, index)
	key := fmt.Sprintf("%v-%v", amount, index)
	if v, ok := coindp[key]; ok {
		fmt.Printf("%v->%v\n", key, v)
		return v
	}
	n := amount / coins[index]
	mincount := math.MaxInt32
	for i := n; i >= 0; i-- {
		count := subCoinChangeV2(coindp, coins, amount-i*coins[index], index-1)
		if count == -1 {
			continue
		}
		mincount = common.Min(mincount, count+i)
		if count+1 < mincount {
			mincount = count + i
		} else {
			break
		}
	}

	if mincount == math.MaxInt32 {
		mincount = -1
	}
	// coindp[key] = mincount
	return mincount
}

/*
贪心
11. 想要总硬币数最少，肯定是优先用大面值硬币，所以对 coins 按从大到小排序
12. 先丢大硬币，再丢会超过总额时，就可以递归下一层丢的是稍小面值的硬币

乘法对加法的加速
21. 优先丢大硬币进去尝试，也没必要一个一个丢，可以用乘法算一下最多能丢几个

k = amount / coins[c_index] 计算最大能投几个
amount - k * coins[c_index] 减去扔了 k 个硬币
count + k 加 k 个硬币

如果因为丢多了导致最后无法凑出总额，再回溯减少大硬币数量
最先找到的并不是最优解
31. 注意不是现实中发行的硬币，面值组合规划合理，会有奇葩情况
32. 考虑到有 [1,7,10] 这种用例，按照贪心思路 10 + 1 + 1 + 1 + 1 会比 7 + 7 更早找到
33. 所以还是需要把所有情况都递归完

ans 疯狂剪枝
41. 贪心虽然得不到最优解，但也不是没用的
42. 我们快速算出一个贪心的 ans 之后，虽然还会有奇葩情况，但是绝大部分普通情况就可以疯狂剪枝了

void coinChange(vector<int>& coins, int amount, int c_index, int count, int& ans) {
    if (amount == 0) {
        ans = min(ans, count);
        return;
    }
    if (c_index == coins.size()) return;

    for (int k = amount / coins[c_index]; k >= 0 && k + count < ans; k--) {
        coinChange(coins, amount - k * coins[c_index], c_index + 1, count + k, ans);
    }
}

int coinChange(vector<int>& coins, int amount) {
    if (amount == 0) return 0;
    sort(coins.rbegin(), coins.rend());
    int ans = INT_MAX;
    coinChange(coins, amount, 0, 0, ans);
    return ans == INT_MAX ? -1 : ans;
}

*/

func CoinChangeV4(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	// 从0元开始凑
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i >= coin && dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
