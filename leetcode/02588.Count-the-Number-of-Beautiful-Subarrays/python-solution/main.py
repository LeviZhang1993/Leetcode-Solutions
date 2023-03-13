from collections import Counter
from typing import List


class Solution:
    def beautifulSubarrays(self, nums: List[int]) -> int:
        d = Counter()
        d[0] = 1
        curr = 0
        ans = 0
        for i in nums:
            for j in range(20):
                if i & (1<<j):
                    curr ^= (1<<j)
            ans += d[curr]
            d[curr] += 1
        return  ans

if __name__ == "__main__":
    solu = Solution()
