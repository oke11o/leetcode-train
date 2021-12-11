import unittest
from typing import List


# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
# https://www.youtube.com/watch?v=aC416p0pjDM&list=PLyHj6yBbnkUgC6T9RpPxEBUv3My9lPd7m&index=56
# 188. Best Time to Buy and Sell Stock IV
# Hard
# Array, Dynamic Programming
class Solution:
    def maxProfit(self, k: int, prices: List[int]) -> int:
        buy = [float('-inf') for _ in range(k)]
        sell = [0 for _ in range(k + 1)]  # !!!!!

        for price in prices:
            next_buy = [float('-inf') for _ in range(k)]
            next_sell = [0 for _ in range(k + 1)]

            for i in range(k):
                next_buy[i] = max(buy[i], sell[i] - price)
                # K+1 нужно для этого, так как мы считаем след покупку на основании предыдущей продажи
                next_sell[i + 1] = max(sell[i + 1], buy[i] + price)

            buy = next_buy
            sell = next_sell

        return max(sell)


#####################################
############### TESTS ###############
#####################################

cases = [
    {
        "prices": [2, 4, 1],
        "k": 2,
        "want": 2,
    },
    {
        "prices": [3, 2, 6, 5, 0, 3],
        "k": 2,
        "want": 7,
    },
]


class SolutionTestCase(unittest.TestCase):
    def setUp(self) -> None:
        self.sol = Solution()

    def test(self):
        for i, t in enumerate(cases):
            want = t['want']
            prices = t['prices']
            k = t['k']

            res = self.sol.maxProfit(k, prices)
            self.assertEqual(res, want, f"Idx: {i}; Prices: {prices}")
