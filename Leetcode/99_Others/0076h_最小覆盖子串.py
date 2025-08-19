# https://leetcode.cn/problems/minimum-window-substring/description/

class Solution:
    # O（n^2）
    def minWindow(self, s: str, t: str) -> str:
        htable = dict()
        for ch in t:
            htable[ch] = htable.get(ch, 0) + 1
        
        ans = ""
        l = 0
        while l <= len(s)-len(t):
            if s[l] in htable:
                r = l
                temptable = dict()
                while r < len(s) and htable != temptable:
                    if s[r] in htable:
                        temptable[s[r]] = temptable.get(s[r], 0) + 1
                    r += 1
                if htable == temptable: # 也可以不等于，所以这里有问题
                    if ans == "":
                        ans = s[l:r]
                    elif r - l < len(ans):
                        ans = s[l:r]
            l += 1
        return ans
    

# O(n)
from collections import Counter, defaultdict

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if not s or not t:
            return ""

        need = Counter(t)           # 需求计数
        have = defaultdict(int)     # 当前窗口计数
        need_kinds = len(need)      # 需要满足的不同字符种类数
        formed = 0                  # 已满足需求（计数达到或超过）的种类数

        ans_len = float('inf')
        ans_l = ans_r = 0

        l = 0
        for r, ch in enumerate(s):
            # 扩张右边界，更新窗口内计数
            if ch in need:
                have[ch] += 1
                # 刚好达到需求阈值时，formed+1（超过不再+1）
                if have[ch] == need[ch]:
                    formed += 1

            # 当窗口已经覆盖全部需求字符时，尝试收缩左边界
            while formed == need_kinds:
                # 更新答案
                if r - l + 1 < ans_len:
                    ans_len = r - l + 1
                    ans_l, ans_r = l, r

                # 收缩左边界
                left_ch = s[l]
                if left_ch in need:
                    have[left_ch] -= 1
                    # 一旦低于需求，formed 减 1，窗口不再有效
                    if have[left_ch] < need[left_ch]:
                        formed -= 1
                l += 1

        return "" if ans_len == float('inf') else s[ans_l:ans_r+1]