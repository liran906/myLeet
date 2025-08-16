package main

import (
	"container/heap"
	"sort"
)

// 方法1: 最佳实现，用最小堆

type ListNode struct {
	Val  int
	Next *ListNode
}

// ListHeap 实现 heap.Interface 的最小堆
type ListHeap []*ListNode

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// mergeKLists 主函数
func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListHeap{}
	heap.Init(h)

	// 把每个链表头加入堆
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	dummy := &ListNode{}
	cur := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		cur = cur.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

// 方法2:排序
func mergeKLists_(lists []*ListNode) *ListNode {
	// 先过滤掉 nil 的链表，避免空指针
	sortList := make([]*ListNode, 0)
	for _, l := range lists {
		if l != nil {
			sortList = append(sortList, l)
		}
	}

	// 如果所有链表都是空的，直接返回 nil
	if len(sortList) == 0 {
		return nil
	}

	// 将所有非空链表的头节点按值升序排序
	sort.Slice(sortList, func(i, j int) bool {
		return sortList[i].Val < sortList[j].Val
	})

	// 创建一个 dummy 节点作为返回链表的头结点
	dummy := &ListNode{}
	cur := dummy

	// 每轮从 sortList 中取出当前最小的节点处理
	for len(sortList) > 1 {
		uNode := sortList[0]    // 最小值节点
		sortList = sortList[1:] // 从列表中移除它
		cur.Next = uNode        // 接到结果链表中
		uNode = uNode.Next      // uNode 向后移动一位

		// 如果该链表还有下一个节点，则重新插入排序列表中
		if uNode != nil {
			inserted := false
			// 按值插入到合适的位置，保持 sortList 升序
			for i := 0; i < len(sortList); i++ {
				if uNode.Val < sortList[i].Val {
					sortList = append(sortList[:i], append([]*ListNode{uNode}, sortList[i:]...)...)
					inserted = true
					break
				}
			}
			// 如果 uNode 的值最大，追加到末尾
			if !inserted {
				sortList = append(sortList, uNode)
			}
		}

		// 继续处理下一个节点
		cur = cur.Next
	}

	// 最后还剩一个链表节点，直接挂接到末尾
	cur.Next = sortList[0]
	return dummy.Next
}
