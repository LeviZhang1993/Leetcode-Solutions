from typing import List


class Solution:
    def vowelStrings(self, words: List[str], left: int, right: int) -> int:
        ans = 0
        for i in words[left:right+1]:
            if i[0] in "aeiou" and i[-1] in "aeiou":
                ans += 1
        return ans

if __name__ == "__main__":
    solu = Solution()
