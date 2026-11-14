package leetcode

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	result := 0

	for _, price := range prices[1:] {
		result = max(result, price-minPrice)
		minPrice = min(minPrice, price)
	}

	return result
}
