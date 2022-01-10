package main

import (
	"fmt"
	"sort"
)

func main() {
	//wd := Constructor()
	//wd.AddWord("a")
	//fmt.Println(wd.Search("."))
	//fmt.Println(wd.Search("bad"))
	//fmt.Println(wd.Search("bab"))
	//fmt.Println(wd.Search(".ad"))
	fmt.Println(shoppingOffers([]int{2,5}, [][]int{{3,0,5},{1,2,10}}, []int{3,2}))
}

type WordDictionary struct {
	isWord bool
	next map[byte]*WordDictionary
}


func Constructor() WordDictionary {
	n := make(map[byte]*WordDictionary)
	n1 := make(map[byte]*WordDictionary)
	next := &WordDictionary{isWord: false, next: n1}
	n['.'] = next
	wd := WordDictionary{isWord: false, next: n}
	return wd
}


func (this *WordDictionary) AddWord(word string)  {
	for i := range word {
		if this.next[word[i]] == nil {
			m := make(map[byte]*WordDictionary)
			wd := &WordDictionary{isWord: false, next: m}
			this.next[word[i]] = wd
		}
		this = this.next[word[i]]
	}
	this.isWord = true
}


func (this *WordDictionary) Search(word string) bool {
	var dfs func(*WordDictionary, string) bool
	dfs = func(wd *WordDictionary, word string) bool {
		if wd == nil {
			return false
		}
		if len(word) == 0 {
			if wd.isWord {
				return true
			} else {
				return false
			}
		}
		if word[0] != '.' {
			if wd.next[word[0]] == nil {
				return false
			}
			return dfs(wd.next[word[0]], word[1:])
		} else {
			for _, v := range wd.next {
				if v != nil && dfs(v, word[1:]) {
					return true
				}
			}
		}
		return false
	}
	return dfs(this, word)
}

/*
https://leetcode-cn.com/problems/shopping-offers/
medium
 */
func shoppingOffers(price []int, special [][]int, needs []int) int {

	var sums []int
	var traverse func([]int, int)
	traverse = func(remain []int, sum int) {
		backup := make([]int, len(remain))
		copy(backup, remain)

		for _, s := range special {
			used := true
			for i := 0; i < len(remain); i++ {
				if s[i] > remain[i] {
					used = false
					break
				}
				remain[i] -= s[i]
			}
			if used {
				sum += s[len(s)-1]
				traverse(remain, sum)
				sum -= s[len(s)-1]
			}
			copy(remain, backup)
		}
		for i := range remain {
			sum += remain[i] * price[i]
		}
		sums = append(sums, sum)
	}

	traverse(needs, 0)
	sort.Ints(sums)
	return sums[0]
}