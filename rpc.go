package MapReduce

type TaskType int
type TaskStatus int

const (
	Map TaskType = iota
	Reduce
)

const (
	NOT_STARTED TaskStatus = iota
	FAILED
	EXECUTING
)

type MasterServer struct {
	TaskQueue TaskQueue
}

type Task struct {
	Type     TaskType
	Status   TaskStatus
	WorkerID int
	File     string
}
