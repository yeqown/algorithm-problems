from lcs import print_mat, mat_constructor


def matrix_path(N: int) -> int:
    '''
    n*n矩阵从左上角到右下角有多少种走法（只限往下和往右走）
    '''
    # dp =
    #   1 2 3
    #  -------
    # 1|1 1 1
    # 2|1 2 3
    # 3|1 3 6
    if N == 0:
        return 0
    # init
    dp_mat = mat_constructor(N+1, N+1, 0)

    # dp_mat[i,j] = dp[i-1][j] + dp[i][j-1]
    # 注意i, j越界
    for i in range(1, N+1):
        for j in range(1, N+1):
            if i == 1 and j == 1:
                dp_mat[1][1] = 1
                continue
            # 只能向右或者向下走
            dp_mat[i][j] = dp_mat[i-1][j] + dp_mat[i][j-1]
    print_mat(dp_mat)

    return dp_mat[N][N]


if __name__ == "__main__":
    N = 4
    res = matrix_path(N)
    print("N = %d, res = %d" % (N, res))
