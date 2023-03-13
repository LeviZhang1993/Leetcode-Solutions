from typing import List


class Solution:
    def maxScore(self, nums: List[int]) -> int:
        nums.sort(reverse=True)
        ans = 0
        curr = 0
        for i in nums:
            curr += i
            if curr <= 0:
                break
            ans += 1
        return ans
    
if __name__ == "__main__":
    solu = Solution()
