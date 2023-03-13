from collections import Counter
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def kthLargestLevelSum(self, root: Optional[TreeNode], k: int) -> int:
        max_lvl = 0
        c = Counter()
        
        def dfs(r, lvl):
            nonlocal max_lvl
            if not r:
                return 
            max_lvl = max(lvl, max_lvl)
            c[lvl] += r.val
            for i in [r.left, r.right]:
                dfs(i, lvl + 1)
        dfs(root, 1)
        arr = [c[i] for i in c]
        if max_lvl < k:
            return -1
        arr.sort(reverse=True)
        return arr[k-1]

if __name__ == "__main__":
    solu = Solution()
    pass
