from typing import List


class Solution1:
    def countSubarrays(self, nums: List[int], minK: int, maxK: int) -> int:
        n = len(nums)
        ans = 0
        if minK > maxK:
            return 0
        i = 0
        while i < n:
            if nums[i] < minK or nums[i] > maxK:
                i += 1
                continue
            j = i
            mi, ma = -1, -1
            while j < n and minK <= nums[j] <= maxK:
                if nums[j] == minK:
                    mi = j
                if nums[j] == maxK:
                    ma = j
                if ma != -1 and mi != -1:
                    ans += min(mi, ma) - i + 1
                j += 1
            i = j
        return ans
    

class Solution2:
    def countSubarrays(self, nums: List[int], minK: int, maxK: int) -> int:
        n = len(nums)
        
        def cal(arr, k, d):
            m = len(arr)
            cnt = m * (m + 1) // 2
            c = Counter()
            j = 0
            for i in range(m):
                c[arr[i]] += 1
                if arr[i] in d and c[arr[i]] == 1:
                    k -= 1
                while k < 1:
                    c[arr[j]] -= 1
                    if arr[j] in d and c[arr[j]] == 0:
                        k += 1
                    j += 1
                cnt -= i - j + 1
            return cnt
        i = 0
        ans = 0
        if minK > maxK:
            return 0
        while i < n:
            if nums[i] < minK or nums[i] > maxK:
                i += 1
                continue
            j = i
            while j < n and minK <= nums[j] <= maxK:
                j += 1
            ans += cal(nums[i:j], 2 if minK < maxK else 1, set([minK, maxK]))
            i = j
        return ans
    

if __name__ == "__main__":
    so1 = Solution1()
    so2 = Solution2()
    print(so1.countSubarrays([1,3,5,2,7,5], 1, 5) == 2)
    print(so2.countSubarrays([1,3,5,2,7,5], 1, 5) == 2)