# 给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。
# 一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。
# 你可以假设网格的四个边均被水包围。
#
# 1 1 1 1 0
# 1 1 0 1 0
# 1 1 0 0 0
# 0 0 0 0 0
#
# output: 1
#

# from ..dp.lcs import mat_constructor, print_mat
# import sys
# import os
# sys.path.append(os.path.abspath("../dp"))


class Solution:

    def numIslands(self, grid: [[str]]) -> int:
        memo = {}
        count = 0

        def bfs(i: int, j: int) -> int:

            # 已经访问过的，不计入新岛
            if (i, j) in memo:
                return 0
            # 遍历的点是水，也不计入岛
            if grid[i][j] == "0":
                return 0

            # 标记当前点已经被访问
            # print(i, j, memo)
            memo[(i, j)] = True
            if (i - 1) >= 0 and grid[i-1][j] == "1":
                bfs(i-1, j)
            if (j - 1) >= 0 and grid[i][j-1] == "1":
                bfs(i, j-1)
            if (i+1) <= (len(grid) - 1) and grid[i+1][j] == "1":
                bfs(i+1, j)
            if (j+1) <= (len(grid[i]) - 1) and grid[i][j+1] == "1":
                bfs(i, j+1)

            return 1

        for i in range(len(grid)):
            for j in range(len(grid[i])):
                # 如何计数
                count += bfs(i, j)
        return count


if __name__ == "__main__":
    sol = Solution()
    # mat = [
    #     ["1", "1", "1", "1", "0"],
    #     ["1", "1", "0", "1", "0"],
    #     ["1", "1", "0", "0", "0"],
    #     ["0", "0", "0", "0", "0"],
    # ]
    # res = sol.numIslands(mat)             # want 1
    # print("num of islands is: %d" % res)  # got 1
    mat = [
        ["1", "1", "1"],
        ["0", "1", "0"],
        ["0", "1", "0"],
    ]
    res = sol.numIslands(mat)             # want 1
    print("num of islands is: %d" % res)  # got 2

    mat = [
        ["1", "1", "0", "0", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "1", "0", "0"],
        ["0", "0", "0", "1", "1"],
    ]

    res = sol.numIslands(mat)             # want 3
    print("num of islands is: %d" % res)  # got 3
