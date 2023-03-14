from typing import List

class BIT:
    
    def __init__(self, n):
        self.a = [0] * (n + 1)
        self.n = n
        
    def lowbit(self, x):
        return x & (-x)
    
    def update(self, k, v):
        while k <= self.n:
            self.a[k] += v
            k += self.lowbit(k)
    
    def query(self, k):
        ans = 0
        while k:
            ans += self.a[k]
            k -= self.lowbit(k)
        return ans
    
    def query_range(self, l, r):
        return self.query(r) - self.query(l - 1)
        

class Solution:
    def findMinimumTime(self, tasks: List[List[int]]) -> int:
        ma = 0
        for l, r, d in tasks:
            ma = max(ma, r)
            
        t = BIT(ma)
        
        tasks.sort(key=lambda x: x[1])
        arr = [0] * (ma + 1)
        for l, r, d in tasks:
            curr = t.query_range(l, r)
            idx = r
            while curr < d:
                if arr[idx]:
                    idx -= 1
                    continue
                arr[idx] = 1
                t.update(idx, 1)
                curr += 1
        return t.query_range(1, ma)
            
if __name__ == "__main__":
    solu = Solution()
