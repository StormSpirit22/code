package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	doClientWork(client)
}

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	} ()

	err := client.Call(
		"KVStoreService.Set", [2]string{"abc", "abc-value2"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Call(
		"KVStoreService.Set", [2]string{"abc", "abc-value-modified2"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second*3)
}