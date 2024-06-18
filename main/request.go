package main

import (
	"MapReduce"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("FATAL ERROR HAS OCCURED")
	}

	var reply MapReduce.Task

	client.Call("MasterServer.ProvideTask", struct{}{}, &reply)
}
