# idx = 752
# 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字：
# '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
# 每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。
# 每次旋转都只能旋转一个拨轮的一位数字。
#
# 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
# 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
# 字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。


class Solution:
    def openLock(self, deadends: [str], target: str) -> int:
        return self.bfs(deadends, target)

    def toggle(self, s: str, i: int, change: int) -> str:
        return s[:i] + str((int(s[i]) + change) % 10) + s[i+1:]
        # n = []
        # for pos, c in enumerate(bytearray(s, encoding='utf-8')):
        #     if idx == pos:
        #         if c == 57 and change > 0:
        #             n.append(48)
        #         elif c == 48 and change < 0:
        #             n.append(57)
        #         else:
        #             n.append(c+1)
        #     else:
        #         n.append(c)
        # return bytearray(n).decode('utf-8')

    def bfs(self, deadends: [str], target: str) -> int:
        q = ['0000']
        visited = {}
        step = 0

        # little trick
        for deadend in deadends:
            visited[deadend] = True
        visited['0000'] = True

        while q:
            nq = len(q)
            for _ in range(nq):
                cur = q.pop(0)
                if cur in visited:
                    continue

                if cur == target:
                    return step
                # 将所有可能性加入q中
                for pos in range(4):
                    # 如果锁访问过，或者死亡密码就跳过
                    up = self.toggle(cur, pos, 1)
                    if up not in visited:
                        q.append(up)
                        visited[up] = True
                    # 如果锁访问过，或者死亡密码就跳过
                    down = self.toggle(cur, pos, -1)
                    if down not in visited:
                        q.append(down)
                        visited[down] = True
            step += 1
        return -1


if __name__ == "__main__":
    sol = Solution()

    # cur = sol.toggle('0000', 1, 1)
    # print('0000 toggle idx=1, up=true, want=0100, got=%s' % cur)

    # cur = sol.toggle('0900', 1, 1)
    # print('0200 toggle idx=1, up=true, want=0000, got=%s' % cur)

    res = sol.openLock(['0000'], '8888')
    print("open 8888 want minStep=-1, got=%d" % res)
