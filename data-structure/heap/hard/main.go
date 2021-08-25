package main

import (
	"container/heap"
	"fmt"
)

func main() {
	lists := []*ListNode{nil, {Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}, {Val: -1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}}
	head := mergeKLists(lists)
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}
/*
https://leetcode-cn.com/problems/merge-k-sorted-lists/
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
*/

// 优先队列
func mergeKLists(lists []*ListNode) *ListNode {
	var pq PQ
	for i := range lists {
		if lists[i] != nil {
			heap.Push(&pq, lists[i])
		}
	}
	head := &ListNode{}
	cur := head
	for pq.Len() > 0 {
		tmpList := heap.Pop(&pq)
		cur.Next = tmpList.(*ListNode)
		cur = cur.Next
		if tmpList.(*ListNode).Next != nil {
			heap.Push(&pq, tmpList.(*ListNode).Next)
		}
	}
	return head.Next
}

type PQ []*ListNode

func (p PQ) Len() int { return len(p) }
func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PQ) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.(*ListNode))
}

func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

// 暴力法
func mergeKLists1(lists []*ListNode) *ListNode {
	var head *ListNode
	for i := range lists {
		head = mergeTwoLists(head, lists[i])
		cur := head
		for cur != nil {
			fmt.Printf("%d ", cur.Val)
			cur = cur.Next
		}
		fmt.Println()
	}
	return head
}


// 分治法
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func mergeTwoLists(left, right *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	a, b := left, right
	for a != nil && b != nil {
		if b.Val <= a.Val {
			cur.Next = b
			b = b.Next
		} else if b.Val > a.Val {
			cur.Next = a
			a = a.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	}
	if b != nil {
		cur.Next = b
	}
	return head.Next
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l > r {
		return nil
	}
	if l == r {
		return lists[l]
	}
	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}