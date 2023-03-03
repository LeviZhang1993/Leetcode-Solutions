class Solution:
    def strStr(self, s: str, h: str) -> int:
        
        def find_next(p):
            m = len(p)
            nxt = [-1] * m
            i, j = 0, -1
            while i < m - 1:
                if j == -1 or p[i] == p[j]:
                    i += 1
                    j += 1
                    nxt[i] = j
                else:
                    j = nxt[j]
            return nxt
        if not h:
            return -1
        nxt = find_next(h)
        n = len(s)
        m = len(h)
        i = 0
        j = 0
        while i < n:
            if j == -1 or s[i] == h[j]:
                i += 1
                j += 1
            else:
                j = nxt[j]
            if j == m:
                return i - m
        return -1

if __name__ == "__main__":
    solu = Solution()
    print(solu.strStr("abcde", "bc") == 1)
    print(solu.strStr("abcde", "") == -1)
    print(solu.strStr("abababcde", "ab") == 0)

