# 若干层楼，若干个鸡蛋，让你算出最少的尝试次数，找到鸡蛋恰好摔不碎的那层楼。
# 你面前有一栋从 1 到 N 共 N 层的楼，然后给你 K 个鸡蛋（K 至少为 1）。
# 现在确定这栋楼存在楼层 0 <= F <= N，在这层楼将鸡蛋扔下去，鸡蛋恰好没摔碎（高于 F 的楼层都会碎，低于 F 的楼层都不会碎）。
# 现在问你，最坏情况下，你至少要扔几次鸡蛋，才能确定这个楼层 F 呢？

# 状态：当前剩余鸡蛋数K，楼层N
# 选择：扔鸡蛋的楼层I


# DP(K, N) = MIN(DP)

def throw_egg_dp(K: int, N: int) -> int:
    # 子问题数：O(KN)
    # 时间复杂度：O(KN) * O(N)

    memo = dict()

    def dp(K, N) -> int:
        if N == 0:
            # 0 楼层，不需要扔
            return 0

        if K == 1:
            # 只有1枚鸡蛋，只能线性扫描
            return N

        # 备忘录，剪枝
        if (K, N) in memo:
            return memo[(K, N)]

        f = float('INF')  # maxest as initial value
        for i in range(1, N+1):
            f = min(
                f,
                max(
                    dp(K-1, i-1),  # 碎
                    dp(K, N-i),  # 没碎
                )+1,
            )

        memo[(K, N)] = f
        return f

    return dp(K, N)


if __name__ == "__main__":
    res = throw_egg_dp(2, 100)
    print(res)
