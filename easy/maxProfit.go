package main

func main() {

}

func maxProfit(prices []int) int {
	minimum := prices[0]
	profit := 0

	for _, v := range prices {
		minimum = min(v, minimum)
		profit = max(profit, v-minimum)
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// const maxProfit = (prices) => {
// 	let left = 0; // Buy
// 	let right = 1; // sell
// 	let max_profit = 0;
// 	while (right < prices.length) {
// 	  if (prices[left] < prices[right]) {
// 		let profit = prices[right] - prices[left]; // our current profit

// 		max_profit = Math.max(max_profit, profit);
// 	  } else {
// 		left = right;
// 	  }
// 	  right++;
// 	}
// 	return max_profit;
//   };
