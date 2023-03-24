from typing import List


class Solution:
    def minReorder(self, n: int, connections: List[List[int]]) -> int:
        nxt = [[] for _ in range(n)]
        pre = [[] for _ in range(n)]
        for u, v in connections:
            nxt[u].append(v)
            pre[v].append(u)
        
        ans = 0
        vis = {}
        def dp(r):
            nonlocal ans
            vis[r] = 1
            for i in nxt[r]:
                if i not in vis:
                    ans += 1
                    dp(i)
            for i in pre[r]:
                if i not in vis:
                    dp(i)
        dp(0)
        return ans

if __name__ == "__main__":
    solu = Solution()
