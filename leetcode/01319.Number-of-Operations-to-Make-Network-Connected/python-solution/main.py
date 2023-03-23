from typing import List


class Solution:
    def makeConnected(self, n: int, connections: List[List[int]]) -> int:
        f = [i for i in range(n)]
        if len(connections) < n - 1:
            return -1
        def find(x):
            if f[x] != x:
                f[x] = find(f[x])
            return f[x]
        for u, v in connections:
            f[find(u)] = find(v)
        
        return len(set([find(i) for i in range(n)])) - 1

if __name__ == "__main__":
    solu = Solution()
