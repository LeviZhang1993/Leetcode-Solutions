# [2585. Number of Ways to Earn Points](https://leetcode.com/problems/number-of-ways-to-earn-points/)
## Description
There is a test that has n types of questions. You are given an integer target and a 0-indexed 2D integer array types where types[i] = [counti, marksi] indicates that there are counti questions of the ith type, and each one of them is worth marksi points.

Return the number of ways you can earn exactly target points in the exam. Since the answer may be too large, return it modulo 109 + 7.
## Solution
Dynamic programming. Dp[i][j] means how many ways to reach j value from [0-i] range.
