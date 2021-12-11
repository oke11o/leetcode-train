import unittest
from typing import List

# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
# https://www.youtube.com/watch?v=aC416p0pjDM&list=PLyHj6yBbnkUgC6T9RpPxEBUv3My9lPd7m&index=56
# 714. Best Time to Buy and Sell Stock with Transaction Fee
# Medium
# Array, Dynamic Programming, Greedy
class Solution:
    def maxProfit(self, prices: List[int], fee: int) -> int:
        buy = float('-inf')
        sell = 0

        for price in prices:
            next_buy = max(buy, sell - (price + fee))
            next_sell = max(sell, buy + price)

            buy = next_buy
            sell = next_sell

        return sell


cases = [
    {
        "prices": [1, 3, 2, 8, 4, 9],
        "fee": 2,
        "want": 8,
    },
    {
        "prices": [1, 3, 7, 5, 10, 3],
        "fee": 3,
        "want": 6,
    },
]


class SolutionTestCase(unittest.TestCase):
    def setUp(self) -> None:
        self.sol = Solution()

    def test(self):
        for i, t in enumerate(cases):
            want = t['want']
            prices = t['prices']
            fee = t['fee']

            res = self.sol.maxProfit(prices, fee)
            self.assertEqual(res, want, f"Idx: {i}; Prices: {prices}")
