from typing import List


class Solution:
    def minimumTime(self, time: List[int], totalTrips: int) -> int:
        l = 0
        r = min(time) * totalTrips
        while l <= r:
            mid = (l + r) // 2
            su = sum(mid // i for i in time)
            if su >= totalTrips:
                r = mid - 1
            else:
                l = mid + 1
        return l

if __name__ == "__main__":
    solu = Solution()
    print(solu.minimumTime([1,2,3], 5) == 3)
    print(solu.minimumTime([2], 1) == 2)
