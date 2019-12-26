package algorithm

/**
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.

Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var (
		buy  = -prices[0]
		sell = 0
	)

	for i := 1; i < len(prices); i++ {
		buy = max(buy, -prices[i])
		sell = max(sell, prices[i]+buy)
	}

	return sell
}

/**
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit.
You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.

Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.

*/

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var (
		buy  = -prices[0]
		sell = 0
	)

	for i := 1; i < len(prices); i++ {
		buy = max(buy, sell-prices[i])
		sell = max(sell, buy+prices[i])
	}
	return sell
}

/**
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.

Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.

Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

*/

func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var (
		fstBuy  = -prices[0]
		fstSell int
		secBuy  = -prices[0]
		secSell int
	)

	for i := 1; i < len(prices); i++ {
		fstBuy = max(fstBuy, -prices[i])
		fstSell = max(fstSell, fstBuy+prices[i])
		secBuy = max(secBuy, fstSell-prices[i])
		secSell = max(secSell, secBuy+prices[i])
	}

	return secSell
}
