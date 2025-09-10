# 快速排序
import random


def quick_sort(alist, start=0, end=None):
    if end is None:
        end = len(alist) - 1
    if start >= end:
        return None
    pivot = alist[start]
    l = start + 1
    r = end
    while l <= r:
        while l <= r and alist[l] < pivot:
            l += 1
        while l <= r and alist[r] > pivot:
            r -= 1
        if l <= r:
            alist[l], alist[r] = alist[r], alist[l]
            l += 1
            r -= 1
    alist[start], alist[r] = alist[r], alist[start]
    quick_sort(alist, start, r - 1)
    quick_sort(alist, r + 1, end)
    return alist

if __name__ == '__main__':
    # 生成随机数
    nums = [random.randint(0, 100) for i in range(30)]
    print(quick_sort(nums))