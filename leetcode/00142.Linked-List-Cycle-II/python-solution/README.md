# [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/description/)
## Description
Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.
## Solution
Just simulate the behavior. Simple math can proof the theory of quick/slow pointer.
Assume we have n nodes and k nodes forms a cycle. h = n - k (means the head of cycle)
if we have s steps when quick and slow pointers meet, then we have 2 * s - s = k * m(m is a int)
then current step can be divided by k, now if we move h steps we will have s + h steps, moved to 
the start of cycle.