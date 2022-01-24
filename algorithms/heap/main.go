package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := &IntHeap{1, 3, 5}
	heap.Init(h)
	heap.Push(h, 2)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}}

// IntHeap 最大堆模板
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


type Frequency struct {
	value int
	priority int
}

// PQ 优先队列模板代码
type PQ []Frequency

func (p PQ) Len() int { return len(p) }
func (p PQ) Less(i, j int) bool { return p[i].priority < p[j].priority }
func (p PQ) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

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