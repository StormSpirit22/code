package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	s := Constructor([]int{1,3})
	fmt.Println(s.PickIndex())
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
