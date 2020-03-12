# 输入: str1 = "abcde", str2 = "ace"
# 输出: 3
# 解释: 最长公共子序列是 "ace"，它的长度是 3


def mat_constructor(row: int, col: int, default_value: int):
    mat = []
    for _ in range(row):
        r = [default_value for _ in range(col)]
        mat.append(r)
    return mat


def print_mat(mat):
    for row in mat:
        print(row)


# 最长公共子序列
# 时间复杂度：O(m * n)
# 空间复杂度：O(m+1 * n+1)
# 如何根据获得序列：在 dp[i][j] = dp[i-1][j-1] + 1 记录该字符
def lcs(str1: str, str2: str) -> str:
    # dp[i,j] = dp[i-1, j-1] + 1,               str1[i] == str[j]
    #         = max(dp[i-1, j], dp[i, j-1]),    str1[i] != str[j]
    #         = 0,                              i = 0, j = 0

    len1 = len(str1)
    len2 = len(str2)

    # 增加 i=0, j=0 行列，减少溢出判断
    # len1 as row 行
    # len2 as col 列
    dp = mat_constructor(len1+1, len2+1, 0)
    # dp:
    #        x=0 x=1 x=2 x=3 [str2 = ace]
    #             a   c   e
    # y=0     0   0   0   0
    # y=1  a  0   1   1   1
    # y=2  b  0   1   1   1
    # y=3  c  0   1   2   2
    # y=4  d  0   1   2   2
    # y=5  e  0   1   2   3
    # [str1 = abcde]

    for y in range(1, len1+1):
        for x in range(1, len2+1):
            # print(str1[y-1], str2[x-1])
            if str1[y-1] == str2[x-1]:
                # true: equal
                dp[y][x] = dp[y-1][x-1] + 1
            else:
                dp[y][x] = max(dp[y-1][x], dp[y][x-1])
    print_mat(dp)
    return dp[y][x]


# 最长公共子串
# 时间复杂度：O(m * n)
# 空间复杂度：O(m+1 * n+1)
# 如何获得序列，记录最大子串的最终位置（str1, str2）和长度，倒推获得字符串切片
def lc_substring(str1: str, str2: str) -> str:
    # dp[i][j] = dp[i-1][j-1] + 1, str1[i] == str2[j]
    #          = 0               , str1[i] != str2[j]

    len1 = len(str1)
    len2 = len(str2)

    # 增加 i=0, j=0 行列，减少溢出判断
    # len1 as row 行
    # len2 as col 列
    dp = mat_constructor(len1+1, len2+1, 0)
    max_mark = 0
    max_y = -1
    max_x = -1

    for y in range(1, len1+1):
        for x in range(1, len2+1):
            # print(str1[y-1], str2[x-1])
            if str1[y-1] == str2[x-1]:
                # true: equal
                dp[y][x] = dp[y-1][x-1] + 1
            else:
                dp[y][x] = 0

            if dp[y][x] > max_mark:
                max_mark = dp[y][x]
                max_y = y
                max_x = x
    print_mat(dp)
    print(max_y, max_x)

    return max_mark


if __name__ == "__main__":
    res = lcs("abcde", "ace")  # 3
    print(res)

    res2 = lc_substring("abcde", "bcdbcde")
    print(res2)
