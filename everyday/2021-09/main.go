package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(smallestK([]int{1,2,3}, 4))
	//fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	//fmt.Println(findMaximizedCapital(2,0,[]int{1,2,3},[]int{0,1,1}))
	//fmt.Println(fullJustify([]string{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}, 20))
	//fmt.Println(findLongestWord("aaa", []string{"aaa","aa"}))
	//fmt.Println(findPeakElement([]int{3,2,1}))
	fmt.Println(findWords([][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}}, []string{"oath","pea","eat","rain","oathi","oathk","oathf","oate","oathii","oathfi","oathfii"}))
}

type ListNode struct {
	Val  int
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
hard
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
	profit  int
}

func (p PQ) Len() int           { return len(p) }
func (p PQ) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PQ) Less(i, j int) bool { return p[i].profit > p[j].profit }
func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.(*Project))
}
func (p *PQ) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	var line string
	var lenLine int
	var wrap []string
	for i := 0; i < len(words); i++ {
		m := len(words[i])
		if lenLine + m < maxWidth {
			wrap = append(wrap, words[i])
			lenLine += m + 1
		} else {
			if lenLine + m == maxWidth {
				tmp := ""
				for i := range wrap {
					tmp += wrap[i] + " "
				}
				tmp += words[i]
				line = tmp
			} else {
				i --
				lenLine -= len(wrap)
				gap := len(wrap) - 1
				var mod int
				var divide int
				if gap > 0 {
					mod = (maxWidth - lenLine) % gap
					divide = (maxWidth - lenLine) / gap
				}
				var length int
				var tmp bytes.Buffer
				count := 0
				for k := 0; k < len(wrap); k++ {
					tmp.WriteString(wrap[k])
					if count < mod {
						length = divide + 1
						count ++
					} else {
						length = divide
					}
					if length == 0 {
						length = maxWidth - tmp.Len()
					}
					for p := 0; p < length && tmp.Len() < maxWidth; p++ {
						tmp.WriteString(" ")
					}
				}
				line = tmp.String()
			}
			fmt.Println("line", line)
			lenLine = 0
			ans = append(ans, line)
			wrap = []string{}
		}
	}
	if len(wrap) > 0 {
		var tmp bytes.Buffer
		for i := range wrap {
			tmp.WriteString(wrap[i])
			if tmp.Len() < maxWidth {
				tmp.WriteString(" ")
			}
		}
		for tmp.Len() < maxWidth {
			tmp.WriteString(" ")
		}
		ans = append(ans, tmp.String())
	}
	return ans
}

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(x, y int) bool {
		if len(dictionary[x]) > len(dictionary[y]) {
			return true
		} else if len(dictionary[x]) < len(dictionary[y]) {
			return false
		}
		for i := range dictionary[x] {
			if dictionary[x][i] > dictionary[y][i] {
				return false
			} else {
				return true
			}
		}
		return true
	})
	fmt.Println(dictionary)
	for _, d := range dictionary {
		i, j := 0, 0
		for ; i < len(d); i++ {
			for j < len(s) && d[i] != s[j] {
				j++
			}
			j++
			if j > len(s) {
				break
			}
		}
		if i == len(d) {
			return d
		}
	}
	return ""
}

func findPeakElement(nums []int) int {
	n := len(nums)
	nums2 := make([]int, n+2)
	nums2[0] = math.MinInt64
	nums2[n+1] = math.MinInt64

	for i := 0; i < n; i++ {
		nums2[i+1] = nums[i]
	}
	left, right := 0, n+1
	mid := (left + right) / 2
	for !(nums2[mid] > nums2[mid-1] && nums2[mid] > nums2[mid+1]) {
		if nums2[mid] < nums2[mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
		if mid == 0 {
			mid ++
		}
		if mid == n+1 {
			mid = n
		}
	}
	return mid - 1
}

/*
https://leetcode-cn.com/problems/word-search-ii/
hard
前缀树 回溯
 */
func findWords(board [][]byte, words []string) []string {
	wMap := make(map[string]bool)
	for _, w := range words {
		wMap[w] = true
	}
	var res []string
	resMap := make(map[string]struct{})
	m, n := len(board), len(board[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var backtrack func([][]bool, [][]byte, int, int, bytes.Buffer)
	backtrack = func(visited [][]bool, board [][]byte, i, j int, track bytes.Buffer) {
		if track.Len() > 10 {
			return
		}
		if wMap[track.String()] {
			if _, ok := resMap[track.String()]; !ok {
				res = append(res, track.String())
				resMap[track.String()] = struct{}{}
				// return
			}
		}

		for _, d := range dirs {
			//fmt.Println(i, j)
			i, j = i + d[0], j + d[1]
			if i < 0 || i > m-1 || j < 0 || j > n-1 || visited[i][j] {
				i, j = i - d[0], j - d[1]
				continue
			}
			visited[i][j] = true
			track.WriteByte(board[i][j])
			backtrack(visited, board, i, j, track)
			visited[i][j] = false
			i, j = i - d[0], j - d[1]
			track.Truncate(track.Len()-1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var s bytes.Buffer
			s.WriteByte(board[i][j])
			visited[i][j] = true
			backtrack(visited, board, i, j, s)
			visited[i][j] = false
		}
	}
	sort.Strings(res)
	return res
}