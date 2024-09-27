package task

import (
	"github.com/docker/docker/client"
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
// A task is the smallest unit of work in this orchestration system and runs on a container.
// Think - process on a machine
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

// Config struct is used to hold the docker container configuration
type Config struct {
	Name          string // name field used to identify the task
	AttachStdin   bool   // if the stdin should be attached
	AttachStdout  bool   // if the stdout should be attached
	AttachStderr  bool   // if the stderr should be attached
	ExposedPorts  nat.PortSet
	Cmd           []string // command within a container
	Image         string   // name of the image that the container runs
	Cpu           float64
	Memory        int64
	Disk          int64
	Env           []string // environment variables
	RestartPolicy string   // always / unless-stopped / on-failure
}

// Docker struct is used to run a task as a Docker container
type Docker struct {
	Client *client.Client
	Config Config
}

// DockerResult contains the result of the docker container execution
type DockerResult struct {
	Error       error  // Error is used to hold error messages
	Action      string // Action being taken, start, stop, etc
	ContainerId string // ContainerId has the current ID of the container being run
	Result      string // Result holds text to provide more information on the output
}
