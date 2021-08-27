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
https://leetcode-cn.com/problems/kth-largest-element-in-an-array
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


/*
https://leetcode-cn.com/problems/super-ugly-number/
超级丑数
 */
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


/*
https://leetcode-cn.com/problems/ugly-number-ii/
丑数II
 */
func nthUglyNumber(n int) int {
	count := 1
	primes := []int{2, 3, 5}
	h := uhp{1}
	ans := 1
	visited := make(map[int]bool)
	for count <= n {
		ans = heap.Pop(&h).(int)
		if count == n {
			return ans
		}
		for _, p := range primes {
			tmp := ans * p
			if !visited[tmp] {
				heap.Push(&h, tmp)
				visited[tmp] = true
			}
		}
		count++
	}
	return ans
}

type uhp []int
func (h uhp) Len() int {return len(h)}
func (h uhp) Less(i, j int) bool {return h[i] < h[j]}
func (h uhp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *uhp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *uhp) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}


/*
https://leetcode-cn.com/problems/top-k-frequent-elements
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
 */
func topKFrequent(nums []int, k int) []int {
	pq := PQs{}
	numMap := make(map[int]int)
	for _, n := range nums {
		numMap[n] += 1
	}
	for k, v := range numMap {
		f := &Frequency{frequency: v, num: k}
		heap.Push(&pq, f)
	}
	var res []int
	for i := 0; i < k; i++ {
		f := heap.Pop(&pq).(*Frequency)
		res = append(res, f.num)
	}
	return res
}

type PQs []*Frequency
type Frequency struct {
	num int
	frequency int
}

func (p PQs) Len() int { return len(p)}
func (p PQs) Less(i, j int) bool { return p[i].frequency > p[j].frequency }
func (p PQs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p *PQs) Push(x interface{}) {
	*p = append(*p, x.(*Frequency))
}
func (p *PQs) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

