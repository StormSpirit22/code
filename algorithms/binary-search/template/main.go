package main

import "fmt"

func main() {
	nums := []int{1,2,3,3,3,5,6}
	fmt.Println(binarySearch(nums, 3))		// 输出索引 3
	fmt.Println(leftBound(nums, 3))			// 输出索引 2
	fmt.Println(rightBound(nums, 3))			// 输出索引 4
}

// 基础二分搜索
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] ==  target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// 寻找左侧边界的二分搜索
func leftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return left
}

// 寻找右侧边界的二分搜索
func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return left - 1
}