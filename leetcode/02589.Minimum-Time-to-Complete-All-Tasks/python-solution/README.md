# [2589. Minimum Time to Complete All Tasks](https://leetcode.com/problems/minimum-time-to-complete-all-tasks/)
## Description
There is a computer that can run an unlimited number of tasks at the same time. You are given a 2D integer array tasks where tasks[i] = [starti, endi, durationi] indicates that the ith task should run for a total of durationi seconds (not necessarily continuous) within the inclusive time range [starti, endi].

You may turn on the computer only when it needs to run a task. You can also turn it off if it is idle.

Return the minimum time during which the computer should be turned on to complete all tasks.

## Solution
Sort the range by end time and fill the duration from end to start.
