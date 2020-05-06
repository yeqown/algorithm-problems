# https://leetcode-cn.com/problems/minimum-cost-for-tickets/
# idx = 983

# 示例 1：
#
# 输入：days = [1,4,6,7,8,20], costs = [2,7,15]
# 输出：11
# 解释：
# 例如，这里有一种购买通行证的方法，可以让你完成你的旅行计划：
# 在第 1 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 1 天生效。
# 在第 3 天，你花了 costs[1] = $7 买了一张为期 7 天的通行证，它将在第 3, 4, ..., 9 天生效。
# 在第 20 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 20 天生效。
# 你总共花了 $11，并完成了你计划的每一天旅行。


class Solution:
    '''
    贪心算法 + DP

    '''

    def mincostTickets(self, days: [int], costs: [int]) -> int:
        daysset = set(days)
        memo = {}

        def dp(i: int) -> int:
            # dp(i)=min{cost(j)+dp(i+j)},j∈{1,7,30}
            if i > 365:
                return 0

            if i in memo:
                return memo[i]

            cost = float("INF")
            if i in daysset:
                # 需要出行
                cost = min(costs[0] + dp(i+1), costs[1] +
                           dp(i+7), costs[2] + dp(i+30))
            else:
                # 不需要出行
                cost = dp(i+1)
            memo[i] = cost
            return cost

        return dp(1)


if __name__ == "__main__":
    sol = Solution()

    days = [1, 4, 6, 7, 8, 20]
    costs = [2, 7, 15]
    res = sol.mincostTickets(days, costs)
    print("days = %s, costs = %s, got mincost =%d, want=11" % (days, costs, res))
