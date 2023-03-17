class Node:
    def __init__(self, cnt):
        self.cnt = 0
        self.next = [None] * 26
        
class Trie:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.head = Node(0)
        self.x = "abcdefghijklmnopqrstuvwxyz"
        self.d = {self.x[i] : i for i in range(26)}
        

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        cur = self.head
        for i in range(len(word)):
            t = word[i]
            t_idx = self.d[t]
            if not cur.next[t_idx]:
                cur.next[t_idx] = Node(0)
            cur = cur.next[t_idx]
        cur.cnt += 1
        

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        cur = self.head
        for i in range(len(word)):
            t = word[i]
            t_idx = self.d[t]
            if not cur.next[t_idx]:
                return False
            cur = cur.next[t_idx]
        return cur.cnt >= 1
        

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        cur = self.head
        for i in range(len(prefix)):
            t = prefix[i]
            t_idx = self.d[t]
            if not cur.next[t_idx]:
                return False
            cur = cur.next[t_idx]
        return cur is not None

if __name__ == "__main__":
    solu = Solution()
