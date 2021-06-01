"""
01背包问题
有N件物品和一个容量为V的背包。第i建物品的费用是c[i],价值是w[i]。
求解将哪些物品装入背包可使价值总和最大（每件物品只能被使用一次）
"""

from typing import List, NamedTuple
from collections import namedtuple

item: NamedTuple = namedtuple('item', ['cost', 'weight'])


class Solution(object):

    def dp(self):
        pass

    def resolve(self, items: List[item], V: int):
        return self.dp(items, V)


if __name__ == "__main__":
    items = [
        item(cost=4, weight=5),
        item(1, 2),
        item(2, 4),
        item(3, 4),
        item(4, 5),
    ]

    sol = Solution()
    sol.resolve(items, 8)
