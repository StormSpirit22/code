package main

func main() {
}

func interchangeableRectangles(rectangles [][]int) int64 {
	gMap := make(map[[2]int]int64)
	for _, r := range rectangles {
		m := gcd(r[0], r[1])
		a, b := r[0]/m, r[1]/m
		gMap[[2]int{a, b}]++
	}
	var ans int64
	for _, v := range gMap {
		ans += v*(v-1)/2
	}
	return ans
}

func gcd(x, y int) int {
	var tmp int
	for {
		tmp = x % y
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}