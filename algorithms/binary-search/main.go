package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	//s := Constructor([]int{1,3})
	//fmt.Println(s.PickIndex())
	//fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 5))
	fmt.Println(shipWithinDays([]int{1,2,3,1,1}, 4))
}


/*
https://leetcode-cn.com/problems/koko-eating-bananas/
 */
func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	sort.Ints(piles)
	left, right := 1, piles[n-1]

	for left < right {
		mid := left + (right - left) / 2
		hours := getHours(mid, piles)
		fmt.Printf("mid %d, hours %d \n", mid, hours)
		if hours == h {
			right = mid
		} else if hours < h {
			// 时间少了说明速度快了，可以减慢一下速度，缩小 mid
			right = mid
		} else if hours > h {
			// 速度慢了，需要加快一下速度，多吃点，增大 mid
			left = mid + 1
		}
	}
	return left
}

func getHours(num int, piles []int) int {
	var sum int
	for _, p := range piles {
		for p > 0 {
			p -= num
			sum++
		}
	}
	return sum
}

/*
https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
这里二分搜索的最小值应该是 max(weights)，因为船至少要能载起来所有货物，不然有货物载不起来肯定不行。最大值就是 sum(weights)。
 */
func shipWithinDays(weights []int, days int) int {
	//sort.Ints(weights)
	var sum int
	var maxWeight int
	for _, w := range weights {
		sum += w
		if maxWeight < w {
			maxWeight = w
		}
	}
	left, right := maxWeight, sum
	for left < right {
		mid := left + (right - left) / 2
		d := getDays(mid, weights)
		fmt.Printf("mid %d, days %d, left %d, right %d \n", mid, d, left, right)
		if d == days {
			right = mid
		} else if d < days {
			right = mid
		} else if d > days {
			left = mid + 1
		}
	}
	return left
}

func getDays(ship int, weights []int) int {
	var days int
	for i := 0; i < len(weights); {
		days++
		count := weights[i]
		for count <= ship {
			i++
			if i == len(weights) {
				break
			}
			count += weights[i]
		}
	}
	return days
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
