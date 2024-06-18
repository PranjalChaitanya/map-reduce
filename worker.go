package MapReduce

import (
	"log"
	"net/rpc"
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

func RequestNextTask() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("FATAL ERROR HAS OCCURED")
	}

	var reply Task

	client.Call("MasterServer.ProvideTask", struct{}{}, &reply)
}

func main() {
}
