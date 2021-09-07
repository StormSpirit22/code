package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(compress([]byte{'a','b','c','c','c','c','c','c'}))
	//fmt.Println(getMaximumGenerated(15))
	//fmt.Println(findCheapestPrice(3, [][]int{{0,1,100},{1,2,100},{0,2,500}}, 0, 2, 1))
	//fmt.Println(allPathsSourceTarget([][]int{{4,3,1},{3,2,4},{3},{4},{}}))
	//fmt.Println(numRescueBoats([]int{1,2}, 3))
	//fmt.Println(sumOddLengthSubarrays([]int{1,4,2,5,3}))
	fmt.Println(compareVersion("1.0.1", "1.0.0"))
}

/*
https://leetcode-cn.com/problems/string-compression/
medium
 */
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

/*
https://leetcode-cn.com/problems/all-paths-from-source-to-target/
输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
 */
// 回溯法
func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	n := len(graph)

	var backtrack func([]int, int)

	backtrack = func(track []int, start int) {
		fmt.Println(track)
		if len(track) > 0 && track[len(track)-1] == n-1 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(graph[start]); i++ {
			track = append(track, graph[start][i])
			backtrack(track, graph[start][i])
			track = track[:len(track)-1]
		}
	}
	backtrack([]int{0}, 0)
	return res
}

/*
https://leetcode-cn.com/problems/boats-to-save-people/
输入：people = [3,5,3,4], limit = 5
输出：4
解释：4 艘船分别载 (3), (3), (4), (5)
 */
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	boats := 0
	for left <= right {
		if people[right] + people[left] <= limit {
			left++
		}
		right--
		boats++
	}
	return boats
}

type MedianFinder struct {
	queMin, queMax mhp
}

/*
https://leetcode-cn.com/problems/find-median-from-data-stream
数据流的中位数
hard
 */

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int)  {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}


func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type mhp struct { sort.IntSlice }
func (h *mhp) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *mhp) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */


/*
https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays
medium
 */
func sumOddLengthSubarrays(arr []int) int {
	var sum int
	n := len(arr)
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + arr[i]
	}
	fmt.Println(prefixSum)
	for i := 0; i < n; i++ {
		for length := 1; length + i <= n; length += 2 {
			sum += prefixSum[length+i] - prefixSum[i]
		}
	}
	return sum
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	var i int
	for i = 0; i < len(v1); i++ {
		if i >= len(v2) {
			a, _ := strconv.Atoi(v1[i])
			if a > 0 {
				return 1
			}
		} else {
			a, _ := strconv.Atoi(v1[i])
			b, _ := strconv.Atoi(v2[i])
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
		}
	}
	if len(v1) < len(v2) {
		for ;i < len(v2); i++ {
			b, _ := strconv.Atoi(v2[i])
			if b > 0 {
				return -1
			}
		}
	}
	return 0
}

/*
https://leetcode-cn.com/problems/corporate-flight-bookings/
利用查分数组来求解。差分数组实际上就是求前缀和的逆运算。举个例子，对于一个数组 arr, 某个区间 [l, r] 需要增加 x，
可以求出来其对应的差分数组 nums（nums[i] = arr[i] - arr[i-1]），然后 nums[l] += x， nums[r+1] -= x，
之后对 nums 求前缀和即可。这是对区间修改的简化操作，不需要去遍历该区间来修改原数组。
在本题中已经给出各个区间的修改值，初始从 0 开始计算，那么查分数组初始也就是 n 个 0。
由于是在 [1,n] 中求解，所以相对应的差分数组应该是 nums[l-1] += x， nums[r] -= x。但是 r 可能等于 n ，
而差分数组最大索引是 n-1，所以当 r == n 时不做修改，而且本来 r = n 时相对于原数组 arr[n] 也没有意义。
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, b := range bookings {
		nums[b[0]-1] += b[2]
		if b[1] < n {
			nums[b[1]] -= b[2]
		}
	}
	for i := range nums {
		if i > 0 {
			nums[i] += nums[i-1]
		}
	}
	return nums
}

/*
https://leetcode-cn.com/problems/random-pick-with-weight
很巧妙的一题，使用前缀和+二分法，
*/

type Solution struct {
	Prefix []int
}


func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}


func (s *Solution) PickIndex() int {
	// 因为 1 <= w[i] ，所以需要使用随机函数参数产生 [1, sum[n - 1]] 范围内的随机数
	// 而 rand.Intn 是产生 [0, sum[n-1]) 的随机数，所以要 + 1
	x := rand.Intn(s.Prefix[len(s.Prefix)-1]) + 1
	index := sort.SearchInts(s.Prefix, x)
	return index
}
