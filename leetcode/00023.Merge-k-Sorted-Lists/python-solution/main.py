class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Node:
    
    def __init__(self, data:ListNode):
        self.data = data
    
    def __lt__(self, other):
        return self.data.val < other.data.val

from heapq import *
from typing import List, Optional
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        
        h = []
        nh = ListNode(-1)
        now = nh
        for i in lists:
            if i:
                heappush(h, Node(i))
        while h:
            u = heappop(h)
            now.next = u.data
            now = now.next
            if u.data.next:
                heappush(h, Node(u.data.next))
        now.next =None
        return nh.next

if __name__ == "__main__":
    solu = Solution()
    pass
