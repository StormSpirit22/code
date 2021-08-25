package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//examples.Example_priorityQueue()
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
}

/*
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
*/
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	var res int
	for i := range nums {
		heap.Push(h, nums[i])
	}
	for i := 0; i < k; i++ {
		res = heap.Pop(h).(int)
	}
	return res
}

// IntHeap 是一个由整数组成的最大堆。
type IntHeap []int

func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}


func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}


func nthSuperUglyNumber(n int, primes []int) (ugly int) {
	seen := map[int]bool{1: true}
	h := &hp{[]int{1}}
	for i := 0; i < n; i++ {
		ugly = heap.Pop(h).(int)
		for _, prime := range primes {
			if next := ugly * prime; !seen[next] {
				seen[next] = true
				heap.Push(h, next)
			}
		}
	}
	return
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }


