package main

import (
	"fmt"
	"time"
)

func goroutineA(a <-chan int) {
	val := <- a
	fmt.Println("G1 received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	val := <- b
	fmt.Println("G2 received data: ", val)
	return
}

func goroutineC(b <-chan int) {
	val := <- b
	fmt.Println("G3 received data: ", val)
	return
}

func goroutineD(b <-chan int) {
	val := <- b
	fmt.Println("G4 received data: ", val)
	return
}

func goroutineE(b <-chan int) {
	val := <- b
	fmt.Println("G5 received data: ", val)
	return
}

func main() {
	ch := make(chan int)
	go goroutineA(ch)
	go goroutineB(ch)
	//go goroutineC(ch)
	//go goroutineD(ch)
	//go goroutineE(ch)
	ch <- 3
	time.Sleep(time.Second)
}
