package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//fmt.Println(compress([]byte{'a','b','c','c','c','c','c','c'}))
	//fmt.Println(getMaximumGenerated(15))
	fmt.Println(findCheapestPrice(3, [][]int{{0,1,100},{1,2,100},{0,2,500}}, 0, 2, 1))
}

func compress(chars []byte) int {
	left, right, write := 0, 0, 0
	for right < len(chars) {
		// 需要统计字符
		if right == len(chars) - 1 || chars[right] != chars[right+1] {
			count := right - left + 1
			chars[write] = chars[right]
			write++
			if count > 1 {
				sc := strconv.Itoa(count)
				for i := range sc {
					chars[write] = sc[i]
					write++
				}
			}
			left = right + 1
		}
		right++
	}
	fmt.Println(string(chars))
	return write
}

/*
https://leetcode-cn.com/problems/get-maximum-in-generated-array/
nums[0] = 0
nums[1] = 1
当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]
 */

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	nums := make([]int, n+1)
	if n + 1 >= 2 {
		nums[0] = 0
		nums[1] = 1
	}
	res := 1
	for i := 2; i < n+1; i++ {
		if i % 2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/
输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
输出: 200
 */
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// 根据题目中给出的数据范围，航班的花费不超过 10^4，最多搭乘航班的次数 k+1 不超过 101
	const inf = 10000*101 + 1
	f := make([][]int, k+2)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][src] = 0
	for t := 1; t <= k+1; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			f[t][i] = min(f[t][i], f[t-1][j]+cost)
		}
	}
	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, f[t][dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
