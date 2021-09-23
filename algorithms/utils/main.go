package main

import "fmt"

func main() {
	fmt.Println(gcd(6, 15))
}


/*
*穷举法：最大公约数
 */
func gcdNormal(x, y int) int {
	var n int
	if x > y {
		n = y
	} else {
		n = x
	}
	for i := n; i >= 1; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}
	return 1
}

/*
*辗转相除法：最大公约数
*递归写法，进入运算是x和y都不为0
 */
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

/*
*辗转相除法：最大公约数
*非递归写法
 */
func gcdx(x, y int) int {
	var tmp int
	for {
		tmp = (x % y)
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}

/*
*穷举写法：最小公倍数
 */
func lcmNormal(x, y int) int {
	var top int = x * y
	var i = x
	if x < y {
		i = y
	}
	for ; i <= top; i++ {
		if i%x == 0 && i%y == 0 {
			return i
		}
	}
	return top
}

/*
*公式解法：最小公倍数=两数之积/最大公约数
 */
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
