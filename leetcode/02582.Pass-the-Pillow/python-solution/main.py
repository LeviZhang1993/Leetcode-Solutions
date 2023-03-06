class Solution:
    def passThePillow(self, n: int, t: int) -> int:
        t =  t % (2 * n - 2)
        if t == 0:
            return 1
        return 1 + t if t <= n -1 else 2 * n - t - 1

if __name__ == "__main__":
    solu = Solution()
    print(solu.passThePillow(3, 2) == 3)
    print(solu.passThePillow(4, 5) == 2)