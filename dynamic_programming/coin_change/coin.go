package coin_change

// We need to use dynamic programming to solve this problem, unless we want to calculate all possible branches.
// We basically view the problem as a set of sub problems, and we start to solve them from the easiest of all, for amount 0.
// The final answer will come from what we previously computed for smaller amounts.
//
// We know that the minimum amount of coins to reach a 0 amount is 0, we can so initialise the first result with 0.
// We also know that we can have a maximum of amount+1 results (0..amount) and we can initialise those too with an unrealistic value like amount+1 or max int available.
// This will also be used to understand if we actually found coins to satisfy the criteria, otherwise we can return -1.
//
// Now, when we approach the next sub problem for amount=1, need to evaluate all coins.
// When we take a coin into consideration, we want to understand if that will actually help us by subtracting it from the amount.
// If the latter is < 0, it means that the coin is too big and we don't need it.
// When we actually find a useful coin, we want to check what's the minimum amount of coins required bewteen what we already have for this amount
// and by adding 1 (since we are adding a coin) and the minimum number of coins that we computed to get amount-coin.
//
// T: O(c*a)
// S: O(a+1)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	for amt := 1; amt < amount+1; amt++ {
		for _, c := range coins {
			if amt-c >= 0 {
				dp[amt] = min(dp[amt], 1+dp[amt-c])
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}
