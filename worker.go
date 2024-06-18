package MapReduce

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

//func call(rpcname string, args interface{}, reply interface{}) bool {
//	c, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//	defer c.Close()
//
//	err = c.Call(rpcname, args, reply)
//	if err != nil {
//		return true
//	}
//
//	fmt.Println(err)
//	return false
//}

func (w *WorkerServer) ExecuteTask(task Task) {
	fmt.Println("Printing inside execute task")
	w.WorkerStatus = EXECUTING

	if task.Type == Map {
		go w.MapFileWordCount(task.File)
	}

	for w.WorkerStatus == EXECUTING {
		time.Sleep(2 * time.Second)
	}
}

func (w *WorkerServer) MapFileWordCount(key string) {
	value, err := os.ReadFile(key)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("intermediate.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(value), " ")
	for i := 0; i < len(words); i++ {
		fmt.Fprintf(file, "%s, %d\n", words[i], 1)
	}

	w.WorkerStatus = IDLE
}

func (w *WorkerServer) RequestTask() {
	fmt.Println("TASK IS BEING REQUESTED")
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	defer client.Close()
	if err != nil {
		log.Fatal("FATAL ERROR HAS OCCURED")
	}

	var reply Task

	client.Call("MasterServer.ProvideTask", struct{}{}, &reply)

	fmt.Println("COMPLETED CALL ABOUT TO RUN THE ROUTINE")

	w.ExecuteTask(reply)
}

func Server() {

}
