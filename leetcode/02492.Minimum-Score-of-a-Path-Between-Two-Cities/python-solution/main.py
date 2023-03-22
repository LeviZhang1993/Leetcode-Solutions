from collections import defaultdict, deque
from typing import List


class Solution:
    def minScore(self, n: int, roads: List[List[int]]) -> int:
        
        nxt = defaultdict(list)
        for u, v, d in roads:
            nxt[u].append((v, d))
            nxt[v].append((u, d))
        
        dis = [1<<31] * (1 + n)
        q = deque()
        q.append(1)
        while q:
            u = q.popleft()
            for v, d in nxt[u]:
                if dis[v] > min(dis[u], d):
                    dis[v] = min(dis[u], d)
                    q.append(v)
        return dis[-1]

if __name__ == "__main__":
    solu = Solution()
