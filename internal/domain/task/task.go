package task

import (
	"time"
)

type Task struct {
	UUID          string
	Name          string
	Description   string
	DockerImage   string
	DockerTag     string
	Port          string
	EnvVars       map[string]string
	Volumes       []string
	NetworkMode   string
	RestartPolicy string
	CpuLimit      int32
	MemoryLimit   int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TaskRepository interface {
	RunTask(id string) (*Task, error)
	CreateTask(task *Task) (*Task, error)
}
