from collections import Counter
from functools import lru_cache

class Solution:

    @lru_cache(None)
    def isScramble(self, s1: str, s2: str) -> bool:
        if not s1 and not s2:
            return True
        if len(s1) != len(s2):
            return False
        if s1 == s2:
            return True
        c = Counter()
        for i in range(len(s1)):
            c[s1[i]] += 1
            c[s2[i]] -= 1
        if any(c[j] != 0 for j in c):
            return False
        for i in range(1, len(s1)):
            if self.isScramble(s1[:i], s2[-i:]) and self.isScramble(s1[i:], s2[:-i]):
                return True
            if self.isScramble(s1[:i], s2[:i]) and self.isScramble(s1[i:], s2[i:]):
                return True
        return False

if __name__ == "__main__":
    solu = Solution()
