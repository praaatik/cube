package worker

import (
	"cube/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Worker is the layer above the Task which
// 1. Runs tasks as Docker containers
// 2. Accepts tasks to run from the Manager (use a Queue)
// 3. Provide Stats
// 4. Track Tasks and their State
type Worker struct {
	Name      string                   // Human-readable Name of the worker
	Queue     queue.Queue              // Manager assigns the Tasks here
	Db        map[uuid.UUID]*task.Task // Used to keep a track of the Tasks and their State
	TaskCount int                      // number of Tasks currently with the worker
}

// CollectStats is used to periodically collect statistics about the worker
func (w *Worker) CollectStats() {
	//TODO: write this
}

// RunTask handles the running of a task where the worker is running
func (w *Worker) RunTask() {
	//TODO: write this
}

// StartTask stops a task
func (w *Worker) StartTask() {
	//TODO: write this
}

// StopTask stops a task
func (w *Worker) StopTask() {
	//TODO: write this
}
