from functools import lru_cache


class Solution:
    def minInsertions(self, s: str) -> int:
        n = len(s)

        @lru_cache(None)
        def get(i, j):
            if i < 0:
                return n - j
            if j >= n:
                return i + 1
            if s[i] == s[j]:
                return get(i - 1, j + 1)
            else:
                return min(get(i - 1, j), get(i, j + 1)) + 1

        ans = n
        for i in range(n):
            ans = min(ans, get(i - 1, i + 1))
            ans = min(ans, get(i, i + 1))
        return ans


if __name__ == "__main__":
    solu = Solution()
    print(solu.minInsertions("zzazz") == 0)
    print(solu.minInsertions("mbadm") == 2)
