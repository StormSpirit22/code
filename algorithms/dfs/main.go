package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(permuteUnique([]int{1,3,2}))
	//fmt.Println(subsets([]int{1,2,3}))
	fmt.Println(generateParenthesis(3))
}

func permuteUnique(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)
	n := len(nums)
	var backtrack func([]int, map[int]bool)
	backtrack = func(track []int, visited map[int]bool) {
		if n == len(track) {
			tmp := make([]int, n)
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		for i := range nums {
			if i > 0 && nums[i] == nums[i-1] && visited[i-1] {
				continue
			}
			if !visited[i] {
				track = append(track, nums[i])
				visited[i] = true
				backtrack(track, visited)
				track = track[:len(track)-1]
				visited[i] = false
			}
		}
	}
	visited := make(map[int]bool)
	backtrack([]int{}, visited)
	return res
}

func subsets(nums []int) [][]int {
	var res [][]int

	n := len(nums)
	var backtrack func([]int, int)
	backtrack = func(track []int, start int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)

		for i := start; i < n; i++ {
			track = append(track, nums[i])
			backtrack(track, i+1)
			track = track[:len(track)-1]
		}
	}
	backtrack([]int{}, 0)
	return res
}

func generateParenthesis(n int) []string {
	var res []string

	var backtrack func(int, int, string)
	backtrack = func(left, right int, track string) {
		if left < 0 || right < 0 {
			return
		}
		if left > right {
			return
		}
		if len(track) == n*2 {
			res = append(res, track)
		}

		track += "("
		backtrack(left-1, right, track)
		track = track[:len(track)-1]

		track += ")"
		backtrack(left, right-1, track)
		track = track[:len(track)-1]
	}
	backtrack(n, n, "")
	return res
}