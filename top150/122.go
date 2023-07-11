package top150

func maxProfit2(prices []int) int {
	mProfit := 0
	stockPrice := -1
	for i := 0; i < len(prices); i++ {
		if stockPrice < 0 {
			if i+1 < len(prices) && prices[i+1] > prices[i] {
				stockPrice = prices[i]
			}
		} else {
			if (i+1 < len(prices) && prices[i+1] < prices[i]) || i == len(prices)-1 {
				mProfit += prices[i] - stockPrice
				stockPrice = -1
			}
		}
	}
	return mProfit
}
