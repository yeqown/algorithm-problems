
# fib(1) = 1
# fib(2) = 1
# fib(3) = fib(2) + fib(1)
# fib(4) = fib(3) + fib(2)
# fib(n) = fib(n-1) + fib(n-2)

import time
from functools import wraps


# def fn_timer(function):
#     @wraps(function)
#     def function_timer(*args, **kwargs):
#         t0 = time.time()
#         result = function(*args, **kwargs)
#         t1 = time.time()
#         print("%s running: %f millseconds" %
#               (function, (t1-t0) * 1000))
#         return result
#     return function_timer

def fib_dp(n):
        # 时间复杂度：O(n) * O(1)
        # 0, 1, 2
    dp_table = [0, 1, 1]
    for i in range(n+1):
        if i <= 2:
            continue

        # print(i, dp_table[i-1], dp_table[i-2])
        dp_table.append(dp_table[i-1] + dp_table[i-2])
        # dp_table[i] = dp_table[i-1] + dp_table[i-2]

    # print(dp_table)
    return dp_table[n]


def fib(n):
    # 子问题个数 * 单个子问题的时间成本
    # 时间复杂度：O(2^n) * O(1)
    if n <= 2:
        return 1
    return fib(n-1) + fib(n-2)


class fibWithRemark(object):
    # 时间复杂度: O(n) * O(1)

    remark = {
        1: 1,
        2: 1,
    }

    def fib(self, n):
        if n in self.remark:
            return self.remark[n]
        self.remark[n] = self.fib(n-1) + self.fib(n-2)
        # print(self.remark, n)
        return self.remark[n]


if __name__ == "__main__":
    N = 100

    t0 = time.time()
    res = fib_dp(N)
    t1 = time.time()
    print(res, (t1-t0) * 1000, "ms")

    # t0 = time.time()
    # res2 = fib(N)
    # t1 = time.time()
    # print(res2, (t1-t0) * 1000, "ms")

    t0 = time.time()
    res3 = fibWithRemark().fib(N)
    t1 = time.time()
    print(res3, (t1-t0) * 1000, "ms")

    # N = 30
    # 832040 0.017881393432617188 ms    [dp]
    # 832040 338.3972644805908 ms       [recursive]
    # 832040 0.027179718017578125 ms    [with remark]

    # N = 100
    # 354224848179261915075 0.03695487976074219 ms  [dp]
    # 354224848179261915075 0.1590251922607422 ms   [with remark]
