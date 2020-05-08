# idx = 221
# 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
#
# 示例:
#
# 输入:
# 1 0 1 0 0
# 1 0 1 1 1
# 1 1 1 1 1
# 1 0 0 1 0
#
# 输出: 4


class Solution:
    def maximalSquare(self, matrix: [[str]]) -> int:
        if matrix is None or len(matrix) == 0 or len(matrix[0]) == 0:
            return 0
        n = len(matrix)
        m = len(matrix[0])

        dp = [[1 if (i == 0 or j == 0) and matrix[i][j] ==
               "1" else 0 for j in range(m)] for i in range(n)]
        maxWidth = 0
        for i in range(0, n):
            for j in range(0, m):
                if i == 0 or j == 0:
                    # dp[i][j] = matrix[i][j]
                    maxWidth = max(maxWidth, dp[i][j])
                    continue

                if matrix[i][j] != "1":
                    dp[i][j] = 0
                    continue
                # matix[i][j] == 1
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
                maxWidth = max(maxWidth, dp[i][j])

        return maxWidth * maxWidth


if __name__ == "__main__":
    sol = Solution()
    mat = [["1", "0", "1", "0", "0"],
           ["1", "0", "1", "1", "1"],
           ["1", "1", "1", "1", "1"],
           ["1", "0", "0", "1", "0"]]
    res = sol.maximalSquare(mat)
    print("sol.maximalSquare(mat)\n matrix=%s \n want=%d, got=%d" %
          (mat, 4, res))
