'''
input : amount, rules
output: solution[ruleId...] which make solution has mininum rules in solution,
and sum of solution is eligible.
'''

import collections
from typing import NamedTuple
from typing import List
from copy import deepcopy

ruleDefination: NamedTuple = collections.namedtuple('RULE', ['id', 'amount'])
# 允许的误差（最大允许超过的部分）期望 150 但是凑得 160 也可以接受
aollow_deviation = 10


def sort_key(item: ruleDefination):
    return item[1]


class Solution:
    choice = []
    choices = []

    def solutionToChange(self, amount: int, rules: List[ruleDefination]):
        '''
        @amount means how much should the fn resolve into rules
        @rules indicate what‘s the scope of result should be in
        '''
        found = False
        perfect = False
        if amount <= 0 or len(rules) == 0:
            found = True
            # end condition
            if amount == 0:
                # 如果找到刚好合适的就直接返回
                perfect = True
                print('perfect_choice', self.choice)

            if amount < 0:
                print('got one solution:', self.choice)
                self.choices.append(deepcopy(self.choice))

            return self.choice, found, perfect

        for idx, rule in enumerate(rules):
            # do choice
            _tmp_amount = amount-rule[1]
            if _tmp_amount < -aollow_deviation:
                # 超过误差则不用考虑
                continue
            self.choice.append(rule[0])
            # 不允许用多次，就将其移除
            # 否则，可以不用移除或者根据计数移除
            _tmp_rules = rules[:idx] + rules[idx+1:]

            # backtrack
            choices, found, perfect = self.solutionToChange(
                amount=_tmp_amount, rules=_tmp_rules)
            if perfect is True:
                return choices, found, perfect

            # undo choice
            self.choice.remove(rule[0])

        return [], found, perfect

    def call(self, amount: int, rules: List[ruleDefination]):
        rules = sorted(rules, key=sort_key, reverse=True)
        rules_mapping = {}
        for r in rules:
            rules_mapping[r[0]] = r

        # print(rules_mapping)

        # 如果有完美的选择
        perfect_choice, found, perfect = self.solutionToChange(
            amount=amount, rules=rules)
        # print("===========result: ", perfect_choice, found, perfect)

        if perfect is True:
            print("has perfect solution: ", [
                  rules_mapping[i] for i in perfect_choice])
            return

        # 如果没有找到, 但是有不那么好的选择，就选用最接近的方案
        min_dif = float('inf')
        min_idx = 0
        if len(self.choices) != 0:
            for idx, c in enumerate(self.choices):
                # print([rules_mapping[i][1] for i in c])
                diff = sum([rules_mapping[i][1] for i in c])-amount
                if diff < min_dif:
                    min_dif = diff
                    min_idx = idx
                print("alternative one choice={} diff={} idx={}".format(c, diff, idx))
            print("has solution: ", [rules_mapping[i]
                  for i in self.choices[min_idx]])
            return

        # 压根不能解决
        print("could not resolve")
        return []


if __name__ == "__main__":
    rules: List[ruleDefination] = [
        ruleDefination(id=1, amount=100),
        ruleDefination(id=4, amount=5),
        ruleDefination(id=3, amount=10),
        ruleDefination(id=2, amount=50),
        ruleDefination(id=5, amount=1)
    ]
    sol = Solution()
    sol.call(104, rules)
    # 115 / res = [1, 4, 3]
    # 104 / res = [1, 4]
    # 154 / res = [1, 4, 2]
    # 120 / FALSE
    # 139 / FALSE
    # 140 / res = [1, 2]
