package MapReduce

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

func (w *WorkerServer) ExecuteTask(task Task) {
	fmt.Println("Printing inside execute task")
	w.WorkerStatus = EXECUTING

	if task.Type == Map {
		go w.MapFileWordCount(task.File)
	}
}

func (w *WorkerServer) MapFileWordCount(key string) {
	value, err := os.ReadFile(key)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("intermediate.txt", os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()
	if err != nil {
		file, err = os.Create("intermediate.txt")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	words := strings.Split(string(value), " ")
	for i := 0; i < len(words); i++ {
		fmt.Fprintf(file, "%s, %d\n", words[i], 1)
	}

	w.WorkerStatus = IDLE
}

func (w *WorkerServer) RequestTask() Task {
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
	return reply
}

func (w *WorkerServer) Server() {
	w.WorkerStatus = IDLE
	for {
		if w.WorkerStatus == IDLE {
			w.RequestTask()
		}
		time.Sleep(time.Second * 2)
	}
}
