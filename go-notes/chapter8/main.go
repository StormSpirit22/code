package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	n := runtime.GOMAXPROCS(0)
	println(n)
	//test(n)
	test2(n)
}

// 测试目标函数
func count() {
	x := 0
	for i := 0; i < math.MaxInt32; i++ {
		x += i
	}

	println(x)
}


// 循环执行
func test(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

// 并发执行
func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}
	wg.Wait()
}


func timeTest() {
	// timeout
	go func() {
		for {
			select {
			case <- time.After(time.Second * 5):
				fmt.Println("timeout...")
				os.Exit(0)
			}
		}
	}()

	// tick
	go func() {
		tick := time.Tick(time.Second)

		for {
			select {
			case <- tick:
				fmt.Println(time.Now())
			}
		}
	}()
}