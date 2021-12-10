import unittest

from typing import List

# 0121 bestTimeToBuyAndSellStock
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
# https://www.youtube.com/watch?v=aC416p0pjDM&list=PLyHj6yBbnkUgC6T9RpPxEBUv3My9lPd7m&index=56
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if not prices:
            return 0

        max_profit = 0
        current_min = prices[0]
        for i in range(1, len(prices)):
            price = prices[i]
            max_profit = max(max_profit, price - current_min)
            current_min = min(current_min, price)

        return max_profit


cases = [
    {
        "name": "test",
        "input": [7, 1, 5, 3, 6, 4],
        "output": 5,
    },
    {
        "name": "test2",
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
