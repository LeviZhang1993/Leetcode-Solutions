# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def sortedListToBST(self, head: ListNode) -> TreeNode:
        if not head:
            return None
        if not head.next:
            return TreeNode(head.val)
        slow = quick = pre = head
        while quick and quick.next:
            pre = slow
            slow = slow.next
            quick = quick.next.next
        pre.next = None
        return TreeNode(slow.val,
                       self.sortedListToBST(head),
                    self.sortedListToBST(slow.next),
                       )

if __name__ == "__main__":
    solu = Solution()
    pass
