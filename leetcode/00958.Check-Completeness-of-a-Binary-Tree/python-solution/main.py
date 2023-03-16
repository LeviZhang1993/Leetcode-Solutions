# Definition for a binary tree node.
from collections import deque
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
class Solution:
    def isCompleteTree(self, root: Optional[TreeNode]) -> bool:
        q = deque()
        q.append(root)
        has_none = False
        while q:
            u = q.popleft()
            if u is None:
                has_none = True
            else:
                if has_none:
                    return False
                q.append(u.left)
                q.append(u.right)
        return True

if __name__ == "__main__":
    solu = Solution()
    pass
