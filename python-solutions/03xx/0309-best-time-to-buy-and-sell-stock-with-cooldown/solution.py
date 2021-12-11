import unittest
from typing import List


# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
# https://www.youtube.com/watch?v=aC416p0pjDM&list=PLyHj6yBbnkUgC6T9RpPxEBUv3My9lPd7m&index=56
# 309. Best Time to Buy and Sell Stock with Cooldown
# Medium
# Array, Dynamic Programming
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy = float('-inf')
        sell = 0
        cooldown = 0

        for price in prices:
            next_buy = max(buy, cooldown - price)  # Мы купили, только если подождали
            next_sell = buy + price  # Мы можем продать, только если мы на предыдущем шаге купили
            cooldown = max(sell, cooldown) # ...

            buy = next_buy
            sell = next_sell

        return max(sell, cooldown)


cases = [
    {
        "prices": [1, 2, 3, 0, 2],
        "want": 3,
    },
    {
        "prices": [1],
        "want": 0,
    },
]


class SolutionTestCase(unittest.TestCase):
    def setUp(self) -> None:
        self.sol = Solution()

    def test(self):
        for i, t in enumerate(cases):
            want = t['want']
            prices = t['prices']

            res = self.sol.maxProfit(prices)
            self.assertEqual(res, want, f"Idx: {i}; Prices: {prices}")
