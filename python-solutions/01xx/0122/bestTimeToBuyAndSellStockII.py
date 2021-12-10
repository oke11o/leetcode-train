import unittest

from typing import List


# 0122 bestTimeToBuyAndSellStockII
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
# https://www.youtube.com/watch?v=aC416p0pjDM&list=PLyHj6yBbnkUgC6T9RpPxEBUv3My9lPd7m&index=56
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy = float('-inf')
        sell = 0

        for price in prices:
            next_buy = max(buy, sell - price)
            next_sell = max(sell, buy + price)

            buy = next_buy
            sell = next_sell

        return sell


cases = [
    {
        "name": "test",
        "input": [7, 1, 5, 3, 6, 4],
        "output": 7,
    },
    {
        "name": "test2",
        "input": [1, 2, 3, 4, 5],
        "output": 4,
    },
    {
        "name": "test3",
        "input": [7, 6, 4, 3, 1],
        "output": 0,
    },
]


class SolutionTestCase(unittest.TestCase):
    def setUp(self) -> None:
        self.sol = Solution()

    def test(self):
        for t in cases:
            name = t["name"]
            input = t["input"]
            output = t['output']

            res = self.sol.maxProfit(input)
            self.assertEqual(res, output, f"Name: {name}, Input: {input}")
