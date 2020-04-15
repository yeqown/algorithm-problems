import math
import collections


def mat_constructor(row: int, col: int, default_value: int):
    mat = []
    for _ in range(row):
        r = [default_value for _ in range(col)]
        mat.append(r)
    return mat


def print_mat(mat):
    for row in mat:
        print(row)


class Solution:
    def updateMatrix(self, matrix: [[int]]) -> [[int]]:
        m = len(matrix)
        n = len(matrix[0])

        dist = [[0] * n for _ in range(m)]
        zeroes_pos = [(i, j) for i in range(m)
                      for j in range(n) if matrix[i][j] == 0]
        # 将所有的 0（源点） 添加进初始队列中
        q = collections.deque(zeroes_pos)
        seen = set(zeroes_pos)
        print(seen)

        # 从0开始广度优先遍历
        while q:
            i, j = q.popleft()
            # 每个点只会更新其附近的点，且只会被更新一次
            # 因为是从0出发的，因此保证路径为最短
            for ni, nj in [(i-1, j), (i+1, j), (i, j-1), (i, j+1)]:
                if 0 <= ni < m and 0 <= nj < n and (ni, nj) not in seen:
                    dist[ni][nj] = dist[i][j]+1
                    seen.add((ni, nj))
                    q.append((ni, nj))

        return dist


if __name__ == "__main__":
    sol = Solution()
    mat = [
        [0, 0, 0],
        [0, 1, 0],
        [0, 0, 0],
    ]
    print_mat(mat)
    res = sol.updateMatrix(mat)
    print_mat(res)
