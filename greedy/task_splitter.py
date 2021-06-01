from typing import Dict, List
import collections
from random import randrange

ruleDefination = collections.namedtuple('RULE', ['id', 'amount'])
# 允许的误差（最大允许超过的部分）期望 150 但是凑得 160 也可以接受
aollow_deviation = 10


def sort_key(item: ruleDefination):
    return item[1]


class rulePool(object):

    _amount: int = 0
    rules: List[ruleDefination] = []

    def __init__(self, amount: int, rules: List[ruleDefination]) -> None:
        super().__init__()
        self._amount = amount
        self.rules = rules

    def amount(self) -> int:
        return self._amount

    def random_one(self) -> ruleDefination:
        # random choose one item in
        idx = randrange(0, len(self.rules), 1)
        return self.rules[idx]


class Solution:

    def greedy(self, amount: int, pools: List[rulePool]) -> List[ruleDefination]:
        '''
        贪心算法：总是从最大的选取，直到
        '''
        print("begin: {} {}".format(amount, len(pools)))
        choice = []
        diff = amount
        cur = 0
        while cur < len(pools):
            if diff < pools[cur].amount():
                # setp on
                cur += 1
                continue

            c = pools[cur].random_one()
            diff = diff - c.amount
            choice.append(c)

        print("lastStep: {} {}".format(cur, diff))

        # 退回最后一个规则处
        cur -= 1
        while diff > 0:
            c = pools[cur].random_one()
            diff = diff - c.amount
            choice.append(c)

        return choice

    def call(self, amount: int, rules: List[ruleDefination]) -> List[ruleDefination]:
        rules = sorted(rules, key=sort_key, reverse=True)
        pools: List[rulePool] = []
        d: Dict[int, List[ruleDefination]] = {}

        for rule in rules:
            if rule.amount in d:
                d[rule.amount].append(rule)
            else:
                d[rule.amount] = [rule]

        # print(d)
        for k, v in d.items():
            pools.append(rulePool(k, v))

        choice = self.greedy(amount, pools)
        self.debug(amount, choice)

        return choice

    def debug(self, want: int, choice: List[ruleDefination]):
        got = 0
        diff = 0
        for c in choice:
            got += c.amount

        diff = want - got
        print("choice={}".format(choice))
        print("want={} got={} diff={}".format(want, got, diff))


if __name__ == "__main__":
    rules: List[ruleDefination] = [
        ruleDefination(id=1, amount=100),
        ruleDefination(id=9, amount=1000),
        ruleDefination(id=10, amount=500),
        ruleDefination(id=6, amount=100),
        ruleDefination(id=7, amount=100),
        ruleDefination(id=4, amount=5),
        ruleDefination(id=3, amount=10),
        ruleDefination(id=2, amount=50),
    ]

    sol = Solution()
    # choice = sol.call(1231, rules)
    # choice = sol.call(923, rules)
    # choice = sol.call(450, rules)
    choice = sol.call(125, rules)
    choice = sol.call(126, rules)
    choice = sol.call(127, rules)
    choice = sol.call(128, rules)
    choice = sol.call(129, rules)
    choice = sol.call(130, rules)
