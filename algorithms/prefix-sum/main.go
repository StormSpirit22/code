package main

func main() {

}

/*
https://leetcode-cn.com/problems/corporate-flight-bookings/
利用查分数组来求解。差分数组实际上就是求前缀和的逆运算。举个例子，对于一个数组 arr, 某个区间 [l, r] 需要增加 x，
可以求出来其对应的差分数组 nums（nums[i] = arr[i] - arr[i-1]），然后 nums[l] += x， nums[r+1] -= x，
之后对 nums 求前缀和即可。这是对区间修改的简化操作，不需要去遍历该区间来修改原数组。
在本题中已经给出各个区间的修改值，初始从 0 开始计算，那么查分数组初始也就是 n 个 0。
由于是在 [1,n] 中求解，所以相对应的差分数组应该是 nums[l-1] += x， nums[r] -= x。但是 r 可能等于 n ，
而差分数组最大索引是 n-1，所以当 r == n 时不做修改，而且本来 r = n 时相对于原数组 arr[n] 也没有意义。
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, b := range bookings {
		nums[b[0]-1] += b[2]
		if b[1] < n {
			nums[b[1]] -= b[2]
		}
	}
	for i := range nums {
		if i > 0 {
			nums[i] += nums[i-1]
		}
	}
	return nums
}