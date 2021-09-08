package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(smallestK([]int{1,2,3}, 4))
	//fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(findMaximizedCapital(2,0,[]int{1,2,3},[]int{0,1,1}))
}

type ListNode struct {
    Val int
    Next *ListNode
}

/*
https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	tmp := head
	for i := 0; i < k; i++ {
		tmp = tmp.Next
	}
	for tmp != nil {
		head = head.Next
		tmp = tmp.Next
	}
	return head
}

/*
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	h := hp{}
	var i int
	for ; i < k; i++ {
		heap.Push(&h, arr[i])
	}

	for ; i < len(arr); i++ {
		tmp := heap.Pop(&h).(int)
		if tmp > arr[i] {
			heap.Push(&h, arr[i])
		} else {
			heap.Push(&h, tmp)
		}
	}
	var ans []int
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(&h).(int))
	}
	sort.Ints(ans)
	return ans
}

type hp []int

func(h hp) Len() int { return len(h) }
func(h hp) Swap(a, b int) { h[a], h[b] = h[b], h[a] }
func(h hp) Less(a, b int) bool { return h[a] > h[b] }
func(h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func(h *hp) Pop() interface{} {
	old := *h
	ans := old[len(old)-1]
	*h = old[:len(old)-1]
	return ans
}
*/
/*
https://leetcode-cn.com/problems/smallest-k-lcci/
官方答案
 */
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

/**
https://leetcode-cn.com/problems/implement-rand10-using-rand7/
古典概型
1. 第一次rand7限定[1,6]，判断奇偶性，概率是1/2
2. 第二次rand7限定[1,5]，概率是1/5
3. 二者结合可以得出10种概率相同的结果
如果是rand11那么就生成 [1, 6] 和 [7, 12]，拒绝 12。
medium
*/
//func rand10() int {
//	a, b := rand7(), rand7()
//	for a > 6 {
//		a = rand7()
//	}
//	for b > 5 {
//		b = rand7()
//	}
//	if a & 1 == 0 {
//		return b
//	}
//	return 5 + b
//}

/*
https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
easy
 */
func balancedStringSplit(s string) int {
	if len(s) == 0 {
		return 0
	}
	var count int
	d := 0
	for i := range s {
		if s[i] == 'L' {
			d++
		} else {
			d--
		}
		if d == 0 {
			count++
		}
	}
	return count
}

/*
https://leetcode-cn.com/problems/ipo/
优先队列
 */
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	p := PQ{}
	n := len(capital)
	projects := make([]*Project, n)
	for i := 0; i < n; i++ {
		projects[i] = &Project{capital: capital[i], profit: profits[i]}
	}
	// 先根据 capital 从小到大排序，这样就可以把小于 w 的一次性找出来
	sort.Slice(projects, func(i, j int) bool { return projects[i].capital < projects[j].capital })
	var cur int
	for ; k > 0; k-- {
		// 将符合要求的项目全部压入大根堆中，然后出栈利润最高的那个项目
		for cur < n && projects[cur].capital <= w {
			heap.Push(&p, projects[cur])
			cur++
		}
		if p.Len() == 0 {
			return w
		}
		w += heap.Pop(&p).(*Project).profit
	}
	return w
}

type PQ []*Project

type Project struct {
	capital int
	profit int
}

func (p PQ) Len() int { return len(p) }
func (p PQ) Swap(i, j int) { p[i],p[j] = p[j],p[i] }
func (p PQ) Less(i, j int) bool { return p[i].profit > p[j].profit }
func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.(*Project))
}
func (p *PQ) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}