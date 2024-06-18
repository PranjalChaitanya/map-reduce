package MapReduce

import "fmt"

type Node struct {
	value Task
	next  *Node
}

type TaskQueue struct {
	head *Node
	tail *Node
	size int
}

func (qe *TaskQueue) Enqueue(value Task) {
	n := &Node{value: value}
	if qe.size == 0 {
		qe.head = n
		qe.tail = n
	} else {
		qe.tail.next = n
		qe.tail = n
	}
	qe.size++
}

func (qe *TaskQueue) Dequeue() Task {
	if qe.size == 0 {
		return Task{}
	}

	tmp := qe.head.value
	qe.head = qe.head.next
	qe.size--

	fmt.Println(tmp.File)

	return tmp
}

func (qe *TaskQueue) Size() int {
	return qe.size
}
