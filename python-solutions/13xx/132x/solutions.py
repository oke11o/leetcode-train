import unittest
from typing import List


# 1329. Sort the Matrix Diagonally
# https://leetcode.com/problems/sort-the-matrix-diagonally/
# Medium
# Array, Sorting, Matrix
class Solution1329:
    def diagonalSort(self, mat: List[List[int]]) -> List[List[int]]:
        row = len(mat)
        if row < 2:
            return mat
        col = len(mat[0])
        if col < 2:
            return mat

        flat = [[] for _ in range(max(row, col) * 2)]
        for fl in range(max(row, col)):
            if fl == 0:
                for i in range(min(row, col)):
                    flat[fl].append(mat[i][i])
            else:
                for j in range(max(row, col)):
                    i = fl + j
                    if i < len(mat) and j < len(mat[0]):
                        flat[fl * 2 - 1].append(mat[i][j])

                    i, j = j, i
                    if i < len(mat) and j < len(mat[0]):
                        flat[fl * 2].append(mat[i][j])

        for fl in flat:
            fl.sort()

        for idx, fl in enumerate(flat):
            if len(fl) == 0:
                continue

            if idx == 0:
                for i, v in enumerate(fl):
                    mat[i][i] = v
            else:
                J = int((idx + 1) / 2)
                JJ = idx % 2

                for i, v in enumerate(fl):
                    j = i + J
                    if JJ == 1:
                        i, j = j, i

                    mat[i][j] = v

        return mat


#####################################
############### TESTS ###############
#####################################

cases = [
    # {
    #     "input": [[3, 3, 1, 1], [2, 2, 1, 2], [1, 1, 1, 2]],
    #     "want": [[1, 1, 1, 1], [1, 2, 2, 2], [1, 2, 3, 3]],
    # },
    # {
    #     "input": [[11, 25, 66, 1, 69, 7], [23, 55, 17, 45, 15, 52], [75, 31, 36, 44, 58, 8], [22, 27, 33, 25, 68, 4],
    #               [84, 28, 14, 11, 5, 50]],
    #     "want": [[5, 17, 4, 1, 52, 7], [11, 11, 25, 45, 8, 69], [14, 23, 25, 44, 58, 15], [22, 27, 31, 36, 50, 66],
    #              [84, 28, 75, 33, 55, 68]],
    # },
    {
        "input": [[75, 21, 13, 24, 8], [24, 100, 40, 49, 62]],
        "want": [[75, 21, 13, 24, 8], [24, 100, 40, 49, 62]],
    },
]


class Solution1329TestCase(unittest.TestCase):
    def setUp(self) -> None:
        self.sol = Solution1329()

    def test(self):
        for i, t in enumerate(cases):
            want = t['want']
            input = t['input']

            res = self.sol.diagonalSort(input)
            self.assertEqual(res, want, f"Idx: {i}; Input: {input}")
