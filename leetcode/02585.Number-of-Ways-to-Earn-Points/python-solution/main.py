from typing import List


class Solution:
    def waysToReachTarget(self, t: int, types: List[List[int]]) -> int:
        dp = [0] * (t + 1)
        mod = 10 ** 9 + 7
        dp[0] = 1
        for c, m in types:
            curr = dp.copy()
            for i in range(1, c+1):
                if i * m > t:
                    break
                for j in range(t, i* m - 1, -1):
                    curr[j] = (curr[j] + dp[j-i*m]) % mod
            dp = curr.copy()
        return dp[t]

if __name__ == "__main__":
    solu = Solution()
    print(solu.waysToReachTarget(6, [[6,1],[3,2],[2,3]]) == 7)
