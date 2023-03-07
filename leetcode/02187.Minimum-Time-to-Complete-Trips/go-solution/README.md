# [2187. Minimum Time to Complete Trips](https://leetcode.com/problems/minimum-time-to-complete-trips/description/)
## Description
You are given an array time where time[i] denotes the time taken by the ith bus to complete one trip.

Each bus can make multiple trips successively; that is, the next trip can start immediately after completing the current trip. Also, each bus operates independently; that is, the trips of one bus do not influence the trips of any other bus.

You are also given an integer totalTrips, which denotes the number of trips all buses should make in total. Return the minimum time required for all buses to complete at least totalTrips trips.
## Solution
It's obvious that we get a value ans in which all the cases < ans we can't have toatl trips but all cases >= ans we can finish total trips. It's a very typical quetion for querying upper bound of a given condition.
