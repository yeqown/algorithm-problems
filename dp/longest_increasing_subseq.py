# 给定一个无序的整数数组，找到其中最长上升子序列的长度
# dp[i] = max(dp[j, 0...i-1]) + 1, arr[j] < arr[i]
# dp[0] = 1


def lis_dp(arr: slice) -> int:
    max_mark = 0
    dp_table = [1]

    def find_pre_less_max(i: int) -> int:
        # 找到前序中，比当前值小的最大长度
        j = i-1
        _max = 0
        while j >= 0:
            if arr[j] < arr[i]:
                if _max < dp_table[j]:
                    _max = dp_table[j]
            j -= 1
        return _max

    for i in range(1, len(arr)):
        last_max = find_pre_less_max(i)
        dp_table.append(last_max + 1)

        if dp_table[i] > max_mark:
            max_mark = dp_table[i]
        print(dp_table)

    return max_mark


if __name__ == "__main__":
    # res = lis_dp([10, 9, 2, 5, 3, 7, 101, 18])
    # dp_table  = [1,1,1,2,2,3,4,4]
    res = lis_dp([10, 9, 2, 5, 3, 2, 101, 4])
    # dp_table  = [1, 1, 1, 2, 2, 1, 3, 3]
    print(res)
