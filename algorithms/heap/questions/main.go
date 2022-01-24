package main

import (
	"container/heap"
)

func main() {
	//findKthLargest([]int{3,2,1,5,6,4}, 2)
	topKFrequent([]int{1,1,1,2,2,3}, 2)
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
	}
	var res int
	for i := 0; i < h.Len(); i++ {
		res = heap.Pop(h).(int)
	}
	return res
}

// IntHeap 是一个由整数组成的最大堆。
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	p := &PQ{}
	numMap := make(map[int]int)
	for _, n := range nums {
		numMap[n]++
	}
	heap.Init(p)
	for n, v := range numMap {
		f := Frequency{value: n, priority: v}
		heap.Push(p, f)
	}
	var res []int
	for i := 0; i < k; i++ {
		x := heap.Pop(p).(Frequency)
		res = append(res, x.value)
	}
	return res
}

type Frequency struct {
	value int
	priority int
}

type PQ []Frequency

func (p PQ) Len() int {return len(p)}
func (p PQ) Less(i, j int) bool {return p[i].priority > p[j].priority}
func (p PQ) Swap(i, j int) {p[i], p[j] = p[j], p[i]}

func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.(Frequency))
}

func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}
