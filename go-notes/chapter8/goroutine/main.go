package main

import (
	"fmt"
	"time"
)

var c int

func counter() int {
	c++
	return c
}

func main()  {
	a := 100

	// 立即计算 counter(), 并复制参数 a
	go func(x, y int) {
		time.Sleep(time.Second)
		fmt.Printf("goroutine: a %d, c %d \n", x, y)
	}(a, counter())

	a += 100
	fmt.Printf("main: a %d, c %d \n", a, counter())
	time.Sleep(time.Second * 3)
}
