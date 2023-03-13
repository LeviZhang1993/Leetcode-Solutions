from collections import Counter
from typing import List


class Solution:
    def findValidSplit(self, nums: List[int]) -> int:
        pre = [-1] * (10 ** 6 + 1)
        cnt = 0
        for i in range(2, 10 ** 6 + 1):
            if pre[i] == -1:
                pre[i] = i
                for j in range(i * i, 10 ** 6 + 1, i):
                    if pre[j] == -1:
                        pre[j] = i
        curr = Counter()
        for i in nums:
            while i != 1:
                curr[pre[i]] +=1 
                i //= pre[i]
        now = Counter()
        idx = -1
        for i in nums[:-1]:
            idx += 1
            s = set()
            while i != 1:
                now[pre[i]] +=1 
                curr[pre[i]] -= 1
                if now[pre[i]] == 1 and curr[pre[i]] > 0:
                    cnt += 1
                if now[pre[i]] > 1 and curr[pre[i]] == 0:
                    cnt -= 1
                i //= pre[i]
            if not cnt:
                return idx
        return -1

if __name__ == "__main__":
    solu = Solution()
