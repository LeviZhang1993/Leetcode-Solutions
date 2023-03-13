class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        
        def is_match(l, r):
            if not l:
                return not r
            if not r:
                return False
            if l.val != r.val:
                return False
            return is_match(l.left, r.right) and is_match(l.right, r.left)
        return is_match(root.left, root.right)

if __name__ == "__main__":
    solu = Solution()
