from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return None
        q = head.next.next
        s = head.next
        while q and q.next and s != q:
            q = q.next.next
            s = s.next
        if not q or not q.next:
            return None
        h = head
        while h and h != q:
            h = h.next
            q = q.next
        return q

if __name__ == "__main__":
    so = Solution()
    node1, node2, node3, node4 = ListNode(3), ListNode(2), ListNode(0), ListNode(-4)
    node1.next = node2
    node2.next = node3
    node3.next = node4
    node4.next = node2
    print(so.detectCycle(node1) == node2)