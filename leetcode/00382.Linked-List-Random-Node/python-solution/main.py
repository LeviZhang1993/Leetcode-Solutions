# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
import random
from typing import Optional
class Solution:

    def __init__(self, head: Optional[ListNode]):
        """
        @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node.
        """
        self.h = head
        

    def getRandom(self) -> int:
        """
        Returns a random node's value.
        """
        ans = self.h
        k = 1
        n = 1
        head = self.h
        while head.next:
            head = head.next
            n += 1
            if random.randint(1, n) <= k:
                ans = head
        return ans.val

if __name__ == "__main__":
    solu = Solution()
    pass
