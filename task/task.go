package task

import (
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"time"
)

type State int

const (
	Pending   State = iota // Initial state
	Scheduled              // Scheduled task => task going to the machine on which it'll be processed
	Running                // Currently working
	Completed              // Stopped by user or task completed
	Failed                 // Crash or stops working unexpected
)

// Task defines one task
type Task struct {
	ID            uuid.UUID         // UUID for each task to uniquely identify them
	Name          string            // Human-readable name
	State         State             // One of the above states
	Image         string            // Docker image the task is based on
	Memory        int               // Memory required for the task
	Disk          int               // Disk required for the task
	ExposedPorts  nat.PortSet       // Used by Docker to allocate the proper ports
	PortBindings  map[string]string // Used by Docker to allocate the proper ports
	RestartPolicy string            // To tell the system what to do in the event that the task stops/fails
	StartTime     time.Time         // Record when the current task starts
	EndTime       time.Time         // Record when the current task stops
}

// TaskEvent is used by the system to trigger Task state changes
type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
