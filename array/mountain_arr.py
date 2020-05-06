# """
# This is MountainArray's API interface.
# You should not implement it, or speculate about its implementation
# """
# class MountainArray:
#    def get(self, index: int) -> int:
#    def length(self) -> int:
#
# https://leetcode-cn.com/problems/find-in-mountain-array/
# idx = 1095


class MountainArray:
    data = []
    sz = 0

    def __init__(self, d: [int]):
        self.data = d
        self.sz = len(d)

    def get(self, index: int) -> int:
        if index < 0 or index >= self.length():
            return -1
        return self.data[index]

    def length(self) -> int:
        return self.sz


# MAX = float('INF')


class Solution:
    """
    利用二分搜索方法，找到“山峰”，再对两边二分搜索。
    """

    def binary_search(self, l: int, r: int, target: int, mountain_arr: 'MountainArray', asc: bool) -> int:
        while l <= r:
            mid = (l + r) // 2
            v = mountain_arr.get(mid)
            if v == target:
                return mid

            if v > target:
                # 升序, 找左边
                r = mid - 1 if asc else r
                # 降序，找右边
                l = mid + 1 if not asc else l
            else:
                # 升序, 找右边
                l = mid + 1 if asc else l
                # 降序，找左边
                r = mid - 1 if not asc else r
        return -1

    def findInMountainArray(self, target: int, mountain_arr: 'MountainArray') -> int:
        n = mountain_arr.length()
        l, r = 0, n-1
        peak = -1
        # find peak
        while l < r:
            mid = (l + r) // 2
            # v = mountain_arr.get(mid)
            # rv = mountain_arr.get(mid + 1)

            if mountain_arr.get(mid) < mountain_arr.get(mid + 1):
                # 升序，找右边
                l = mid + 1
            else:
                # 降序，找左边
                r = mid
            # lv = mountain_arr.get(mid - 1)
            # if v >= rv and v > lv:
            #     # peak
            #     peak = mid
            #     break
            # else if lv < v < rv:
            #     # asc
            #     l = mid + 1
            # else:
            #     # desc
            #     r = mid - 1
        peak = l
        print(peak)
        # 优先找左边
        idx = self.binary_search(0, peak, target, mountain_arr, True)
        if idx != -1:
            return idx

        return self.binary_search(peak + 1, n-1, target, mountain_arr, False)


if __name__ == "__main__":
    sol = Solution()
    arr = [1, 2, 3, 4, 5, 3, 1]
    ma = MountainArray(arr)
    idx = sol.binary_search(0, 4, 4, ma, True)
    print("arr=%s, search(0, 4)=%d, want=3" % (arr, idx))
    idx = sol.binary_search(4, len(arr)-1, 4, ma, False)
    print("arr=%s, search(0, 4)=%d, want=-1" % (arr, idx))
    idx = sol.binary_search(4, len(arr)-1, 3, ma, False)
    print("arr=%s, search(0, 4)=%d, want=5" % (arr, idx))
    idx = sol.binary_search(0, 4, 5, ma, True)
    print("arr=%s, search(0, 4)=%d, want=4" % (arr, idx))

    minIdx = sol.findInMountainArray(3, ma)
    print("sol.findInMountainArray(3, ma)=%d, want=2" % minIdx)
    minIdx = sol.findInMountainArray(5, ma)
    print("sol.findInMountainArray(5, ma)=%d, want=4" % minIdx)
