package main

func test() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		s = append(s, func() {				// 将多个匿名函数添加到列表
			println(&i, i)
		})
	}
	return s												// 返回匿名函数列表
}

func main() {
	for _, f := range test2() {			// 迭代执行所有匿名函数
		f()
	}
}

func test2() []func() {
	var s []func()
	for i := 0; i < 2; i++ {
		x := i									// x 每次循环都重新定义
		s = append(s, func() {
			println(&x, x)
		})
	}
	return s
}