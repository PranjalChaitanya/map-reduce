package MapReduce

import (
	"fmt"
	"net/http"
	"net/rpc"
)

func (m *MasterServer) Init(empty struct{}, reply *Task) error {
	m.TaskQueue = TaskQueue{}
	a := Task{0, 0, 1, "tunu.txt"}
	b := Task{0, 0, 1, "hello.txt"}

	m.TaskQueue.Enqueue(a)
	m.TaskQueue.Enqueue(b)

	return nil
}

func (m *MasterServer) ProvideTask(empty struct{}, reply *Task) error {
	fmt.Println("IT IS GOING INSIDE OF HERE")
	if m.TaskQueue.Size() == 0 {
		fmt.Println("IT GOES INSIDE DOES IT FATAL?")
	}

	poppedTask := m.TaskQueue.Dequeue()

	reply.Type = poppedTask.Type
	reply.File = poppedTask.File
	reply.Status = poppedTask.Status
	reply.WorkerID = poppedTask.WorkerID

	return nil
}

func (m *MasterServer) Server() {
	var masterServer = new(MasterServer)
	rpc.Register(masterServer)
	rpc.HandleHTTP()
	http.ListenAndServe(":8080", nil)
}
