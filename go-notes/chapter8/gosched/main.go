package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		defer close(exit)

		go func() {
			fmt.Println("b")
		}()

		for i := 0; i < 4; i++ {
			fmt.Println("a, ", i)
			if i == 1 {
				runtime.Gosched()
			}
		}
	}()

	<- exit
}


func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		runtime.Gosched()
	}
}

func test() {
	runtime.GOMAXPROCS(1)
	go say("world")
	say("hello")
}
