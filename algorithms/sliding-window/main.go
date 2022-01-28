package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 1
	n := len(s)
	var start, end, length int
	length = math.MaxInt32

	for left < right && right < n+1 {
		c := s[right-1]
		right++
		window[c]++

		for isEqual(need, window) {
			if length > right-left {
				length = right-left
				start, end = left, right-1
			}
			b := s[left]
			window[b]--
			left++
		}
	}
	if start == end {
		return ""
	}
	return s[start:end-1]
}

func isEqual(need, window map[byte]int) bool {
	for k, v1 := range need {
		if v2, ok := window[k]; !ok {
			return false
		} else if v1 > v2 {
			return false
		}
	}
	return true
}
