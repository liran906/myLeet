# 归并排序
import random

def merge_sort(alist):
    if len(alist) <= 1:
        return alist
    mid = len(alist) // 2
    left = merge_sort(alist[:mid])
    right = merge_sort(alist[mid:])
    return merge(left, right)

def merge(left, right):
    res = []
    while left and right:
        if left[0] <= right[0]: # 加上等于号保证是稳定排序
            res.append(left[0])
            left = left[1:]
        else:
            res.append(right[0])
            right = right[1:]
    res.extend(left)
    res.extend(right)
    return res

if __name__ == '__main__':
    # 生成随机数
    nums = [random.randint(0, 100) for i in range(30)]
    print(merge_sort(nums))