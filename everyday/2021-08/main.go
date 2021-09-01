package main

import (
	"fmt"
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

//func sumOddLengthSubarrays(arr []int) int {
//	var sum int
//	n := len(arr)
//	prefixSum := make([]int, n+1)
//	for i := 0; i < n; i++ {
//		prefixSum[i+1] = prefixSum[i] + arr[i]
//	}
//	fmt.Println(prefixSum)
//	for i := 0; i < n; i++ {
//		for length := 1; length + i <= n; length += 2 {
//			sum += prefixSum[length+i] - prefixSum[i]
//		}
//	}
//	return sum
//	heap.Fix()
//}

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