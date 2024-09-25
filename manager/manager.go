package manager

import (
	"cube/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Manager struct defines the manager.
// Manager will
// 1. Accept requests from users to start/stop tasks
// 2. Schedule tasks on the worker machine
// 3. Keep track of the tasks and on which machine they run
type Manager struct {
	Pending       queue.Queue             // Pending tasks which need to be placed on machines
	TaskDb        map[string][]*task.Task // in-memory database to store the Tasks
	EventDb       map[string][]*task.Task // in-memory database to store the Events
	Workers       []string                // Keep a track of the workers in the cluster
	WorkerTaskMap map[string][]uuid.UUID  // To keep a track of the jobs assigned to each worker
	TaskWorkerMap map[uuid.UUID]string    // To keep a track of the workers running a task
}

func (m *Manager) SelectWorker() {
	// TODO: add this
}

func (m *Manager) UpdateTasks() {
	// TODO: add this
}

func (m *Manager) SendWork() {
	// TODO: add this
}
