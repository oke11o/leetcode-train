import unittest
from typing import List


# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
# https://www.youtube.com/watch?v=aC416p0pjDM&list=PLyHj6yBbnkUgC6T9RpPxEBUv3My9lPd7m&index=56
# 123. Best Time to Buy and Sell Stock III
# Hard
# Array, Dynamic Programming
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy_1 = float('-inf')  # buy 1 sell 0
        buy_2 = float('-inf')  # buy 2 sell 1
        sell_1 = 0  # buy 1 sell 1
        sell_2 = 0  # buy 2 sell 2

        for price in prices:
            next_buy_1 = max(buy_1, 0 - price)
            next_buy_2 = max(buy_2, sell_1 - price)
            next_sell_1 = max(sell_1, buy_1 + price)
            next_sell_2 = max(sell_2, buy_2 + price)

            buy_1 = next_buy_1
            buy_2 = next_buy_2
            sell_1 = next_sell_1
            sell_2 = next_sell_2

        return max(sell_1, sell_2)


#################
##### TESTS #####
#################

cases = [
    {
        "prices": [3, 3, 5, 0, 0, 3, 1, 4],
        "want": 6,
    },
    {
        "prices": [1, 2, 3, 4, 5],
        "want": 4,
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
