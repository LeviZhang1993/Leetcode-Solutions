from typing import List


class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        l = 1
        r = max(piles)
        
        n = len(piles)
        while l <= r:
            mid = (l + r) //2
            tmp = sum(((i-1)//mid + 1) for i in piles)
            if tmp <= h:
                r = mid - 1
            else:
                l = mid + 1
        return l

if __name__ == "__main__":
    solu = Solution()
    print(solu.minEatingSpeed([3,6,7,11], 8) == 4)