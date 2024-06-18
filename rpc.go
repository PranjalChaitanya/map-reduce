package MapReduce

type TaskType int
type TaskStatus int

type WorkerStatus int

const (
	Map TaskType = iota
	Reduce
)

const (
	NOT_STARTED TaskStatus = iota
	FAILED
	EXECUTING_TASK
)

const (
	IDLE WorkerStatus = iota
	EXECUTING
)

type MasterServer struct {
	TaskQueue TaskQueue
}

type WorkerServer struct {
	WorkerStatus WorkerStatus
}

type Task struct {
	Type     TaskType
	Status   TaskStatus
	WorkerID int
	File     string
}
