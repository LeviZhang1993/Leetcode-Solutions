# [983. Minimum Cost For Tickets](https://leetcode.com/problems/minimum-cost-for-tickets/description/)
## Description
You have planned some train traveling one year in advance. The days of the year in which you will travel are given as an integer array days. Each day is an integer from 1 to 365.

Train tickets are sold in three different ways:

a 1-day pass is sold for costs[0] dollars,
a 7-day pass is sold for costs[1] dollars, and
a 30-day pass is sold for costs[2] dollars.
The passes allow that many days of consecutive travel.

For example, if we get a 7-day pass on day 2, then we can travel for 7 days: 2, 3, 4, 5, 6, 7, and 8.
Return the minimum number of dollars you need to travel every day in the given list of days.
## Solution
DP[i] means the minimum cost to cover [0-i].
dp[i] = min(dp[i-1] + cost[1], dp[k] + cost[7], dp[j] + cost[30]), j, k is the best place we end last pass.
