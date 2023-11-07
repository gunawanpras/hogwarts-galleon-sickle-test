package calculate

// CountWaysToMakeChange will find the different ways to get G3.5
func CountWaysToMakeChange(coins []int, target int) int {
	ways := make([]int, target+1)
	ways[0] = 1

	for _, coin := range coins {
		for amount := coin; amount <= target; amount++ {
			ways[amount] += ways[amount-coin]
		}
	}

	return ways[target]
}
