package main

import (
	"fmt"
)

func main() {
	nums := []int{1,3,2}
	nextPermutation(nums)
	fmt.Println(nums)
}

/*
https://leetcode-cn.com/problems/next-permutation
先从右边往左遍历，通过判断 nums[i] > nums[i-1] 找到左边的较小数 minLeft ，比如 [4,5,2,6,3,1]，较小数为 2 。
再从右边往左边李，通过判断 nums[i] > nums[minLeft] 找到右边的较大数 maxRight， 为 3 。
交换 nums[minLeft]， nums[maxRight]，成为 [4,5,3,6,2,1]， 可以知道交换后 minLeft 后面的必为降序序列，
然后将 nums[minLeft+1] 后面的序列交换即可。
*/
func nextPermutation(nums []int)  {
	n := len(nums)
	right := n-1
	minLeft, maxRight := 0, n-1
	for right > 0 {
		if nums[right] > nums[right-1] {
			minLeft = right-1
			break
		}
		right--
	}
	right = n - 1
	for right > minLeft {
		if nums[right] > nums[minLeft] {
			maxRight = right
			break
		}
		right--
	}
	if nums[minLeft] < nums[maxRight] {
		nums[minLeft], nums[maxRight] = nums[maxRight], nums[minLeft]
		reverse(minLeft+1, n-1, nums)
	} else {
		reverse(0, n-1, nums)
	}
}

func reverse(left, right int, nums []int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}