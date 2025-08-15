# https://leetcode.cn/problems/integer-to-roman/?envType=study-plan-v2&envId=top-interview-150

class Solution:
    symbols = [
        (1000, "M"),
        (900, "CM"),
        (500, "D"),
        (400, "CD"),
        (100, "C"),
        (90, "XC"),
        (50, "L"),
        (40, "XL"),
        (10, "X"),
        (9, "IX"),
        (5, "V"),
        (4, "IV"),
        (1, "I"),
    ]
  
    def intToRoman(self, num: int):
        res = ""
        for s in Solution.symbols:
            while num >= s[0]:
                num -= s[0]
                res += s[1]
            if num == 0:
                break
        return res
        