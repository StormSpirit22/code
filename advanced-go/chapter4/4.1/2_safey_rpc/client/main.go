package main

import (
	"code/advanced-go/chapter4/4.1/2_safey_rpc/client/client"
	"fmt"
	"log"
)

func main() {
	client, err := client.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	log.Println("dial tcp successful")

	var reply string
	err = client.Hello("hello", &reply)
	fmt.Println(reply)
	if err != nil {
		log.Fatal(err)
	}
}
