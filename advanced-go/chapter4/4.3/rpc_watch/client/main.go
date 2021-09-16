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
	defer func(start time.Time) {
		fmt.Printf("doClientWork time cost %fs", time.Since(start).Seconds())
	}(time.Now())
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 10, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	} ()

	var err error
	//err := client.Call(
	//	"KVStoreService.Set", [2]string{"abc", "abc-value2"},
	//	new(struct{}),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	reply := ""
	err = client.Call("KVStoreService.Get", "abc", &reply)
	fmt.Println(reply)
	err = client.Call(
		"KVStoreService.Set", [2]string{"abc", "abc-value-modified2"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second*20)
}