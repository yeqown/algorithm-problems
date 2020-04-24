class Solution:
    ways = 0

    # TLE
    def canJump(self, nums: [int]) -> bool:
        n = len(nums) - 1
        if n <= 0:
            return True

        # reset
        self.ways = 0

        # idx = [0, n-1], (init) max_skip = nums[idx]
        def backtrack(idx: int, max_skip: int):
            if self.ways != 0:
                return

            # print(idx, max_skip, self.ways)
            if idx + max_skip >= n:
                self.ways += 1
                return

            if not max_skip:
                return

            for skip in range(max_skip, 0, -1):
                # do choice
                idx += skip
                if idx >= n:
                    self.ways += 1
                    continue
                backtrack(idx, nums[idx])
                idx -= skip

        backtrack(0, nums[0])
        # print(self.ways)
        return self.ways != 0

    # AC
    def canJump2(self, nums: [int]) -> bool:
        n = len(nums)
        rightmax = 0

        for idx in range(n):
            if idx <= rightmax:
                rightmax = max(rightmax, idx+nums[idx])
                if rightmax >= n - 1:
                    return True
        return False


if __name__ == "__main__":
    sol = Solution()
    arr = [2, 3, 1, 1, 4]
    res = sol.canJump(arr)
    print("can jump arr=%s, res=%s" % (arr, res))

    arr = [3, 2, 1, 0, 4]
    res = sol.canJump(arr)
    print("can jump arr=%s, res=%s" % (arr, res))

    arr = [2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0, 0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8,
           0, 7, 9, 3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5, 8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6]
    res = sol.canJump2(arr)
    print("can jump arr=%s, res=%s" % (arr, res))
