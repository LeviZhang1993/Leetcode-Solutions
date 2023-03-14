# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        self.ans = 0
        if not root:
            return 0
        
        def dfs(r, now):
            if not r:
                return
            if not r.left and not r.right:
                self.ans += now
                return
            if r.left:
                dfs(r.left, now * 10 + r.left.val)
            if r.right:
                dfs(r.right, now * 10 + r.right.val)
        dfs(root, root.val)
        return self.ans

if __name__ == "__main__":
    solu = Solution()
    pass
