package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{0,0,0,0,0}))
}

/*
输入：nums = [-1,0,1,2,-1,-4]
-4 -1 -1 0 1 2
输出：[[-1,-1,2],[-1,0,1]]
先固定一个值，其他两个用双指针去不断逼近
 */
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var res [][]int
	for i := range nums {
		//枚举数字，如果和前一个相等那么 target 也相同，已经计算过一遍了直接跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			//枚举数字，如果和前一个相等，已经计算过一遍了直接跳过
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if nums[left] + nums[right] > target {
				right --
			} else if nums[left] + nums[right] < target {
				left ++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return res
}
