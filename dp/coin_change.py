class Solution:
    def waysToChange1(self, n: int) -> int:
        # dp [0...n]
        dp = [1] + [0] * n
        coins = [1, 5, 10, 25]

        for c in coins:
            for i in range(1, n+1):
                if i < c:
                    continue
                dp[i] += dp[i-c]

        return dp[n] % 1000000007

    def waysToChange(self, n: int) -> int:
        # n = a + 5b + 10c = 25d => 5m + a = a + 5(b + 2c + 5d)
        # b + 2c + 5d = m 的非负整数解个数，m = [0, n/5]
        # b + 2c = m - 5d, d = [0, m/5] => (m - 5d) / 2 + 1
        #
        #

        # 10 , m = [0, 1, 2], d = [0]
        # 1 + 1 + 2

        # 20, m = [0,1,2,3,4], d = [0]
        # 9
        # dp = {0: 1, 1: 1, 2:2, 3:2, 4:3}
        m = int(n / 5)
        cnt = 0

        for im in range(0, m+1):
            d = int(im / 5)
            for _d in range(0, d+1):
                cnt += int((im - 5*_d) / 2) + 1

        return cnt


if __name__ == "__main__":
    sol = Solution()
    res = sol.waysToChange(10)
    print("amount = 10, ways = %d, want=4" % res)
    res = sol.waysToChange(20)
    print("amount = 10, ways = %d, want=9" % res)
    res = sol.waysToChange(30)
    print("")
