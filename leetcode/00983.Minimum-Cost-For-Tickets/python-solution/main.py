from typing import List


class Solution:
    def mincostTickets(self, days: List[int], costs: List[int]) -> int:
        n = len(days)
        days = [-1000] + days
        dp = [0] * (n + 1)
        for i in range(1, n + 1):
            dp[i] = dp[i-1] + costs[0]
            for idx, d in [(1, 7), (2, 30)]:
                c = costs[idx]
                j = i - 1
                while days[j] > days[i] - d:
                    j -= 1
                dp[i] = min(dp[j] + c, dp[i])
        return dp[-1]

if __name__ == "__main__":
    solu = Solution()
