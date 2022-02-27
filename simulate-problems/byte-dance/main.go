package main

import "fmt"

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	fmt.Printf("%+v", reverseKGroup(head, 2))
}

type ListNode struct {
	Val int
	Next *ListNode
}

/*
https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
hard
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var count int
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	cur = head
	var start int
	for i := 0; i < count/k; i++ {
		cur = reverseGroup(cur, start, k)
		start += k
	}
	return cur
}

func reverseGroup(head *ListNode, start, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur, next := dummy, dummy, dummy
	for i := 0; i < start; i++ {
		pre = pre.Next
	}
	cur = pre.Next
	next = pre.Next

	preNext := next
	for i := 0; i < k; i++ {
		preNext = next
		next = next.Next
	}
	preNext.Next = nil
	back := cur
	pre.Next = reverseEntireList(cur)
	back.Next = next
	return dummy.Next
}

func reverseEntireList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur, next := head, head.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}