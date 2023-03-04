# [2444. Count Subarrays With Fixed Bounds](https://leetcode.com/problems/count-subarrays-with-fixed-bounds/description/)

## Description
Find the number of subarrays whose max=Maxk min=minK

## Solution

### Solution1 iteration by given definition
use two index i, j to represent the last seen value index of maxK and minK, we also record the last index k with value between minK and maxK, it's obvious that for current the good subarrarys whose right side is max[i, j] has min[i, j] - k + 1.By adding all the values, we can get the final result.

### Solution2 transform into a sliding window question
1: first filter all the value not in [mink, maxk], then we get several subarrays. And every array values lay between mink and maxk.
2: for each array, we can do a sliding window to find the number of subarray as NUM which have at most k - 1 values in set[mink, maxk]. Then we get the number by n*(n-1)/2 - NUM
